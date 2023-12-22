package counter64

import (
	"encoding/binary"

	"github.com/seanmcadam/counter/counterint"
	"github.com/seanmcadam/counter/common"
)

type Counter64 uint64

func NewCount(c uint64) counterint.CounterInt {
	c64 := Counter64(c)
	return &c64
}

func (*Counter64) Bits() common.CounterBits {
	return common.BIT64
}

func (c *Counter64) Uint() uint64 {
	return uint64(*c)
}

func (c *Counter64) ToByte() (b []byte) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(*c))
	return b
}

func (c *Counter64) Copy() counterint.CounterInt {
	copy := *c
	return &copy
}
