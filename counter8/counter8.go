package counter8

import (
	"fmt"

	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/counterint"
	"github.com/seanmcadam/loggy"
)

type Counter8 uint8

// NewCount create a new 8bit counter
func NewCount(c uint8) counterint.CountInt {
	c8 := Counter8(c)
	return &c8
}

func (c *Counter8) Bits() common.CounterBits {
	c.checkfornil()
	return common.BIT8
}

func (c *Counter8) Uint() uint64 {
	return uint64(*c)
}

func (c *Counter8) ToByte() (b []byte) {
	c.checkfornil()
	b = make([]byte, 1)
	b[0] = byte(*c)
	return b
}

func (c *Counter8) Copy() counterint.CountInt {
	c.checkfornil()
	copy := *c
	return &copy
}

func (c *Counter8) String() string {
	c.checkfornil()
	return fmt.Sprintf("%d", c)
}

func (c *Counter8) checkfornil() {
	if c == nil {
		loggy.FatalStack("nil method")
	}
}
