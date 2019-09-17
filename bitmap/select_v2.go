// These functions are experimental.

package bitmap

// // for test and evaluate new implementation
// type bitmap struct {
//     words []uint64
//     ranks []int32
//     index []int32
//     i, j  int32
// }

// // slower, 9 ns.
// // Calling an assembly function still cost a lot.
// func select2withPDEP(bm *bitmap, i int32) {

//     wordI := bm.index[i>>5]
//     for ; bm.ranks[wordI+1] <= i; wordI++ {
//     }

//     w := bm.words[wordI]
//     a := int32(selectU64WithPDEP(w, uint64(i-bm.ranks[wordI])))

//     i++
//     wp := wordI << 6
//     bm.i = wp + a

//     if bm.ranks[wordI+1] > i {
//         // a&63 to elide boundary check
//         b := int32(bits.TrailingZeros64(w & RMaskUpto[a&63]))

//         bm.j = b + wp
//         return
//     }

//     wordI++

//     for ; bm.ranks[wordI+1] <= i; wordI++ {
//     }

//     bm.j = int32(bits.TrailingZeros64(bm.words[wordI])) + wordI<<6
// }

// // Use a rank index and a select index.
// // This is the best for now. about 8 ns
// func select2withBsearch(bm *bitmap, i int32) (int32, int32) {

//     a := int32(0)
//     l := int32(len(bm.words))

//     wordI := bm.index[i>>5]
//     for ; bm.ranks[wordI+1] <= i; wordI++ {
//     }

//     w := bm.words[wordI]
//     base := wordI << 6
//     findIth := int(i - bm.ranks[wordI])

//     offset := int32(0)

//     ones := bits.OnesCount32(uint32(w))
//     if ones <= findIth {
//         findIth -= ones
//         offset |= 32
//         w >>= 32
//     }

//     ones = bits.OnesCount16(uint16(w))
//     if ones <= findIth {
//         findIth -= ones
//         offset |= 16
//         w >>= 16
//     }

//     ones = bits.OnesCount8(uint8(w))
//     if ones <= findIth {
//         a = int32(tbl2[(w>>5)&(0x7f8)|uint64(findIth-ones)]) + offset + 8
//     } else {
//         a = int32(tbl2[(w&0xff)<<3|uint64(findIth)]) + offset
//     }

//     a += base

//     // "& 63" elides boundary check
//     w &= RMaskUpto[a&63]
//     if w != 0 {
//         return a, base + int32(bits.TrailingZeros64(w))
//     }

//     wordI++
//     for ; wordI < l; wordI++ {
//         w = bm.words[wordI]
//         if w != 0 {
//             return a, wordI<<6 + int32(bits.TrailingZeros64(w))
//         }
//     }
//     return a, l << 6
// }
