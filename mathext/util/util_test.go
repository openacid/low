package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinMaxClap(t *testing.T) {

	ta := require.New(t)

	ta.Equal(int(1), MinI(1, 2))
	ta.Equal(int8(1), MinI8(1, 2))
	ta.Equal(int16(1), MinI16(1, 2))
	ta.Equal(int32(1), MinI32(1, 2))
	ta.Equal(int64(1), MinI64(1, 2))

	ta.Equal(uint(1), MinU(1, 2))
	ta.Equal(uint8(1), MinU8(1, 2))
	ta.Equal(uint16(1), MinU16(1, 2))
	ta.Equal(uint32(1), MinU32(1, 2))
	ta.Equal(uint64(1), MinU64(1, 2))

	ta.Equal(int(2), MaxI(1, 2))
	ta.Equal(int8(2), MaxI8(1, 2))
	ta.Equal(int16(2), MaxI16(1, 2))
	ta.Equal(int32(2), MaxI32(1, 2))
	ta.Equal(int64(2), MaxI64(1, 2))

	ta.Equal(uint(2), MaxU(1, 2))
	ta.Equal(uint8(2), MaxU8(1, 2))
	ta.Equal(uint16(2), MaxU16(1, 2))
	ta.Equal(uint32(2), MaxU32(1, 2))
	ta.Equal(uint64(2), MaxU64(1, 2))

	ta.Equal(int(2), ClapI(1, 2, 3))
	ta.Equal(int(3), ClapI(4, 2, 3))

	ta.Equal(int8(2), ClapI8(1, 2, 3))
	ta.Equal(int8(3), ClapI8(4, 2, 3))

	ta.Equal(int16(2), ClapI16(1, 2, 3))
	ta.Equal(int16(3), ClapI16(4, 2, 3))

	ta.Equal(int32(2), ClapI32(1, 2, 3))
	ta.Equal(int32(3), ClapI32(4, 2, 3))

	ta.Equal(int64(2), ClapI64(1, 2, 3))
	ta.Equal(int64(3), ClapI64(4, 2, 3))

	ta.Equal(uint(2), ClapU(1, 2, 3))
	ta.Equal(uint(3), ClapU(4, 2, 3))

	ta.Equal(uint8(2), ClapU8(1, 2, 3))
	ta.Equal(uint8(3), ClapU8(4, 2, 3))

	ta.Equal(uint16(2), ClapU16(1, 2, 3))
	ta.Equal(uint16(3), ClapU16(4, 2, 3))

	ta.Equal(uint32(2), ClapU32(1, 2, 3))
	ta.Equal(uint32(3), ClapU32(4, 2, 3))

	ta.Equal(uint64(2), ClapU64(1, 2, 3))
	ta.Equal(uint64(3), ClapU64(4, 2, 3))

}
