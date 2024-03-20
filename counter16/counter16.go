package counter16

import (
	"encoding/binary"
	"fmt"

	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/counterint"
)

type Counter16 uint16

// NewCount create a new 16bit counter
func NewCount(c uint16) counterint.CountInt {
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

func (c *Counter16) Copy() counterint.CountInt {
	copy := *c
	return &copy
}

func (c *Counter16) String() string {
	return fmt.Sprintf("%d", c)
}
