package counter32

import (
	"encoding/binary"

	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/countererrors"
	"github.com/seanmcadam/counter/counterint"
	"github.com/seanmcadam/ctx"
	"github.com/seanmcadam/loggy"
)

type Counter32Struct struct {
	cx      *ctx.Ctx
	countCh chan *Counter32
}

func New(cx *ctx.Ctx) counterint.CounterStructInt {
	if cx == nil {
		loggy.Fatal("nil ctx")
	}
	c := &Counter32Struct{
		cx:      cx,
		countCh: make(chan *Counter32, common.ChanDepth),
	}
	go c.goRun()
	return c
}

func (*Counter32Struct) Bits() common.CounterBits {
	return common.BIT32
}

func (c *Counter32Struct) Next() counterint.CountInt {
	c.checkfornil()
	return <-c.countCh
}

func (c *Counter32Struct) ByteToCounter(b []byte) (ci counterint.CountInt, err error) {
	c.checkfornil()
	if len(b) != 4 {
		return nil, countererrors.ErrCounterBadParameter(loggy.Errf("Count data len:%d, :%0x", len(b), b))
	}
	c32 := Counter32(binary.BigEndian.Uint32(b))
	return &c32, nil
}

// -
// goRun()
// -
func (c *Counter32Struct) goRun() {
	c.checkfornil()

	defer c.emptych()

	var counter Counter32 = 0
	for {
		counter += 1
		c32 := counter
		select {
		case c.countCh <- &c32:
		case <-c.cx.DoneChan():
			return
		}
	}
}

func (c *Counter32Struct) emptych() {
	c.checkfornil()
	if c == nil {
		return
	}

	for {
		select {
		case <-c.countCh:
		default:
			close(c.countCh)
			return
		}
	}
}

func (c *Counter32Struct) checkfornil() {
	if c == nil {
		loggy.FatalStack("nil method")
	}
}