package api

import "github.com/holiman/uint256"

type Distance struct {
	Dist uint256.Int
}

func (d Distance) SSZBytes() ([]byte, error) {
	ssz, err := d.Dist.MarshalSSZ()
	if err != nil {
		return nil, err
	}

	return ssz, nil
}

type Metrics interface {
	Distance(x [32]byte, y [32]byte) Distance
}
