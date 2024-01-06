package counterint

import "github.com/seanmcadam/counter/common"

type CounterInt interface {
	// CounterInt
	// Used to facilitate a generic counter value of 1,2 4, or 8 bytes and be able to
	// reduce to the exact size and send and recieve over a binary structure on the wire
	String() string
	//
	// Returns a converted integer as a uint64
	//
	Uint() uint64
	// 
	// Returns CounterBits (1,2,4,8) based on the underlying integer base
	//
	Bits() common.CounterBits
	//
	// Returns a byte array coresponding to the length of the integer.
	//
	ToByte() []byte
	//
	// Create a duplicate of the CounterInt
	Copy() CounterInt
}

type CounterStructInt interface {
	// 
	// Returns CounterBits (1,2,4,8) based on the underlying integer base
	//
	Bits() common.CounterBits
	//
	// Increments the counter and returns a CounterInt version of it
	//
	Next() CounterInt
	//
	// Utility to convert bytes to a CounterInt, based on the bit size
	//
	ByteToCounter([]byte) (CounterInt, error)
}
