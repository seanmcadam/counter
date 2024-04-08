package counter16

import (
	"encoding/binary"

	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/countererrors"
	"github.com/seanmcadam/counter/counterint"
	"github.com/seanmcadam/ctx"
	"github.com/seanmcadam/loggy"
)

type Counter16Struct struct {
	cx      *ctx.Ctx
	countCh chan *Counter16
}

func New(cx *ctx.Ctx) counterint.CounterStructInt {
	if cx == nil {
		loggy.Fatal("nil ctx")
	}
	c := &Counter16Struct{
		cx:      cx,
		countCh: make(chan *Counter16, common.ChanDepth),
	}
	go c.goRun()
	return c
}

func (*Counter16Struct) Bits() common.CounterBits {
	return common.BIT16
}

func (*Counter16Struct) ByteToCounter(b []byte) (c counterint.CountInt, err error) {
	if len(b) != 2 {
		return nil, countererrors.ErrCounterBadParameter(loggy.Errf("Count data len:%d, :%0x", len(b), b))
	}
	c16 := Counter16(binary.BigEndian.Uint16(b))
	return &c16, nil
}

func (c *Counter16Struct) Next() counterint.CountInt {
	if c == nil {
		loggy.FatalStack("Nil counter pointer")
	}

	return <-c.countCh
}

// -
// goRun()
// -
func (c *Counter16Struct) goRun() {
	if c == nil {
		loggy.Fatal()
	}

	defer c.emptych()

	var counter Counter16 = 0
	for {
		counter += 1
		c16 := counter
		select {
		case c.countCh <- &c16:
		case <-c.cx.DoneChan():
			return
		}
	}
}

func (c *Counter16Struct) emptych() {
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
