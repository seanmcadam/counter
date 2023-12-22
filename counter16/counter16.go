package counter16

import (
	"encoding/binary"

	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/counterint"
)

type Counter16 uint16

func NewCount(c uint16) counterint.CounterInt {
	c16 := Counter16(c)
	return &c16
}

func (*Counter16) Bits() common.CounterBits {
	return common.BIT16
}

func (c *Counter16) Uint() uint64 {
	return uint64(*c)
}

func (c *Counter16) ToByte() (b []byte) {
	b = make([]byte, 4)
	binary.BigEndian.PutUint16(b, uint16(*c))
	return b
}

func (c *Counter16) Copy() counterint.CounterInt {
	copy := *c
	return &copy
}
