package counter8

import (
	"log"
	"testing"

	"github.com/seanmcadam/ctx"
)

func TestCounter8(t *testing.T) {
	cx := ctx.New()
	c8 := New(cx)
	_ = c8.Next()
	cx.Cancel()
}

func TestCounter8_overflow(t *testing.T) {
	cx := ctx.New()
	c8 := New(cx)

	for i := 0; i < 258; i++ {
		log.Printf("I:%d", c8.Next())
	}

	cx.Cancel()
}


func TestCounterInt8(t *testing.T) {
	cx := ctx.New()
	c8 := New(cx)
	_ = c8.Bits()

	for i := 0; i < 100; i++ {
		ci := c8.Next()
		_ = ci.Bits()
		_ = ci.Copy()
		_ = ci.Uint()
		_ = ci.ToByte()
		_ = ci.String()
		_ = NewCount(uint8(ci.Uint()))
		_, _ = c8.ByteToCounter(ci.ToByte())
	}

	cx.Cancel()
}
