package counterint

import "github.com/seanmcadam/counter/common"

type CounterInt interface {
	Uint() uint64
	Bits() common.CounterBits
	ToByte() []byte
	Copy() CounterInt
}

type CounterStructInt interface {
	Bits() common.CounterBits
	Next() CounterInt
	ByteToCounter([]byte) (CounterInt, error)
}
