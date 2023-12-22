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
