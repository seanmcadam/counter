package counter

import (
	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/counter16"
	"github.com/seanmcadam/counter/counter32"
	"github.com/seanmcadam/counter/counter64"
	"github.com/seanmcadam/counter/counter8"
	"github.com/seanmcadam/counter/counterint"
	"github.com/seanmcadam/ctx"
	log "github.com/seanmcadam/loggy"
)

const BIT8 = common.BIT8
const BIT16 = common.BIT16
const BIT32 = common.BIT32
const BIT64 = common.BIT64

func New(cx *ctx.Ctx, b common.CounterBits) counterint.CounterStructInt {
	switch b {
	case BIT8:
		return NewCounter8(cx)
	case BIT16:
		return NewCounter16(cx)
	case BIT32:
		return NewCounter32(cx)
	case BIT64:
		return NewCounter64(cx)
	default:
		log.Fatalf("Unknown BIT value %d", b)
	}
	return nil
}

func NewCounter64(cx *ctx.Ctx) counterint.CounterStructInt {
	cs := counter64.New(cx)
	return cs
}

func NewCounter32(cx *ctx.Ctx) counterint.CounterStructInt {
	cs := counter32.New(cx)
	return cs
}

func NewCounter16(cx *ctx.Ctx) counterint.CounterStructInt {
	cs := counter16.New(cx)
	return cs
}

func NewCounter8(cx *ctx.Ctx) counterint.CounterStructInt {
	cs := counter8.New(cx)
	return cs
}

func NewCount(c interface{}) counterint.CounterInt {
	switch val := c.(type) {
	case uint8:
		return counter8.NewCount(val)
	case uint16:
		return counter16.NewCount(val)
	case uint32:
		return counter32.NewCount(val)
	case uint64:
		return counter64.NewCount(val)
	default:
		log.Fatalf("NewCount() type:%v", val)
	}

	return nil
}
