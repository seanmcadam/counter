package counter

import (
	"fmt"
	"reflect"

	"github.com/seanmcadam/ctx"
	"github.com/seanmcadam/loggy"
	"golang.org/x/exp/constraints"
)

type Count[T constraints.Unsigned] struct{ count T }

type CounterStruct[T constraints.Unsigned] struct {
	cx *ctx.Ctx
	ch chan *Count[T]
}

func New64(cx *ctx.Ctx) *CounterStruct[uint64] {
	return New[uint64](cx)
}
func New32(cx *ctx.Ctx) *CounterStruct[uint32] {
	return New[uint32](cx)
}
func New16(cx *ctx.Ctx) *CounterStruct[uint16] {
	return New[uint16](cx)
}
func New8(cx *ctx.Ctx) *CounterStruct[uint8] {
	return New[uint8](cx)
}

func New[T constraints.Unsigned](cx *ctx.Ctx) (c *CounterStruct[T]) {
	if cx == nil {
		loggy.Fatal("nil ctx")
	}
	c = &CounterStruct[T]{
		cx: cx,
		ch: make(chan *Count[T], 5),
	}
	go c.goRun()
	return c
}

func (c *CounterStruct[T]) Next() *Count[T] {
	c.checkfornil()
	return <-c.ch
}

func (c *Count[T]) Uint() (u T) {
	u = c.count
	return u
}

func (c *Count[T]) ToBEByte() (b []byte) {
	c.checkfornil()
	switch reflect.TypeOf(c.count).Kind() {
	case reflect.Uint64:
		return Uint64BigEndianToBytes(uint64(c.count))
	case reflect.Uint32:
		return Uint32BigEndianToBytes(uint32(c.count))
	case reflect.Uint16:
		return Uint16BigEndianToBytes(uint16(c.count))
	case reflect.Uint8:
		return Uint8BigEndianToBytes(uint8(c.count))
	default:
		loggy.FatalfStack("Type:%s", reflect.TypeOf(c).Kind())
	}
	return nil
}

func (c *Count[T]) String() string {
	c.checkfornil()
	return fmt.Sprintf("%d", c)
}

// -
// goRun()
// -
func (c *CounterStruct[T]) goRun() {
	c.checkfornil()
	defer c.emptych()
	defer close(c.ch)

	var cnt T = 0
	for {
		cnt += 1
		count := &Count[T]{count: cnt}
		select {
		case c.ch <- count:
		case <-c.cx.DoneChan():
			return
		}
	}
}

func (c *CounterStruct[T]) emptych() {
	c.checkfornil()
	for l := len(c.ch); l > 0; l = len(c.ch) {
		if nil == <-c.ch {
			break
		}
	}
}

func (c *CounterStruct[T]) checkfornil() {
	if c == nil {
		loggy.FatalStack("nil method")
	}
}

func (c *Count[T]) checkfornil() {
	if c == nil {
		loggy.FatalStack("nil method")
	}
}
