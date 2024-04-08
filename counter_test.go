package counter

import (
	"testing"
	"time"

	"github.com/seanmcadam/ctx"
)

func TestCompile(t *testing.T) {}

func TestNew64(t *testing.T) {
	cx := ctx.New().WithCancel()
	c := New64(cx)

	for i := 1; i < 10000; i++ {
		j := c.Next()
		if uint64(i) != j.Uint() {
			t.Fatalf("i:%d != j:%d", i, *j)
		}
		j.ToBEByte()
		j.String()
	}
	cx.Cancel()
	time.Sleep(time.Millisecond*10)
}

func TestNew32(t *testing.T) {
	cx := ctx.New().WithCancel()
	c := New32(cx)

	for i := 1; i < 10000; i++ {
		j := c.Next()
		if uint32(i) != j.Uint() {
			t.Fatalf("i:%d != j:%d", i, *j)
		}
		j.ToBEByte()
		j.String()
	}
}

func TestNew16(t *testing.T) {
	cx := ctx.New().WithCancel()
	c := New16(cx)

	for i := 1; i < 10000; i++ {
		j := c.Next()
		if uint16(i) != j.Uint() {
			t.Fatalf("i:%d != j:%d", i, *j)
		}
		j.ToBEByte()
		j.String()
	}
}

func TestNew8(t *testing.T) {
	cx := ctx.New().WithCancel()
	c := New8(cx)

	for i := 1; i < 10000; i++ {
		j := c.Next()
		if uint8(i) != j.Uint() {
			t.Fatalf("i:%d != j:%d", i, *j)
		}
		j.ToBEByte()
		j.String()
	}
}
