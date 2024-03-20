package counter64

import (
	"encoding/binary"
	"fmt"

	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/counterint"
)

type Counter64 uint64

// NewCount create a new 64bit counter
func NewCount(c uint64) counterint.CountInt {
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

func (c *Counter64) Copy() counterint.CountInt {
	copy := *c
	return &copy
}

func (c *Counter64) String() string {
	return fmt.Sprintf("%d", c)
}
