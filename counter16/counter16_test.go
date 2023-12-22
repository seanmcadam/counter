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
