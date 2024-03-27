package compare

import (
	"cmp"

	"github.com/stretchr/testify/assert"
)

type Logger interface {
	Error(args ...any)
	Errorf(format string, args ...any)
}

func Unique[T comparable](w Logger, args ...T) bool {
	for i, v := range args {
		for j := i + 1; j < len(args); j++ {
			if v == args[j] {
				w.Errorf("Should not be: %#v  @%d %d", v, i, j)
				return false
			}
		}
	}
	return true
}

func Equal[T comparable](w Logger, args ...T) bool {
	for i := 1; i < len(args); i++ {
		if args[0] != args[i] {
			w.Errorf("Not equal: \n"+
				"expected: %s\n"+
				"actual  : %s @%d", args[0], args[i], i)
			return false
		}
	}
	return true
}

func UniqueValues(w Logger, args ...any) bool {
	for i, v := range args {
		for j := i + 1; j < len(args); j++ {
			if assert.ObjectsAreEqualValues(v, args[j]) {
				w.Errorf("Not unique: \n"+
					"%d: %s @\n"+
					"%d: %s ", i, v, j, args[j])
				return false
			}
		}
	}
	return true
}

func EqualValues(w Logger, args ...any) bool {
	for i, v := range args {
		for j := i + 1; j < len(args); j++ {
			if !assert.ObjectsAreEqualValues(v, args[j]) {
				w.Errorf("Not equal: \n"+
					"expected: %s @%d\n"+
					"actual  : %s @%d", v, i, args[j], j)
				return false
			}
		}
	}
	return true
}

// Greater returns a function that returns true when a value is greater
// than a threshold.
func Greater[T cmp.Ordered](t T) func(T) bool {
	return func(s T) bool {
		return t > s
	}
}

// GreaterOrEqual returns a function that returns true when a value is
// greater than or equal to a threshold.
func GreaterOrEqual[T cmp.Ordered](t T) func(T) bool {
	return func(s T) bool {
		return t >= s
	}
}
