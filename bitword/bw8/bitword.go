// Package bw8 provides string operations functions
package bw8

var (
	// mask for 1, 2, 4, 8 bit word
	wordMask = []byte{
		// 1, 2, 3, 4, 5, 6, 7, 8
		0, 1, 3, 0, 15, 0, 0, 0, 255,
	}
)

// FromStr split a string into a slice of byte.
// A byte in string is split into 8/`n` `n`-bit words
// Value of every byte is in range [0, 2^n-1].
// `n` must be a one of [1, 2, 4, 8].
//
// Significant bits in a byte is place at left.
// Thus the result byte slice keeps order with the original string.
//
// Since 0.1.2
func FromStr(s string, n int) []byte {
	if wordMask[n] == 0 {
		panic("n must be one of 1, 2, 4, 8")
	}

	mask := wordMask[n]

	// number of words per byte
	m := 8 / n
	lenSrc := len(s)
	words := make([]byte, lenSrc*m)

	for i := 0; i < lenSrc; i++ {
		b := s[i]

		for j := 0; j < m; j++ {
			words[i*m+j] = (b >> uint(8-n*j-n)) & mask
		}
	}
	return words
}

// FromStrs converts a `[]string` to a n-bit word `[][]byte`.
//
// Since 0.1.2
func FromStrs(strs []string, n int) [][]byte {
	rst := make([][]byte, len(strs))
	for i, s := range strs {
		rst[i] = FromStr(s, n)
	}
	return rst
}

// ToStr is the reverse of FromStr.
// It composes a string of which each byte is formed from 8/n words from bs.
//
// Since 0.1.2
func ToStr(bs []byte, n int) string {
	if wordMask[n] == 0 {
		panic("n must be one of 1, 2, 4, 8")
	}

	// number of words per byte
	m := 8 / n
	sz := (len(bs) + m - 1) / m
	strbs := make([]byte, sz)

	var b byte
	for i := 0; i < len(strbs); i++ {
		b = 0
		for j := 0; j < m; j++ {
			if i*m+j < len(bs) {
				b = (b << uint(n)) + bs[i*m+j]
			} else {
				b = b << uint(n)
			}
		}
		strbs[i] = b
	}

	return string(strbs)
}

// ToStrs converts a `[][]byte` back to a `[]string`.
//
// Since 0.1.2
func ToStrs(bytesslice [][]byte, n int) []string {
	rst := make([]string, len(bytesslice))
	for i, s := range bytesslice {
		rst[i] = ToStr(s, n)
	}
	return rst
}

// Get returns i-th n-bit word from a string.
//
// n must be one of 1, 2, 4, 8
//
// Since 0.1.2
func Get(s string, n int, ith int) byte {
	i := n * ith
	end := (i + n - 1) & 7

	word := s[i>>3]
	return (word >> uint(7-end)) & wordMask[n]
}

// FirstDiff returns the index of the first different n-bit word,
// that ge "from" and lt "end".
// If "end" is -1 it means to look up upto end of a or b.
//
// Since 0.1.2
func FirstDiff(a, b string, n int, from, end int) int {
	la, lb := len(a)*8/n, len(b)*8/n

	if end == -1 {
		end = la
	}

	if end > la {
		end = la
	}

	if end > lb {
		end = lb
	}

	for i := from; i < end; i++ {
		if Get(a, n, i) != Get(b, n, i) {
			return i
		}
	}
	return end
}
