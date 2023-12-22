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
