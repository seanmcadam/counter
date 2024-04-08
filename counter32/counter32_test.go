package counter32

import (
	"testing"

	"github.com/seanmcadam/ctx"
)

func TestCounter32(t *testing.T) {
	cx := ctx.New()
	c32 := New(cx)
	_ = c32.Next()
	cx.Cancel()
}


func TestCounterInt32(t *testing.T) {
	cx := ctx.New()
	c32 := New(cx)
	_ = c32.Bits()

	for i := 0; i < 100; i++ {
		ci := c32.Next()
		_ = ci.Bits()
		_ = ci.Copy()
		_ = ci.Uint()
		_ = ci.ToByte()
		_ = ci.String()
		_ = NewCount(uint32(ci.Uint()))
		_, _ = c32.ByteToCounter(ci.ToByte())
	}

	cx.Cancel()
}
