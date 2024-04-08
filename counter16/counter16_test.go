package counter16

import (
	"testing"

	"github.com/seanmcadam/ctx"
)

func TestCounter16(t *testing.T) {
	cx := ctx.New()
	c16 := New(cx)
	_ = c16.Next()
	cx.Cancel()
}


func TestCounterInt16(t *testing.T) {
	cx := ctx.New()
	c16 := New(cx)
	_ = c16.Bits()

	for i := 0; i < 100; i++ {
		ci := c16.Next()
		_ = ci.Bits()
		_ = ci.Copy()
		_ = ci.Uint()
		_ = ci.ToByte()
		_ = ci.String()
		_ = NewCount(uint16(ci.Uint()))
		_, _ = c16.ByteToCounter(ci.ToByte())
	}

	cx.Cancel()
}
