package counter

import (
	"encoding/binary"
	"fmt"

	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/counter16"
	"github.com/seanmcadam/counter/counter32"
	"github.com/seanmcadam/counter/counter64"
	"github.com/seanmcadam/counter/counter8"
	"github.com/seanmcadam/counter/countererrors"
	"github.com/seanmcadam/counter/counterint"
	"github.com/seanmcadam/ctx"
	"github.com/seanmcadam/loggy"
	log "github.com/seanmcadam/loggy"
)

type Counter counterint.CounterStructInt
type Count counterint.CountInt

const BIT8 = common.BIT8
const BIT16 = common.BIT16
const BIT32 = common.BIT32
const BIT64 = common.BIT64

func New(cx *ctx.Ctx, b common.CounterBits) Counter {
	switch b {
	case BIT8:
		return counter8.New(cx)
	case BIT16:
		return counter16.New(cx)
	case BIT32:
		return counter32.New(cx)
	case BIT64:
		return counter64.New(cx)
	default:
		log.Fatalf("Unknown BIT value %d", b)
	}
	return nil
}

func NewCount(c interface{}) Count {
	switch val := c.(type) {
	case int8:
		if val < 0 {
			loggy.Panicf("negative integer")
		}
		return counter8.NewCount(uint8(val))
	case int16:
		if val < 0 {
			loggy.Panicf("negative integer")
		}
		return counter16.NewCount(uint16(val))
	case int32:
		if val < 0 {
			loggy.Panicf("negative integer")
		}
		return counter32.NewCount(uint32(val))
	case int64:
		if val < 0 {
			loggy.Panicf("negative integer")
		}
		return counter64.NewCount(uint64(val))
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

func ByteToCount(b []byte) (c Count, err error) {
	x := len(b)
	switch x {
	case 1:
		c8 := counter8.Counter8(b[0])
		return &c8, nil
	case 2:
		c16 := counter16.Counter16(binary.BigEndian.Uint16(b))
		return &c16, nil
	case 4:
		c32 := counter32.Counter32(binary.BigEndian.Uint32(b))
		return &c32, nil
	case 8:
		c64 := counter64.Counter64(binary.BigEndian.Uint64(b))
		return &c64, nil
	}

	return nil, countererrors.ErrCounterBadParameter(fmt.Errorf("ByteToCount() Byte len(b): %d", x))
}
