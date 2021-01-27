// Package util provides most used math functions that golang does not provided.
//
// Since 0.1.21
package util

func MinI(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxI(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinI8(a, b int8) int8 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxI8(a, b int8) int8 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinI16(a, b int16) int16 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxI16(a, b int16) int16 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinI32(a, b int32) int32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxI32(a, b int32) int32 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinI64(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxI64(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

// uint:

func MinU(a, b uint) uint {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxU(a, b uint) uint {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinU8(a, b uint8) uint8 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxU8(a, b uint8) uint8 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinU16(a, b uint16) uint16 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxU16(a, b uint16) uint16 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinU32(a, b uint32) uint32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxU32(a, b uint32) uint32 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinU64(a, b uint64) uint64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxU64(a, b uint64) uint64 {
	if a > b {
		return a
	} else {
		return b
	}
}

// clap

func ClapI(n, min, max int) int {
	if n < min {
		n = min
	}
	if n > max {
		n = max
	}
	return n
}

func ClapI8(n, min, max int8) int8 {
	if n < min {
		n = min
	}
	if n > max {
		n = max
	}
	return n
}

func ClapI16(n, min, max int16) int16 {
	if n < min {
		n = min
	}
	if n > max {
		n = max
	}
	return n
}

func ClapI32(n, min, max int32) int32 {
	if n < min {
		n = min
	}
	if n > max {
		n = max
	}
	return n
}

func ClapI64(n, min, max int64) int64 {
	if n < min {
		n = min
	}
	if n > max {
		n = max
	}
	return n
}

func ClapU(n, min, max uint) uint {
	if n < min {
		n = min
	}
	if n > max {
		n = max
	}
	return n
}

func ClapU8(n, min, max uint8) uint8 {
	if n < min {
		n = min
	}
	if n > max {
		n = max
	}
	return n
}

func ClapU16(n, min, max uint16) uint16 {
	if n < min {
		n = min
	}
	if n > max {
		n = max
	}
	return n
}

func ClapU32(n, min, max uint32) uint32 {
	if n < min {
		n = min
	}
	if n > max {
		n = max
	}
	return n
}

func ClapU64(n, min, max uint64) uint64 {
	if n < min {
		n = min
	}
	if n > max {
		n = max
	}
	return n
}
