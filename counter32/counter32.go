package counter32

import (
	"encoding/binary"
	"fmt"

	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/counterint"
)

type Counter32 uint32

func NewCount(c uint32) counterint.CounterInt {
	c32 := Counter32(c)
	return &c32
}

func (*Counter32) Bits() common.CounterBits {
	return common.BIT32
}

func (c *Counter32) Uint() uint64 {
	return uint64(*c)
}

func (c *Counter32) ToByte() (b []byte) {
	b = make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(*c))
	return b
}

func (c *Counter32) Copy() counterint.CounterInt {
	copy := *c
	return &copy
}

func (c *Counter32) String() string {
	return fmt.Sprintf("%d", c)
}
