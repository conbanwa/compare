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
			name: "2",
			args: []float64{6, 6},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Equal(t, tt.args...), "Equal(t, %v)", tt.args)
			assert.NotEqual(t, Equal(t, tt.args...), Unique(t, tt.args...))
			for _, v := range tt.args {
				assert.False(t, Greater(v)(6))
			}
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
			name: "nan",
			args: []any{math.NaN(), math.NaN()},
			want: false,
		},
		{
			name: "6",
			args: []any{6, 6.0, float32(6), int8(6), uint(6)},
			want: true,
		},
		{
			name: "six",
			args: []any{6, 6.0, float32(6), int8(6), uint(6), "6"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, EqualValues(t, tt.args...), "Equal(t, %v)", tt.args)
		})
	}
}

func FuzzEqualValuesTypes(t *testing.F) {
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
	t.Add(a, b, c, d, e, g, h, j, k, l)
	t.Fuzz(func(t *testing.T, a int, b uint, c int8, d int16, e int32, g uint64, h byte, j float32, k float64, l bool) {
		arr := []any{a, b, c, d, e, g, j, k}
		if UniqueValues(t, arr...) {
			assert.True(t, Unique(t, arr...))
			assert.True(t, !Equal(t, arr...))
			assert.True(t, !EqualValues(t, arr...))
		}
		if Unique(t, arr...) {
			assert.True(t, !Equal(t, arr...))
			assert.True(t, !EqualValues(t, arr...))
		}
		if Equal(t, arr...) {
			assert.True(t, EqualValues(t, arr...))
			assert.True(t, !Unique(t, arr...))
		}
	})
}
func FuzzEqualValues(t *testing.F) {
	var (
		a, b, c, d, e, g, h, j, k, l int
	)
	t.Add(a, b, c, d, e, g, h, j, k, l)
	t.Fuzz(func(t *testing.T, a, b, c, d, e, g, h, j, k, l int) {
		arr := []any{a, b, c, d, e, g, j, k}
		//if assert.True(t, EqualValues(t,arr...)){
		//	return
		//}
		if UniqueValues(t, arr...) {
			assert.True(t, Unique(t, arr...))
			assert.True(t, !Equal(t, arr...))
			assert.True(t, !EqualValues(t, arr...))
		}
		if Unique(t, arr...) {
			assert.True(t, UniqueValues(t, arr...))
			assert.True(t, !Equal(t, arr...))
			assert.True(t, !EqualValues(t, arr...))
		}
		if EqualValues(t, arr...) {
			assert.True(t, Equal(t, arr...))
			assert.True(t, !Unique(t, arr...))
			assert.True(t, !UniqueValues(t, arr...))
		}
	})
}

func FuzzEqualInt(t *testing.F)     { fuzzEqual[int](t) }
func FuzzEqualUnt(t *testing.F)     { fuzzEqual[uint](t) }
func FuzzEqualInt8(t *testing.F)    { fuzzEqual[int8](t) }
func FuzzEqualInt16(t *testing.F)   { fuzzEqual[int16](t) }
func FuzzEqualInt32(t *testing.F)   { fuzzEqual[int32](t) }
func FuzzEqualUint64(t *testing.F)  { fuzzEqual[uint64](t) }
func FuzzEqualFloat32(t *testing.F) { fuzzEqual[float32](t) }
func FuzzEqualFloat64(t *testing.F) { fuzzEqual[float64](t) }
func FuzzEqualString(t *testing.F)  { fuzzEqual[string](t) }
func FuzzEqualByte(t *testing.F)    { fuzzEqual[byte](t) }

func fuzzEqual[T comparable](t *testing.F) {
	var (
		a T
		b T
		c T
		d T
	)
	t.Add(a, b, c, d)
	t.Fuzz(func(t *testing.T, a, b, c, d T) {
		assert.Equal(t, Equal(t, a, b, c, d), EqualValues(t, a, b, c, d))
		assert.Equal(t, Unique(t, a, b, c, d), UniqueValues(t, a, b, c, d))
	})
}
