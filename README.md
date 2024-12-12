# Go Benchmarks: Performance Comparisons

This repository contains benchmarks for various Go operations, starting with the performance comparison of appending integers to slices.

## Benchmark 1: Array Append

This benchmark compares two methods of appending integers to slices:
1. **Dynamic Slice Appending**: Appending to a slice that grows dynamically.
2. **Preallocated Slice Assignment**: Assigning values to a preallocated slice.

### Results

The benchmark was executed on the following system:
- **OS**: macOS (Darwin)
- **Architecture**: ARM64
- **Go Version**: Compatible with `go test -bench`.

| Benchmark Name                      | Iterations   | Time per Op (ns/op) | Memory per Op (B/op) | Allocations per Op |
|-------------------------------------|--------------|----------------------|-----------------------|---------------------|
| `AppendToDynamicSlice`              | 251,445,637  | 9.260               | 44                    | 0                   |
| `AssignToPreallocatedSlice`         | 1,000,000,000| 1.466               | 0                     | 0                   |


```
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^BenchmarkArrayAppend$ github.com/insomnius/benchmark

goos: darwin
goarch: arm64
pkg: github.com/insomnius/benchmark
BenchmarkArrayAppend/AppendToDynamicSlice-8         	251445637	         9.260 ns/op	      44 B/op	       0 allocs/op
BenchmarkArrayAppend/AssignToPreallocatedSlice-8    	1000000000	         1.466 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/insomnius/benchmark	6.099s
```

### Insights
- Preallocated slices are faster and more memory-efficient, as they avoid runtime resizing and memory allocation.