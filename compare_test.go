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
			assert.NotEqual(t, AreEqual(tt.args...), AreUnique(tt.args...))
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

func FuzzEqualValuesTypes(f *testing.F) {
	var (
		a int
		b uint
		c int8
		d int16
		e int32
		g uint64
		h byte
		j float32 = float32(math.NaN())
		k float64 = math.NaN()
		l bool
	)
	f.Add(a, b, c, d, e, g, h, j, k, l)
	f.Fuzz(func(t *testing.T, a int, b uint, c int8, d int16, e int32, g uint64, h byte, j float32, k float64, l bool) {
		arr := []any{a, b, c, d, e, g, j, k}
		if AreUniqueValues(arr...) {
			assert.True(t, AreUnique(arr...))
			assert.True(t, !AreEqual(arr...))
			assert.True(t, !AreEqualValues(arr...))
		}
		if AreUnique(arr...) {
			assert.True(t, !AreEqual(arr...))
			assert.True(t, !AreEqualValues(arr...))
		}
		if AreEqual(arr...) {
			assert.True(t, AreEqualValues(arr...))
			assert.True(t, !AreUnique(arr...))
		}
	})
}
func FuzzEqualValues(f *testing.F) {
	var (
		a, b, c, d, e, g, h, j, k, l int
	)
	f.Add(a, b, c, d, e, g, h, j, k, l)
	f.Fuzz(func(t *testing.T, a, b, c, d, e, g, h, j, k, l int) {
		arr := []any{a, b, c, d, e, g, j, k}
		//if assert.True(t, AreEqualValues(arr...)){
		//	return
		//}
		if AreUniqueValues(arr...) {
			assert.True(t, AreUnique(arr...))
			assert.True(t, !AreEqual(arr...))
			assert.True(t, !AreEqualValues(arr...))
		}
		if AreUnique(arr...) {
			assert.True(t, AreUniqueValues(arr...))
			assert.True(t, !AreEqual(arr...))
			assert.True(t, !AreEqualValues(arr...))
		}
		if AreEqualValues(arr...) {
			assert.True(t, AreEqual(arr...))
			assert.True(t, !AreUnique(arr...))
			assert.True(t, !AreUniqueValues(arr...))
		}
	})
}

func FuzzEqualInt(f *testing.F)     { fuzzEqual[int](f) }
func FuzzEqualUnt(f *testing.F)     { fuzzEqual[uint](f) }
func FuzzEqualInt8(f *testing.F)    { fuzzEqual[int8](f) }
func FuzzEqualInt16(f *testing.F)   { fuzzEqual[int16](f) }
func FuzzEqualInt32(f *testing.F)   { fuzzEqual[int32](f) }
func FuzzEqualUint64(f *testing.F)  { fuzzEqual[uint64](f) }
func FuzzEqualFloat32(f *testing.F) { fuzzEqual[float32](f) }
func FuzzEqualFloat64(f *testing.F) { fuzzEqual[float64](f) }
func FuzzEqualString(f *testing.F)  { fuzzEqual[string](f) }
func FuzzEqualByte(f *testing.F)    { fuzzEqual[byte](f) }

func fuzzEqual[T comparable](f *testing.F) {
	var (
		a T
		b T
		c T
		d T
	)
	f.Add(a, b, c, d)
	f.Fuzz(func(t *testing.T, a, b, c, d T) {
		assert.Equal(t, AreEqual(a, b, c, d), AreEqualValues(a, b, c, d))
		assert.Equal(t, AreUnique(a, b, c, d), AreUniqueValues(a, b, c, d))
	})
}
