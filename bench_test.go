package compare

import (
	"testing"
)

func BenchmarkAreEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equal(b, i, i, i)
	}
}

func BenchmarkAreUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unique(b, i, i+1, i+2)
	}
}

func BenchmarkGreater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greater(i)(i+1)
	}
}

func BenchmarkGreaterLegacy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i+1>i{}
	}
}

func BenchmarkGreaterOrEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreaterOrEqual(i)(i)
	}
}

func BenchmarkGreaterThan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := float32(i)
		Greater(float32(-1.0))(j)
	}
}

func BenchmarkGreaterThanEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := float64(i)
		GreaterOrEqual(i)(j)
	}
}
