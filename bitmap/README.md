# bitmap - bitmap operation creation, query, rank/select impl

A bitmap is a `[]uint64`.


```
BenchmarkFromStr32-4            158607060                7.51 ns/op
BenchmarkStringToBytes-4        195268784                6.11 ns/op
BenchmarkRank64_5_bits-4        593165232                2.07 ns/op
BenchmarkRank128_5_bits-4       447032264                2.50 ns/op
BenchmarkRank64_64k_bits-4      894464229                1.33 ns/op
BenchmarkRank128_64k_bits-4     436505224                2.83 ns/op
BenchmarkSelect-4               84540423                14.2 ns/op
BenchmarkSelect2-4              132366105                8.55 ns/op
BenchmarkSelectU64Indexed-4     391063209                3.07 ns/op
```
