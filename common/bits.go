package common

import (
	log "github.com/seanmcadam/loggy"
)

type CounterBits uint8

const (
	BIT8  = 8
	BIT16 = 16
	BIT32 = 32
	BIT64 = 64
)

func (c CounterBits) Number() uint8 {
	switch c {
	case BIT8:
		return uint8(8)
	case BIT16:
		return uint8(16)
	case BIT32:
		return uint8(32)
	case BIT64:
		return uint8(64)
	default:
		log.Fatal("c:%d", uint8(c))
	}
	return 0
}

func (c CounterBits) String() string {
	switch c {
	case BIT8:
		return "8"
	case BIT16:
		return "16"
	case BIT32:
		return "32"
	case BIT64:
		return "64"
	default:
		log.Fatal("c:%d", uint8(c))
	}
	return " _unknown_ "
}
