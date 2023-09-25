package net

import (
	"crypto/ecdsa"
	"net"
	"reflect"
	"testing"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/netutil"
	"github.com/holiman/uint256"
	kbucket "github.com/libp2p/go-libp2p-kbucket"
)

func TestDefaultConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewProtocol(t *testing.T) {
	type args struct {
		config     *Config
		protocolId string
		privateKey *ecdsa.PrivateKey
	}
	tests := []struct {
		name    string
		args    args
		want    *Protocol
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProtocol(tt.args.config, tt.args.protocolId, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProtocol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProtocol() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProtocol_Ping(t *testing.T) {
	type fields struct {
		buckets        *kbucket.RoutingTable
		protocolId     string
		nodeRadius     *uint256.Int
		Discovery      *discover.UDPv5
		ListenAddr     string
		localNode      *enode.LocalNode
		log            log.Logger
		discmix        *enode.FairMix
		PrivateKey     *ecdsa.PrivateKey
		NetRestrict    *netutil.Netlist
		BootstrapNodes []*enode.Node
		radiusCache    *fastcache.Cache
	}
	type args struct {
		node *enode.Node
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Protocol{
				buckets:        tt.fields.buckets,
				protocolId:     tt.fields.protocolId,
				nodeRadius:     tt.fields.nodeRadius,
				Discovery:      tt.fields.Discovery,
				ListenAddr:     tt.fields.ListenAddr,
				localNode:      tt.fields.localNode,
				log:            tt.fields.log,
				discmix:        tt.fields.discmix,
				PrivateKey:     tt.fields.PrivateKey,
				NetRestrict:    tt.fields.NetRestrict,
				BootstrapNodes: tt.fields.BootstrapNodes,
				radiusCache:    tt.fields.radiusCache,
			}
			if err := p.Ping(tt.args.node); (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProtocol_Start(t *testing.T) {
	type fields struct {
		buckets        *kbucket.RoutingTable
		protocolId     string
		nodeRadius     *uint256.Int
		Discovery      *discover.UDPv5
		ListenAddr     string
		localNode      *enode.LocalNode
		log            log.Logger
		discmix        *enode.FairMix
		PrivateKey     *ecdsa.PrivateKey
		NetRestrict    *netutil.Netlist
		BootstrapNodes []*enode.Node
		radiusCache    *fastcache.Cache
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Protocol{
				buckets:        tt.fields.buckets,
				protocolId:     tt.fields.protocolId,
				nodeRadius:     tt.fields.nodeRadius,
				Discovery:      tt.fields.Discovery,
				ListenAddr:     tt.fields.ListenAddr,
				localNode:      tt.fields.localNode,
				log:            tt.fields.log,
				discmix:        tt.fields.discmix,
				PrivateKey:     tt.fields.PrivateKey,
				NetRestrict:    tt.fields.NetRestrict,
				BootstrapNodes: tt.fields.BootstrapNodes,
				radiusCache:    tt.fields.radiusCache,
			}
			if err := p.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProtocol_handlePing(t *testing.T) {
	type fields struct {
		buckets        *kbucket.RoutingTable
		protocolId     string
		nodeRadius     *uint256.Int
		Discovery      *discover.UDPv5
		ListenAddr     string
		localNode      *enode.LocalNode
		log            log.Logger
		discmix        *enode.FairMix
		PrivateKey     *ecdsa.PrivateKey
		NetRestrict    *netutil.Netlist
		BootstrapNodes []*enode.Node
		radiusCache    *fastcache.Cache
	}
	type args struct {
		id   enode.ID
		ping *Ping
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Protocol{
				buckets:        tt.fields.buckets,
				protocolId:     tt.fields.protocolId,
				nodeRadius:     tt.fields.nodeRadius,
				Discovery:      tt.fields.Discovery,
				ListenAddr:     tt.fields.ListenAddr,
				localNode:      tt.fields.localNode,
				log:            tt.fields.log,
				discmix:        tt.fields.discmix,
				PrivateKey:     tt.fields.PrivateKey,
				NetRestrict:    tt.fields.NetRestrict,
				BootstrapNodes: tt.fields.BootstrapNodes,
				radiusCache:    tt.fields.radiusCache,
			}
			got, err := p.handlePing(tt.args.id, tt.args.ping)
			if (err != nil) != tt.wantErr {
				t.Errorf("handlePing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handlePing() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProtocol_handleTalkRequest(t *testing.T) {
	type fields struct {
		buckets        *kbucket.RoutingTable
		protocolId     string
		nodeRadius     *uint256.Int
		Discovery      *discover.UDPv5
		ListenAddr     string
		localNode      *enode.LocalNode
		log            log.Logger
		discmix        *enode.FairMix
		PrivateKey     *ecdsa.PrivateKey
		NetRestrict    *netutil.Netlist
		BootstrapNodes []*enode.Node
		radiusCache    *fastcache.Cache
	}
	type args struct {
		id   enode.ID
		addr *net.UDPAddr
		msg  []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Protocol{
				buckets:        tt.fields.buckets,
				protocolId:     tt.fields.protocolId,
				nodeRadius:     tt.fields.nodeRadius,
				Discovery:      tt.fields.Discovery,
				ListenAddr:     tt.fields.ListenAddr,
				localNode:      tt.fields.localNode,
				log:            tt.fields.log,
				discmix:        tt.fields.discmix,
				PrivateKey:     tt.fields.PrivateKey,
				NetRestrict:    tt.fields.NetRestrict,
				BootstrapNodes: tt.fields.BootstrapNodes,
				radiusCache:    tt.fields.radiusCache,
			}
			if got := p.handleTalkRequest(tt.args.id, tt.args.addr, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleTalkRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProtocol_processPong(t *testing.T) {
	type fields struct {
		buckets        *kbucket.RoutingTable
		protocolId     string
		nodeRadius     *uint256.Int
		Discovery      *discover.UDPv5
		ListenAddr     string
		localNode      *enode.LocalNode
		log            log.Logger
		discmix        *enode.FairMix
		PrivateKey     *ecdsa.PrivateKey
		NetRestrict    *netutil.Netlist
		BootstrapNodes []*enode.Node
		radiusCache    *fastcache.Cache
	}
	type args struct {
		node *enode.Node
		pong *Pong
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Protocol{
				buckets:        tt.fields.buckets,
				protocolId:     tt.fields.protocolId,
				nodeRadius:     tt.fields.nodeRadius,
				Discovery:      tt.fields.Discovery,
				ListenAddr:     tt.fields.ListenAddr,
				localNode:      tt.fields.localNode,
				log:            tt.fields.log,
				discmix:        tt.fields.discmix,
				PrivateKey:     tt.fields.PrivateKey,
				NetRestrict:    tt.fields.NetRestrict,
				BootstrapNodes: tt.fields.BootstrapNodes,
				radiusCache:    tt.fields.radiusCache,
			}
			if err := p.processPong(tt.args.node, tt.args.pong); (err != nil) != tt.wantErr {
				t.Errorf("processPong() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProtocol_sendReqAndHandleResp(t *testing.T) {
	type fields struct {
		buckets        *kbucket.RoutingTable
		protocolId     string
		nodeRadius     *uint256.Int
		Discovery      *discover.UDPv5
		ListenAddr     string
		localNode      *enode.LocalNode
		log            log.Logger
		discmix        *enode.FairMix
		PrivateKey     *ecdsa.PrivateKey
		NetRestrict    *netutil.Netlist
		BootstrapNodes []*enode.Node
		radiusCache    *fastcache.Cache
	}
	type args struct {
		node             *enode.Node
		err              error
		talkRequestBytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Protocol{
				buckets:        tt.fields.buckets,
				protocolId:     tt.fields.protocolId,
				nodeRadius:     tt.fields.nodeRadius,
				Discovery:      tt.fields.Discovery,
				ListenAddr:     tt.fields.ListenAddr,
				localNode:      tt.fields.localNode,
				log:            tt.fields.log,
				discmix:        tt.fields.discmix,
				PrivateKey:     tt.fields.PrivateKey,
				NetRestrict:    tt.fields.NetRestrict,
				BootstrapNodes: tt.fields.BootstrapNodes,
				radiusCache:    tt.fields.radiusCache,
			}
			if err := p.sendReqAndHandleResp(tt.args.node, tt.args.err, tt.args.talkRequestBytes); (err != nil) != tt.wantErr {
				t.Errorf("sendReqAndHandleResp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProtocol_setupDiscovery(t *testing.T) {
	type fields struct {
		buckets        *kbucket.RoutingTable
		protocolId     string
		nodeRadius     *uint256.Int
		Discovery      *discover.UDPv5
		ListenAddr     string
		localNode      *enode.LocalNode
		log            log.Logger
		discmix        *enode.FairMix
		PrivateKey     *ecdsa.PrivateKey
		NetRestrict    *netutil.Netlist
		BootstrapNodes []*enode.Node
		radiusCache    *fastcache.Cache
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Protocol{
				buckets:        tt.fields.buckets,
				protocolId:     tt.fields.protocolId,
				nodeRadius:     tt.fields.nodeRadius,
				Discovery:      tt.fields.Discovery,
				ListenAddr:     tt.fields.ListenAddr,
				localNode:      tt.fields.localNode,
				log:            tt.fields.log,
				discmix:        tt.fields.discmix,
				PrivateKey:     tt.fields.PrivateKey,
				NetRestrict:    tt.fields.NetRestrict,
				BootstrapNodes: tt.fields.BootstrapNodes,
				radiusCache:    tt.fields.radiusCache,
			}
			if err := p.setupDiscovery(); (err != nil) != tt.wantErr {
				t.Errorf("setupDiscovery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProtocol_setupUDPListening(t *testing.T) {
	type fields struct {
		buckets        *kbucket.RoutingTable
		protocolId     string
		nodeRadius     *uint256.Int
		Discovery      *discover.UDPv5
		ListenAddr     string
		localNode      *enode.LocalNode
		log            log.Logger
		discmix        *enode.FairMix
		PrivateKey     *ecdsa.PrivateKey
		NetRestrict    *netutil.Netlist
		BootstrapNodes []*enode.Node
		radiusCache    *fastcache.Cache
	}
	tests := []struct {
		name    string
		fields  fields
		want    *net.UDPConn
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Protocol{
				buckets:        tt.fields.buckets,
				protocolId:     tt.fields.protocolId,
				nodeRadius:     tt.fields.nodeRadius,
				Discovery:      tt.fields.Discovery,
				ListenAddr:     tt.fields.ListenAddr,
				localNode:      tt.fields.localNode,
				log:            tt.fields.log,
				discmix:        tt.fields.discmix,
				PrivateKey:     tt.fields.PrivateKey,
				NetRestrict:    tt.fields.NetRestrict,
				BootstrapNodes: tt.fields.BootstrapNodes,
				radiusCache:    tt.fields.radiusCache,
			}
			got, err := p.setupUDPListening()
			if (err != nil) != tt.wantErr {
				t.Errorf("setupUDPListening() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupUDPListening() got = %v, want %v", got, tt.want)
			}
		})
	}
}
