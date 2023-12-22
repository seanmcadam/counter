package counter8

import (
	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/counterint"
)

type Counter8 uint8

func NewCount(c uint8) counterint.CounterInt {
	c8 := Counter8(c)
	return &c8
}

func (*Counter8) Bits() common.CounterBits {
	return common.BIT8
}

func (c *Counter8) Uint() uint64 {
	return uint64(*c)
}

func (c *Counter8) ToByte() (b []byte) {
	b = make([]byte, 1)
	b[0] = byte(*c)
	return b
}

func (c *Counter8) Copy() counterint.CounterInt {
	copy := *c
	return &copy
}
