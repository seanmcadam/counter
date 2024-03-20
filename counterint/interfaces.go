package counterint

import "github.com/seanmcadam/counter/common"

//
// CountInt Used to facilitate a generic counter value of 1,2 4, or 8 bytes and be able to
// reduce to the exact size and send and recieve over a binary structure on the wire
//
type CountInt interface {
	//
	// String returns the string prepresentation of the numeric value
	//
	String() string
	//
	// Uint Returns a converted integer as a uint64
	//
	Uint() uint64
	// 
	// Bits Returns CounterBits (1,2,4,8) based on the underlying integer base
	//
	Bits() common.CounterBits
	//
	// ToByte Returns a byte array coresponding to the length of the integer.
	//
	ToByte() []byte
	//
	// Copy Create a duplicate of the CounterInt
	Copy() CountInt
}

//
// CounterStructInt - used to return a counting object
//
type CounterStructInt interface {
	// 
	// Bits Returns CounterBits (1,2,4,8) based on the underlying integer base
	//
	Bits() common.CounterBits
	//
	// Next Increments the counter and returns a CounterInt version of it
	//
	Next() CountInt
	//
	// ByteToCounter Utility to convert bytes to a CounterInt, based on the bit size
	//
	ByteToCounter([]byte) (CountInt, error)
}
