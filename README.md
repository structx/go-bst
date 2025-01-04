## BST 

Simple binary search tree. Implemented to be used with other `structx` products.

Utilizes a string as the key and stores a payload. Tree can be traversed and searched.

## Benchmark

```
goos: linux
goarch: amd64
pkg: github.com/structx/go-bst
cpu: 13th Gen Intel(R) Core(TM) i5-1340P
=== RUN   BenchmarkInsert
BenchmarkInsert
BenchmarkInsert-16         23011             72093 ns/op              64 B/op          2 allocs/op
PASS
ok      github.com/structx/go-bst       3.262s
```