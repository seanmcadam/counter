package counter8

import (
	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/countererrors"
	"github.com/seanmcadam/counter/counterint"
	"github.com/seanmcadam/ctx"
	"github.com/seanmcadam/loggy"
)

type Counter8Struct struct {
	cx      *ctx.Ctx
	countCh chan *Counter8
}

func New(cx *ctx.Ctx) counterint.CounterStructInt {
	if cx == nil {
		loggy.Fatal("nil ctx")
	}
	c := &Counter8Struct{
		cx:      cx,
		countCh: make(chan *Counter8, common.ChanDepth),
	}
	go c.goRun()
	return c
}

func (*Counter8Struct) Bits() common.CounterBits {
	return common.BIT8
}

func (c *Counter8Struct) ByteToCounter(b []byte) (ci counterint.CountInt, err error) {
	c.checkfornil()
	if len(b) != 1 {
		return nil, countererrors.ErrCounterBadParameter(loggy.Errf("Count data len:%d, :%0x", len(b), b))
	}
	c8 := Counter8(b[0])
	return &c8, nil
}

func (c *Counter8Struct) Next() counterint.CountInt {
	c.checkfornil()
	if c == nil {
		loggy.FatalStack("Nil counter pointer")
	}

	return <-c.countCh
}

// -
// goRun()
// -
func (c *Counter8Struct) goRun() {
	c.checkfornil()

	defer c.emptych()

	var counter Counter8 = 0
	for {
		counter += 1
		c8 := counter
		select {
		case c.countCh <- &c8:
		case <-c.cx.DoneChan():
			return
		}
	}
}

func (c *Counter8Struct) emptych() {
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

func (c *Counter8Struct) checkfornil() {
	if c == nil {
		loggy.FatalStack("nil method")
	}
}
