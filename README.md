# counter
Atomic counter



NewCounterX() creates a interface to a counter with X bits and returns a CounterStructInt{}
The different sizes of counters are for limiting the size of the required byte array when
the data is transmitted over the wire.  Internaly the counters are processed as 64 bits when used.


type Counter counterint.CounterStructInt

You can atomicly get the next counter number by calling Next(), which returns a CounterInt
To find out the depth of counter call Bits() => 8, 16, 32, 64
To convert a byte string to a CounterInt{} use ByteToCounter()

type Count counterint.CountInt

To get a uint64 of the counter value use Uint()
To get the depth of the CounterInt{} use Bits()
To convert to a properly sized []byte value use ToByte()
To get a copy of the CounterInt use Copy()
