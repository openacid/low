# bitmap - bitmap operation creation, query, rank/select impl

A bitmap is a `[]uint64`.


```
BenchmarkFromStr32-4                  147230166                8.07 ns/op
BenchmarkStringToBytes-4              201391240                5.95 ns/op
BenchmarkNextOne/FoundInCurrentWord-4 334609323                3.56 ns/op
BenchmarkNextOne/EnumerateAllBits-4   294496100                4.03 ns/op
BenchmarkPrevOne/FoundInCurrentWord-4 334022793                3.59 ns/op
BenchmarkPrevOne/EnumerateAllBits-4   286702246                4.12 ns/op
BenchmarkRank64_5_bits-4              672754074                1.77 ns/op
BenchmarkRank128_5_bits-4             486679581                2.49 ns/op
BenchmarkRank64_64k_bits-4            881509602                1.35 ns/op
BenchmarkRank128_64k_bits-4           424685930                2.84 ns/op
BenchmarkSelect-4                     70999674                16.1 ns/op
BenchmarkSelect32-4                   141631495                8.54 ns/op
BenchmarkSelect32R64-4                153702387                7.85 ns/op
BenchmarkSelectU64Indexed-4           365140676                3.29 ns/op
```
