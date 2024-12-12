package benchmark_test

import (
	"testing"
)

// BenchmarkArrayAppend tests the performance of appending to a slice
func BenchmarkArrayAppend(b *testing.B) {
	// Test appending to an uninitialized slice
	b.Run("AppendToDynamicSlice", func(b *testing.B) {
		arr := []int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			arr = append(arr, i)
		}
	})

	// Test assigning values to a preallocated slice
	b.Run("AssignToPreallocatedSlice", func(b *testing.B) {
		arr := make([]int, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			arr[i] = i
		}
	})
}
