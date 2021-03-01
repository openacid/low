package randext

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeo(t *testing.T) {
	for _, g := range []func(int) int{
		geoByLoop,
		Geo,
	} {
		for _, ratio := range []int{2, 4, 8} {
			testGeo(t, ratio, g)
		}
	}

}

func testGeo(t *testing.T, ratio int, geo func(int) int) {
	name := fmt.Sprintf("%s/%d", GetFunctionName(geo), ratio)

	t.Run(name, func(t *testing.T) {
		ta := require.New(t)

		n := 1000 * 1000
		sample := make([]int, 64)
		for i := 0; i < n; i++ {
			v := geo(ratio)
			sample[v]++
		}

		// fmt.Println(sample[:4])
		for i := 0; i < 4; i++ {
			got := float64(sample[i]) / float64(sample[i+1])
			ta.InDelta(ratio, got, float64(ratio)/5, "%d %d", sample[i], sample[i+1])
		}
	})
}

var OutputGeo int

func BenchmarkGeo(b *testing.B) {
	for _, g := range []func(int) int{
		geoByLoop,
		Geo,
	} {
		for _, ratio := range []int{2, 4, 8} {
			name := fmt.Sprintf("%s/%d", GetFunctionName(g), ratio)
			b.Run(name, func(b *testing.B) {
				var s int
				for i := 0; i < b.N; i++ {
					s += g(ratio)
				}
				OutputGeo = s
			})
		}
	}
}

func Benchmark_rand_Uint64(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += int(rand.Uint64())
	}
	OutputGeo = s
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
