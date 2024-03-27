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
		Unique(b, i, i, i)
	}
}

func BenchmarkGreater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greater(-1)(i)
	}
}

func BenchmarkGreaterOrEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreaterOrEqual(0)(i)
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
		GreaterOrEqual(-0.0)(j)
	}
}
