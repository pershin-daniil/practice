package slices

import "testing"

var input = []string{"1", "2", "3", "4", "5"}

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copyListFixed(input)
	}
}

//func BenchmarkCopyFixed(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		copyListFixed(input)
//	}
//}
