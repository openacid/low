package sigbits

// endsOf find every ending bit positions of keys.
//
// Since 0.1.9
func endsOf(keys []string) []int {

	l := len(keys)

	ends := make([]int, l)
	for i := 0; i < l; i++ {
		ends[i] = len(keys[i]) * 8
	}

	return ends
}
