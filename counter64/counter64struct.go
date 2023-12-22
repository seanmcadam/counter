package counter64

import (
	"encoding/binary"

	"github.com/seanmcadam/counter/countererrors"
	"github.com/seanmcadam/counter/counterint"
	"github.com/seanmcadam/ctx"
	log "github.com/seanmcadam/loggy"
	"github.com/seanmcadam/counter/common"
)

type Counter64Struct struct {
	cx      *ctx.Ctx
	countCh chan *Counter64
}

func New(cx *ctx.Ctx) counterint.CounterStructInt {
	if cx == nil {
		log.Fatal("nil ctx")
	}
	c := &Counter64Struct{
		cx:      cx,
		countCh: make(chan *Counter64),
	}
	go c.goRun()
	return c
}

func (*Counter64Struct) Bits() common.CounterBits {
	return common.BIT64
}

func (*Counter64Struct) ByteToCounter(b []byte) (c counterint.CounterInt, err error) {
	if len(b) != 8 {
		return nil, countererrors.ErrCounterBadParameter(log.Errf("Count data len:%d, :%0x", len(b), b))
	}
	c64 := Counter64(binary.BigEndian.Uint64(b))
	c = &c64
	return c, nil
}

func (c *Counter64Struct) Next() counterint.CounterInt {
	if c == nil {
		log.FatalStack("Nil counter pointer")
	}

	return <-c.countCh
}

// -
// goRun()
// -
func (c *Counter64Struct) goRun() {
	if c == nil {
		log.Fatal()
	}

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
