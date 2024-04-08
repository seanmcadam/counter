# counter
Atomic counter

New[uintxx]
New8
New16
New32
New64

New[uint] or NewNN creates a counter with X bits and returns a CounterStuct{}
The different sizes of counters are for limiting the size of the required byte array when
the data is transmitted over the wire.  Counts are returned as a struct ptr, and can take 
advantage of some methods.  .Uint() will return the count as the correct sized uint

To get a uintNN of the counter value use Uint()
To convert to a properly sized Big Endian []byte value use ToBEByte()
To get a copy of the CounterInt use Copy()

You can atomicly get the next counter number by calling Next(), which returns a CounterInt

The counterStruct primes the channel with hard coded number of values so Next() is fast.

