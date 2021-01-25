package sigbits

// ShardByPrefix splits a sorted string slice into shards each of which has at most maxSize elts.
//
// The rule is that every shard has a unique prefix. All prefixes are still a
// sorted string slice.
//
// It returns a slice of prefix lengths in bytes of each shard, and a slice of
// the starting index of each shard.
//
// Since 0.1.20
func ShardByPrefix(keys []string, maxSize int32) ([]int32, []int32) {

	firstDiffs := FirstDiffBits(keys)

	n := int32(len(firstDiffs) + 1)

	prefixes := make([]int32, 0, len(keys)/32)
	keyCnts := make([]int32, 0, len(keys)/32)
	keyCnts = append(keyCnts, 0)

	var dfs func(s, e int32)

	dfs = func(s, e int32) {
		if e-s <= maxSize {

			// find the longest prefix:
			min := int32(len(keys[s]))
			for i := s; i < e-1; i++ {
				if min > firstDiffs[i]>>3 {
					min = firstDiffs[i] >> 3
				}
			}

			prefixes = append(prefixes, min)
			keyCnts = append(keyCnts, e)
			return
		}

		// range is too large, split it

		longest := int32(len(keys[s]))

		// endsAt collects where to split current range
		endsAt := make([]int32, 0, 256)

		for i := s; i < e-1; i++ {
			prefixLen := firstDiffs[i] >> 3

			if prefixLen < longest {
				longest = prefixLen
				endsAt = endsAt[0:0]
				endsAt = append(endsAt, i+1)
			} else if prefixLen == longest {
				endsAt = append(endsAt, i+1)
			}
		}

		endsAt = append(endsAt, e)

		for i := 0; i < len(endsAt); i++ {
			end := endsAt[i]
			dfs(s, end)
			s = end
		}
	}

	dfs(0, n)

	return prefixes, keyCnts
}
