package compare

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreEqual(t *testing.T) {
	type testCase struct {
		name string
		args []float64
		want bool
	}
	tests := []testCase{
		{
			name: "1",
			args: []float64{math.NaN(), math.NaN()},
			want: false,
		},
		{
			name: "1",
			args: []float64{6, 6},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, AreEqual(tt.args...), "AreEqual(%v)", tt.args)
		})
	}
}

func TestAreEqualValue(t *testing.T) {
	type testCase struct {
		name string
		args []any
		want bool
	}
	tests := []testCase{
		{
			name: "1",
			args: []any{math.NaN(), math.NaN()},
			want: false,
		},
		{
			name: "1",
			args: []any{6, 6.0, float32(6), int8(6), uint(6)},
			want: true,
		},
		{
			name: "1",
			args: []any{6, 6.0, float32(6), int8(6), uint(6), "6"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, AreEqualValues(tt.args...), "AreEqual(%v)", tt.args)
		})
	}
}

func FuzzEqualValues(f *testing.F) {
	var (
		a int
		b uint
		c int8
		d int16
		e int32
		g uint64
		j float32 = float32(math.NaN())
		k float64 = math.NaN()
	)
	f.Add(a, b, c, d, e, g, j, k)
	f.Fuzz(func(t *testing.T, a int, b uint, c int8, d int16, e int32, g uint64, j float32, k float64) {
		arr := []any{a, b, c, d, e, g, j, k}
		if AreEqual(arr...) {
			assert.True(t, AreEqualValues(arr...))
		}
		if AreUniqueValues(arr...) {
			assert.True(t, AreUnique(arr...))
		}
	})
}
func FuzzInt(f *testing.F)     { fuzzEqual[int](f) }
func FuzzUnt(f *testing.F)     { fuzzEqual[uint](f) }
func FuzzInt8(f *testing.F)    { fuzzEqual[int8](f) }
func FuzzInt16(f *testing.F)   { fuzzEqual[int16](f) }
func FuzzInt32(f *testing.F)   { fuzzEqual[int32](f) }
func FuzzUint64(f *testing.F)  { fuzzEqual[uint64](f) }
func FuzzFloat32(f *testing.F) { fuzzEqual[float32](f) }
func FuzzFloat64(f *testing.F) { fuzzEqual[float64](f) }
func FuzzString(f *testing.F)  { fuzzEqual[string](f) }

func fuzzEqual[T comparable](f *testing.F) {
	var (
		a T
		b T
	)
	f.Add(a, b)
	f.Fuzz(func(t *testing.T, a, b T) {
		if AreEqual(a, b) {
			assert.True(t, AreEqualValues(a, b))
		}
		if AreUniqueValues(a, b) {
			assert.True(t, AreUnique(a, b))
		}
	})
}
