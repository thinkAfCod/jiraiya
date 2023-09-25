package net

import (
	"crypto/ecdsa"
	"net"
	"time"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/bits-and-blooms/bitset"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/netutil"
	"github.com/holiman/uint256"
	kbucket "github.com/libp2p/go-libp2p-kbucket"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/host/peerstore"
)

const (
	// This is the fairness knob for the discovery mixer. When looking for peers, we'll
	// wait this long for a single source of candidates before moving on and trying other
	// sources.
	discmixTimeout = 5 * time.Second
)

// ProtocolId is the protocol id for the overlay protocol.
const (
	StateNetwork             = "0x500a"
	HistoryNetwork           = "0x500b"
	TxGossipNetwork          = "0x500c"
	HeaderGossipNetwork      = "0x500d"
	CanonicalIndicesNetwork  = "0x500e"
	BeaconLightClientNetwork = "0x501a"
	UTPNetwork               = "0x757470"
	Rendezvous               = "0x72656e"
)

// Message codes for the overlay protocol.
const (
	PING        = 0x00
	PONG        = 0x01
	FINDNODES   = 0x02
	NODES       = 0x03
	FINDCONTENT = 0x04
	CONTENT     = 0x05
	OFFER       = 0x06
	ACCEPT      = 0x07
)

type Config struct {
	BootstrapNodes []*enode.Node

	ListenAddr      string
	NetRestrict     *netutil.Netlist
	NodeRadius      *uint256.Int
	RadiusCacheSize int
}

func DefaultConfig() *Config {
	nodeRadius, _ := uint256.FromHex("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	return &Config{
		BootstrapNodes:  make([]*enode.Node, 0),
		ListenAddr:      ":9000",
		NetRestrict:     &netutil.Netlist{},
		NodeRadius:      nodeRadius,
		RadiusCacheSize: 32 * 1024 * 1024,
	}
}

type OverlayMsgId [16]byte

type RawContentKey []byte

type ContentItem struct {
	ContentKey RawContentKey
	Content    []byte
}

type PortalWireProtocolMsg interface {
	Id() OverlayMsgId
}

type (
	PingPongCustomData struct {
		Radius []byte `ssz-size:"32"`
	}

	Ping struct {
		EnrSeq        uint64
		CustomPayload []byte `ssz-max:"2048"`
	}

	FindNodes struct {
		Distances []uint16
	}

	FindContent struct {
		ContentKey []byte
	}

	Offer struct {
		ContentKeys []RawContentKey
	}

	//PopulatedOffer struct {
	//	ContentItems []*ContentItem
	//}
)

type (
	Pong struct {
		EnrSeq        uint64
		CustomPayload []byte `ssz-max:"2048"`
	}

	Nodes struct {
		Total uint8
		Enrs  []*enode.Node
	}

	//ConnectionIdContent struct {
	//	Id uint16
	//}
	//
	//RawContent struct {
	//	Raw []byte
	//}
	//
	//EnrsContent struct {
	//	Enrs []*enode.Node
	//}

	Accept struct {
		ConnectionId uint16
		ContentKeys  bitset.BitSet
	}
)

type Protocol struct {
	buckets *kbucket.RoutingTable

	protocolId string

	nodeRadius     *uint256.Int
	Discovery      *discover.UDPv5
	ListenAddr     string
	localNode      *enode.LocalNode
	log            log.Logger
	discmix        *enode.FairMix
	PrivateKey     *ecdsa.PrivateKey
	NetRestrict    *netutil.Netlist
	BootstrapNodes []*enode.Node

	radiusCache *fastcache.Cache
}

func NewProtocol(config *Config, protocolId string, privateKey *ecdsa.PrivateKey) (*Protocol, error) {
	protocol := &Protocol{
		protocolId:     protocolId,
		ListenAddr:     config.ListenAddr,
		log:            log.New("protocol", protocolId),
		PrivateKey:     privateKey,
		NetRestrict:    config.NetRestrict,
		BootstrapNodes: config.BootstrapNodes,
		nodeRadius:     config.NodeRadius,
		radiusCache:    fastcache.New(config.RadiusCacheSize),
	}

	return protocol, nil

}

func (p *Protocol) Start() error {
	err := p.setupDiscovery()
	if err != nil {
		return err
	}

	buckets, err := kbucket.NewRoutingTable(256, p.localNode.ID().Bytes(), time.Minute, peerstore.NewMetrics(), 2*time.Hour, nil)
	if err != nil {
		return err
	}
	p.buckets = buckets
	p.Discovery.RegisterTalkHandler(p.protocolId, p.handleTalkRequest)
	return nil
}

func (p *Protocol) setupUDPListening() (*net.UDPConn, error) {
	listenAddr := p.ListenAddr

	addr, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	laddr := conn.LocalAddr().(*net.UDPAddr)
	p.localNode.SetFallbackUDP(laddr.Port)
	p.log.Debug("UDP listener up", "addr", laddr)
	// TODO: NAT
	//if !laddr.IP.IsLoopback() && !laddr.IP.IsPrivate() {
	//	srv.portMappingRegister <- &portMapping{
	//		protocol: "UDP",
	//		name:     "ethereum peer discovery",
	//		port:     laddr.Port,
	//	}
	//}

	return conn, nil
}

func (p *Protocol) setupDiscovery() error {
	p.discmix = enode.NewFairMix(discmixTimeout)

	conn, err := p.setupUDPListening()
	if err != nil {
		return err
	}

	cfg := discover.Config{
		PrivateKey:  p.PrivateKey,
		NetRestrict: p.NetRestrict,
		Bootnodes:   p.BootstrapNodes,
		Log:         p.log,
	}
	p.Discovery, err = discover.ListenV5(conn, p.localNode, cfg)
	if err != nil {
		return err
	}

	return nil
}

func (p *Protocol) Ping(node *enode.Node) error {
	enrSeq := p.Discovery.LocalNode().Seq()
	radiusBytes, err := p.nodeRadius.MarshalSSZ()
	if err != nil {
		return err
	}
	customPayload := &PingPongCustomData{
		Radius: radiusBytes,
	}

	customPayloadBytes, err := customPayload.MarshalSSZ()
	if err != nil {
		return err
	}

	pingRequest := &Ping{
		EnrSeq:        enrSeq,
		CustomPayload: customPayloadBytes,
	}

	pingRequestBytes, err := pingRequest.MarshalSSZ()
	if err != nil {
		return err
	}

	talkRequestBytes := make([]byte, 0, len(pingRequestBytes)+1)
	talkRequestBytes = append(talkRequestBytes, PING)
	talkRequestBytes = append(talkRequestBytes, pingRequestBytes...)

	return p.sendReqAndHandleResp(node, err, talkRequestBytes)
}

func (p *Protocol) sendReqAndHandleResp(node *enode.Node, err error, talkRequestBytes []byte) error {
	talkResp, err := p.Discovery.TalkRequest(node, p.protocolId, talkRequestBytes)
	if err != nil {
		p.buckets.RemovePeer(peer.ID(node.ID().String()))
		return err
	}

	p.buckets.UpdateLastSuccessfulOutboundQueryAt(peer.ID(node.ID().String()), time.Now())

	switch talkResp[0] {
	case PONG:
		pong := &Pong{}
		err = pong.UnmarshalSSZ(talkResp[1:])
		if err != nil {
			return err
		}

		err = p.processPong(node, pong)
		if err != nil {
			return err
		}

	}

	return nil
}

func (p *Protocol) processPong(node *enode.Node, pong *Pong) error {
	customPayload := &PingPongCustomData{}
	err := customPayload.UnmarshalSSZ(pong.CustomPayload)
	if err != nil {
		return err
	}
	p.radiusCache.Set([]byte(node.ID().String()), customPayload.Radius)
	return nil
}

func (p *Protocol) handleTalkRequest(id enode.ID, addr *net.UDPAddr, msg []byte) []byte {
	if p.buckets.Find(peer.ID(id.String())) == "" {
		nodes := p.Discovery.Lookup(id)
		if len(nodes) > 0 && nodes[0].ID() == id {
			_, err := p.buckets.TryAddPeer(peer.ID(id.String()), true, true)
			if err != nil {
				p.log.Error("failed to add peer to buckets", "err", err)
				return nil
			}
		}
		// TODO: emit event
	}

	msgCode := msg[0]

	switch msgCode {
	case PING:
		pingRequest := &Ping{}
		err := pingRequest.UnmarshalSSZ(msg[1:])
		if err != nil {
			p.log.Error("failed to unmarshal ping request", "err", err)
			return nil
		}

		p.log.Trace("received ping request", "protocol", p.protocolId, "source", id, "pingRequest", pingRequest)
		resp, err := p.handlePing(id, pingRequest)
		if err != nil {
			p.log.Error("failed to handle ping request", "err", err)
			return nil
		}

		return resp
	}

	return nil
}

func (p *Protocol) handlePing(id enode.ID, ping *Ping) ([]byte, error) {
	pingCustomPayload := &PingPongCustomData{}
	err := pingCustomPayload.UnmarshalSSZ(ping.CustomPayload)
	if err != nil {
		return nil, err
	}

	p.radiusCache.Set([]byte(id.String()), pingCustomPayload.Radius)

	enrSeq := p.Discovery.LocalNode().Seq()
	radiusBytes, err := p.nodeRadius.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	pongCustomPayload := &PingPongCustomData{
		Radius: radiusBytes,
	}

	pongCustomPayloadBytes, err := pongCustomPayload.MarshalSSZ()
	if err != nil {
		return nil, err
	}

	pong := &Pong{
		EnrSeq:        enrSeq,
		CustomPayload: pongCustomPayloadBytes,
	}

	pongBytes, err := pong.MarshalSSZ()

	return pongBytes, err
}
