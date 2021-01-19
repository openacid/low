package mathext

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZipf(t *testing.T) {

	ta := require.New(t)

	n := 100

	a := float64(1)
	b := float64(n)
	s := float64(1.5)

	got := make([]int, n)
	z := NewZipf(a, b, s)

	want := []int{0, 3255, 1442, 859, 587, 433, 336, 271, 225, 190, 163, 143,
		126, 112, 101, 91, 83, 76, 70, 64, 60, 56, 52, 49, 45, 44, 40, 39, 36, 35,
		33, 31, 30, 29, 27, 27, 25, 24, 23, 23, 21, 21, 20, 20, 18, 18, 18, 17, 16,
		16, 16, 15, 14, 15, 13, 14, 13, 13, 12, 12, 12, 11, 12, 11, 10, 11, 10, 10,
		10, 10, 9, 9, 9, 9, 9, 8, 8, 9, 8, 7, 8, 8, 7, 7, 7, 7, 7, 7, 7, 6, 7, 6, 6,
		7, 6, 6, 5, 6, 6, 6}

	sampleCnt := float64(10000)
	for u := float64(0); u < 1; u += 1 / sampleCnt {
		x := int(z.Float64(float64(u)))
		got[x]++
	}

	ta.Equal(want, got)
}

func TestOfficalZipf(t *testing.T) {

	// Not to run. Only when I want.
	t.Skip()

	start := 1
	end := 100
	n := end
	s := float64(1.5)

	ta := require.New(t)
	_ = ta

	sample := make([]int, n)
	r := rand.New(rand.NewSource(44))

	sampleCnt := float64(10000)
	rz := rand.NewZipf(r, s, float64(start), uint64(end-start))
	for i := 0; i < int(sampleCnt); i++ {
		x := rz.Uint64()
		sample[x]++
	}
}

var Output float64

func BenchmarkZipf(b *testing.B) {
	start := 10
	end := 10000
	a := float64(start)
	bb := float64(end)
	s := float64(1.5)

	z := NewZipf(a, bb, s)
	ss := float64(0)
	for i := 0; i < b.N; i++ {
		v := z.Float64(float64(i) / float64(b.N))
		ss += v
	}

	Output = ss
}

func BenchmarkOfficialZipf(b *testing.B) {

	start := 10
	end := 10000
	s := float64(1.5)
	r := rand.New(rand.NewSource(44))

	rz := rand.NewZipf(r, s, float64(start), uint64(end-start))
	ss := uint64(0)
	for i := 0; i < b.N; i++ {
		x := rz.Uint64()
		ss += x
	}

	Output = float64(ss)
}
