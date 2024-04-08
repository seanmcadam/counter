package counter64

import (
	"testing"

	"github.com/seanmcadam/ctx"
)

func TestCounter64(t *testing.T) {
	cx := ctx.New()
	c64 := New(cx)
	_ = c64.Next()
	cx.Cancel()
}

func TestCounterInt64(t *testing.T) {
	cx := ctx.New()
	c64 := New(cx)
	_ = c64.Bits()

	for i := 0; i < 100; i++ {
		ci := c64.Next()
		_ = ci.Bits()
		_ = ci.Copy()
		_ = ci.Uint()
		_ = ci.ToByte()
		_ = ci.String()
		_ = NewCount(ci.Uint())
		_, _ = c64.ByteToCounter(ci.ToByte())
	}

	cx.Cancel()
}
