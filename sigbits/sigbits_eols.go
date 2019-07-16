package sigbits

import "sort"

// EOLs returns a list of index of the end of keys relative to "frombit".
// If a key ends after "tobit", it uses "tobit" as end.
// If "tobit" is -1 then it does check it.
//
//
// Since 0.1.9
func (sb *SigBits) EOLs(keyStart, keyEnd, frombit, tobit int32, dedup bool) []int32 {

	l := keyEnd - keyStart

	eols := make([]int, l)

	for i := keyStart; i < keyEnd; i++ {
		v := sb.ends[i]
		if tobit != -1 && v > int(tobit) {
			v = int(tobit)
		}
		v -= int(frombit)
		eols[i-keyStart] = v
	}

	sort.Ints(eols)

	rst := make([]int32, 0, l)
	for i, v := range eols {
		if !dedup || i == 0 || (v > eols[i-1]) {
			rst = append(rst, int32(v))
		}
	}
	return rst
}
