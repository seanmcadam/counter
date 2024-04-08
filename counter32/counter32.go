package counter32

import (
	"encoding/binary"
	"fmt"

	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/counterint"
	"github.com/seanmcadam/loggy"
)

type Counter32 uint32

// NewCount create a new 32bit counter
func NewCount(c uint32) counterint.CountInt {
	c32 := Counter32(c)
	return &c32
}

func (c *Counter32) Bits() common.CounterBits {
	c.checkfornil()
	return common.BIT32
}

func (c *Counter32) Uint() uint64 {
	return uint64(*c)
}

func (c *Counter32) ToByte() (b []byte) {
	c.checkfornil()
	b = make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(*c))
	return b
}

func (c *Counter32) Copy() counterint.CountInt {
	c.checkfornil()
	copy := *c
	return &copy
}

func (c *Counter32) String() string {
	c.checkfornil()
	return fmt.Sprintf("%d", c)
}

func (c *Counter32) checkfornil() {
	if c == nil {
		loggy.FatalStack("nil method")
	}
}
