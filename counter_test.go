package counter

import (
	"log"
	"testing"

	"github.com/seanmcadam/counter/common"
	"github.com/seanmcadam/counter/counter16"
	"github.com/seanmcadam/counter/counter32"
	"github.com/seanmcadam/counter/counter64"
	"github.com/seanmcadam/counter/counter8"
	"github.com/seanmcadam/ctx"
)

func TestCounter_compile(t *testing.T) {
}

func TestNewCounters(t *testing.T) {
	bits := []common.CounterBits{BIT64, BIT32, BIT16, BIT8}
	cx := ctx.New()

	for _, b := range bits {
		cs := New(cx, b)
		v := cs.Next()
		log.Printf("Bits:%d", v.Bits())
		log.Printf("cs64:%d", v.Uint())
	}

}

func TestNewCounterX(t *testing.T) {
	cx := ctx.New()

	cs64 := counter64.New(cx)
	v := cs64.Next()
	log.Printf("Bits:%d", v.Bits())
	log.Printf("cs64:%d", v.Uint())

	cs32 := counter32.New(cx)
	v = cs32.Next()
	log.Printf("Bits:%d", v.Bits())
	log.Printf("cs32:%d", v.Uint())

	cs16 := counter16.New(cx)
	v = cs16.Next()
	log.Printf("Bits:%d", v.Bits())
	log.Printf("cs16:%d", v.Uint())

	cs8 := counter8.New(cx)
	v = cs8.Next()
	log.Printf("Bits:%d", v.Bits())
	log.Printf("cs8:%d", v.Uint())

	_ = v
}

func TestNewCounter64(t *testing.T) {
	cx := ctx.New()

	cs64 := counter64.New(cx)
	v := cs64.Next()
	log.Printf("Count:%s\n", string(v.ToByte()))

	cs32 := counter32.New(cx)
	v = cs32.Next()

	cs16 := counter16.New(cx)
	v = cs16.Next()

	cs8 := counter8.New(cx)
	v = cs8.Next()

	_ = v
}

func TestByteToCounter64(t *testing.T) {

	var b []byte = []byte{0, 0, 0, 0, 0, 0, 0, 0}

	cx := ctx.New()
	cs64 := counter64.New(cx)
	c, err := cs64.ByteToCounter(b)
	if err != nil {
		t.Fatalf("Error:%s", err)
	}

	if c.Bits() != 64 {
		t.Fatalf("Bit != 64")
	}
}

func TestByteToCounter32(t *testing.T) {

	var b []byte = []byte{0, 0, 0, 0}

	cx := ctx.New()
	cs32 := counter32.New(cx)
	c, err := cs32.ByteToCounter(b)
	if err != nil {
		t.Fatalf("Error:%s", err)
	}

	if c.Bits() != 32 {
		t.Fatalf("Bit != 32")
	}
}

func TestByteToCounter16(t *testing.T) {

	var b []byte = []byte{0, 0}

	cx := ctx.New()
	cs16 := counter16.New(cx)
	c, err := cs16.ByteToCounter(b)
	if err != nil {
		t.Fatalf("Error:%s", err)
	}

	if c.Bits() != 16 {
		t.Fatalf("Bit != 16")
	}
}

func TestByteToCounter8(t *testing.T) {

	var b []byte = []byte{0}

	cx := ctx.New()
	cs8 := counter8.New(cx)
	c, err := cs8.ByteToCounter(b)
	if err != nil {
		t.Fatalf("Error:%s", err)
	}

	if c.Bits() != 8 {
		t.Fatalf("Bit != 8")
	}
}
