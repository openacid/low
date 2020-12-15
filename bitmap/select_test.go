package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndexSelect64(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input []uint64
		want  []int32
	}{
		{nil, []int32{}},
		{[]uint64{}, []int32{}},
		{[]uint64{0}, []int32{}},
		{[]uint64{1}, []int32{0}},
		{[]uint64{2}, []int32{1}},
		{[]uint64{3}, []int32{0}},
		{[]uint64{4, 0}, []int32{2}},
		{[]uint64{0xffffffff, 0xffffffff}, []int32{0, 64}},
		{[]uint64{0xffffffff, 0xffffffff, 1}, []int32{0, 64, 128}},

		{[]uint64{0xffffffffffffffff, 0xffffffffffffffff, 1}, []int32{0, 32, 64, 96, 128}},
		{[]uint64{0, 0xffffffffffffffff, 1}, []int32{64, 96, 128}},
	}

	for i, c := range cases {
		got := IndexSelect32(c.input)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}

// func TestSelect_panic(t *testing.T) {

//     ta := require.New(t)

//     cases := []struct {
//         input  []uint64
//         inputI int32
//     }{
//         // {nil, 0},
//         {nil, 1},
//         {[]uint64{}, 0},
//         {[]uint64{}, 1},
//         {[]uint64{0}, 0},
//         {[]uint64{0}, 1},
//         {[]uint64{1}, 1},
//         {[]uint64{1}, 2},

//         {[]uint64{2}, 1},
//         {[]uint64{2}, 2},
//         {[]uint64{3}, 2},
//         {[]uint64{0xffffffff, 0xffffffff, 1}, 65},
//     }

//     for i, c := range cases {
//         idx := IndexSelect64(c.input)
//         _ = i
//         _ = ta
//         _ = Select(c.input, idx, c.inputI)
//         // ta.Panics(func() { Select(c.input, idx, c.inputI) }, "%d-th: case: %+v", i+1, c)
//     }
// }

func TestSelect(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input []uint64
	}{
		{nil},
		{[]uint64{}},
		{[]uint64{0}},
		{[]uint64{1}},
		{[]uint64{2}},
		{[]uint64{3}},
		{[]uint64{4, 0}},
		{[]uint64{0xffffffff, 0xffffffff}},
		{[]uint64{0xffffffff, 0xffffffff, 1}},
		{[]uint64{0x6668}}, // 000101100110011
	}

	for i, c := range cases {
		idx := IndexSelect32(c.input)

		ta.Equal(int32(-1), select32single(c.input, idx, -1))

		nth := int32(-1)
		for j := 0; j < len(c.input)*64; j++ {

			if c.input[j>>6]&(1<<uint(j&63)) != 0 {
				nth++
				got := select32single(c.input, idx, nth)
				ta.Equal(int32(j), got, "%d-th: case: %+v, select: %d", i+1, c, nth)
			}
		}

		// test select more "1" than there are
		nth++
		ta.Equal(int32(len(c.input)*64), select32single(c.input, idx, nth))
	}
}

func TestSelect32(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input []uint64
	}{
		{nil},
		{[]uint64{}},
		{[]uint64{0}},
		{[]uint64{1}},
		{[]uint64{2}},
		{[]uint64{3}},
		{[]uint64{4, 0}},
		{[]uint64{0xf, 0xf}},
		{[]uint64{0xf, 0, 0xf}},
		{[]uint64{0xfffffffffffffff0}},
		{[]uint64{0xffffffffffffffff}},
		{[]uint64{0xffffffff, 0xffffffff}},
		{[]uint64{0xffffffff, 0xffffffff, 1}},
		{[]uint64{0x6668}}, // 000101100110011
	}

	for i, c := range cases {
		idx := IndexSelect32(c.input)

		all := make([]int32, 0)

		ta.Equal(int32(-1), select32single(c.input, idx, -1))

		nth := int32(-1)
		for j := 0; j < len(c.input)*64; j++ {

			if c.input[j>>6]&(1<<uint(j&63)) != 0 {
				nth++
				got := select32single(c.input, idx, nth)
				ta.Equal(int32(j), got, "%d-th: case: %+v, select: %d", i+1, c, nth)
				all = append(all, int32(j))
			}
		}

		// test select more "1" than there are
		nth++
		ta.Equal(int32(len(c.input)*64), select32single(c.input, idx, nth))
		all = append(all, int32(len(c.input)*64))

		// a, b := Select32(c.input, idx, -1)
		// ta.Equal(int32(-1), a)
		// ta.Equal(all[0], b)

		for j := 0; j < len(all)-1; j++ {
			a, b := Select32(c.input, idx, int32(j))
			ta.Equal(all[j], a, "select2-1: %d, case: %+v", j, c)
			ta.Equal(all[j+1], b, "select2-2: %d, case: %+v", j, c)
		}

	}
}

func TestSelect32R64(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input []uint64
	}{
		{nil},
		{[]uint64{}},
		{[]uint64{0}},
		{[]uint64{1}},
		{[]uint64{2}},
		{[]uint64{3}},
		{[]uint64{4, 0}},
		{[]uint64{0xf, 0xf}},
		{[]uint64{0xf, 0, 0xf}},
		{[]uint64{0xfffffffffffffff0}},
		{[]uint64{0xffffffffffffffff}},
		{[]uint64{0xffffffff, 0xffffffff}},
		{[]uint64{0xffffffff, 0xffffffff, 1}},
		{[]uint64{0x6668}}, // 000101100110011
	}

	for _, c := range cases {

		sidx, ridx := IndexSelect32R64(c.input)

		all := ToArray(c.input)

		for j := 0; j < len(all)-1; j++ {
			a, b := Select32R64(c.input, sidx, ridx, int32(j))
			ta.Equal(all[j], a, "select2-1: %d, case: %+v", j, c)
			ta.Equal(all[j+1], b, "select2-2: %d, case: %+v", j, c)
		}

	}
}

func TestSelect32_panic(t *testing.T) {

	ta := require.New(t)

	ta.Panics(func() { Select32([]uint64{0}, []int32{0}, -1) })
	ta.Panics(func() { Select32([]uint64{0}, []int32{0}, 32) })
}

func TestIndexSelectU64(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input uint64
		want  uint64
	}{
		{0x00, 0x8080808080808080},
		{0x01, 0x8181818181818181},
		{0x02, 0x8181818181818181},
		{0x03, 0x8282828282828282},
		{0x33, 0x8484848484848484},
		{0x3300, 0x8484848484848480},
		{0xff, 0x8888888888888888},
		{0xff00, 0x8888888888888880},
		{0xff0000, 0x8888888888888080},
		{0xffffff, 0x9898989898989088},
		{0xffffffffffffffff, 0x8080808080808080 | 0x4038302820181008},
	}

	for i, c := range cases {
		got := indexSelectU64(c.input)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}

func TestSelectU64Indexed(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input uint64
	}{
		{0},
		{1},
		{2},
		{3},
		{4},
		{0xffffffff},
		{0xffffffff},
		{0x6668}, // 000101100110011
	}

	for i, c := range cases {

		idx := indexSelectU64(c.input)

		all := make([]int32, 0)

		nth := int32(-1)
		for j := 0; j < 64; j++ {

			if c.input&(1<<uint(j)) != 0 {
				nth++
				got, gotn := selectU64Indexed(c.input, idx, uint64(nth))
				ta.Equal(int32(j), got, "%d-th: case: %+v, select: %d", i+1, c, nth)
				ta.Equal(0, gotn, "%d-th: case: %+v, select: %d", i+1, c, nth)
				all = append(all, int32(j))
			}
		}

		// // test select more "1" than there are
		// nth++
		// got, gotn := selectU64Indexed(c.input, idx, uint64(nth))
		// ta.Equal(int32(64), got)
		// ta.Equal(bits.OnesCount64(c.input), gotn)

		all = append(all, int32(64))
		_ = all

		// for j := 0; j < len(all)-1; j++ {
		//     a, b := Select32(c.input, idx, int32(j))
		//     ta.Equal(all[j], a, "select2-1: %d, case: %+v", j, c)
		//     ta.Equal(all[j+1], b, "select2-2: %d, case: %+v", j, c)
		// }

	}
}
