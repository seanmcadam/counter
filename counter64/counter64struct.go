package counter64

import (
	"encoding/binary"

	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/countererrors"
	"github.com/seanmcadam/counter/counterint"
	"github.com/seanmcadam/ctx"
	"github.com/seanmcadam/loggy"
)

type Counter64Struct struct {
	cx      *ctx.Ctx
	countCh chan *Counter64
}

func New(cx *ctx.Ctx) counterint.CounterStructInt {
	if cx == nil {
		loggy.Fatal("nil ctx")
	}
	c := &Counter64Struct{
		cx:      cx,
		countCh: make(chan *Counter64, common.ChanDepth),
	}
	go c.goRun()
	return c
}

func (c *Counter64Struct) Bits() common.CounterBits {
	c.checkfornil()
	return common.BIT64
}

func (c *Counter64Struct) ByteToCounter(b []byte) (ci counterint.CountInt, err error) {
	c.checkfornil()
	if len(b) != 8 {
		return nil, countererrors.ErrCounterBadParameter(loggy.Errf("Count data len:%d, :%0x", len(b), b))
	}
	c64 := Counter64(binary.BigEndian.Uint64(b))
	ci = &c64
	return ci, nil
}

func (c *Counter64Struct) Next() counterint.CountInt {
	c.checkfornil()

	return <-c.countCh
}

// -
// goRun()
// -
func (c *Counter64Struct) goRun() {
	c.checkfornil()
	defer c.emptych()

	var counter Counter64 = 0
	for {
		counter += 1
		c64 := counter
		select {
		case c.countCh <- &c64:
		case <-c.cx.DoneChan():
			return
		}
	}
}

func (c *Counter64Struct) emptych() {
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

func (c *Counter64Struct) checkfornil() {
	if c == nil {
		loggy.FatalStack("nil method")
	}
}
