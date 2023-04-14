package funcs

import "testing"

func BenchmarkSlowSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortSlice(false)
	}
}

func BenchmarkFastSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortSlice(true)
	}
}
