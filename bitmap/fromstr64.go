package bitmap

// FromStr64 returns a bit array from string and put them in the least
// significant "to-from" bits of a uint64.
// "to-from" must be smaller than or equal 32.
//
// It returns actual number of bits used from the string, and a uint64.
//
// Since 0.1.9
func FromStr64(s string, frombit, tobit int32) (int32, uint64) {

	size := tobit - frombit

	frombyte := frombit >> 3
	blen := int32(len(s)<<3) - frombit

	var b uint64

	if blen > size {
		blen = size
	}

	if blen <= 0 {
		return 0, 0
	}

	firstI := frombit & 7
	spanSize := size + firstI

	if firstI != 0 && spanSize > 64 {

		b = get64Bits(s[frombyte+1:]) >> uint(72-spanSize)
		b |= (uint64(s[frombyte]) << uint(spanSize-8))

	} else {
		b = get64Bits(s[frombyte:]) >> uint(64-spanSize)
	}

	b &= Mask[size]
	return blen, b

}

// get64Bits converts a string to a uint64,
// in big endian.
// Less than 8 byte string will be filled with trailing 0.
// More than 8 bytes will be ignored.
//
// Since 0.1.9
func get64Bits(s string) uint64 {

	if len(s) >= 8 {

		return ((uint64(s[0]) << 56) +
			(uint64(s[1]) << 48) +
			(uint64(s[2]) << 40) +
			(uint64(s[3]) << 32) +
			(uint64(s[4]) << 24) +
			(uint64(s[5]) << 16) +
			(uint64(s[6]) << 8) +
			uint64(s[7]))

	} else {

		bs := make([]byte, 8)
		copy(bs, s)
		return ((uint64(bs[0]) << 56) +
			(uint64(bs[1]) << 48) +
			(uint64(bs[2]) << 40) +
			(uint64(bs[3]) << 32) +
			(uint64(bs[4]) << 24) +
			(uint64(bs[5]) << 16) +
			(uint64(bs[6]) << 8) +
			uint64(bs[7]))
	}
}
