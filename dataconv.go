package counter

import "github.com/seanmcadam/loggy"

func Uint8BigEndianToBytes(u uint8) []byte {
	if u > 0xFF {
		loggy.FatalStack("")
	}
	return []byte{
		byte(u),
	}
}
func Uint16BigEndianToBytes(u uint16) []byte {
	if u > 0xFFFF {
		loggy.FatalStack("")
	}
	return []byte{
		byte(u >> 8),
		byte(u & 0xFF),
	}

}
func Uint32BigEndianToBytes(u uint32) []byte {
	if u > 0xFFFFFFFF {
		loggy.FatalStack("")
	}
	return []byte{
		byte(u >> 24),
		byte(u >> 16 & 0xFF),
		byte(u >> 8 & 0xFF),
		byte(u & 0xFF),
	}
}
func Uint64BigEndianToBytes(u uint64) []byte {
	if u > 0xFFFFFFFFFFFFFFFF {
		loggy.FatalStack("")
	}
	return []byte{
		byte(u >> 56),
		byte(u >> 48 & 0xFF),
		byte(u >> 40 & 0xFF),
		byte(u >> 32 & 0xFF),
		byte(u >> 24 & 0xFF),
		byte(u >> 16 & 0xFF),
		byte(u >> 8 & 0xFF),
		byte(u & 0xFF),
	}

}
