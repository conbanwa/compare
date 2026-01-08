package compare

import (
	"cmp"
	"fmt"
	"strings"

	"github.com/stretchr/testify/assert"
)

type Logger interface {
	Error(args ...any)
	Errorf(format string, args ...any)
}

// helper to format the full argument list nicely
func formatArgs(args ...any) string {
	var sb strings.Builder
	for i, a := range args {
		if i > 0 {
			sb.WriteString(", ")
		}
		fmt.Fprintf(&sb, "%d: %#v", i, a)
	}
	return sb.String()
}

// helper to convert []T to []any
func anySlice[T any](s []T) []any {
	res := make([]any, len(s))
	for i, v := range s {
		res[i] = v
	}
	return res
}

// Unique checks if all values are strictly unique using ==
func Unique[T comparable](w Logger, args ...T) bool {
	for i, v := range args {
		for j := i + 1; j < len(args); j++ {
			if v == args[j] {
				w.Errorf("Unique[T] FAILED — duplicate found:\n"+
					"  Value: %#v\n"+
					"  Positions: %d and %d\n"+
					"  Full args: [%s]",
					v, i, j, formatArgs(anySlice(args)...))
				return false
			}
		}
	}
	return true
}

// Equal checks if all values are strictly equal using ==
func Equal[T comparable](w Logger, args ...T) bool {
	if len(args) == 0 {
		return true
	}
	first := args[0]
	for i := 1; i < len(args); i++ {
		if first != args[i] {
			w.Errorf("Equal[T] FAILED — not all equal:\n"+
				"  Expected all to be: %#v\n"+
				"  Found different: %#v at position %d\n"+
				"  Full args: [%s]",
				first, args[i], i, formatArgs(anySlice(args)...))
			return false
		}
	}
	return true
}

// UniqueValues checks if all values are unique using assert.ObjectsAreEqualValues
func UniqueValues(w Logger, args ...any) bool {
	for i, v := range args {
		for j := i + 1; j < len(args); j++ {
			if assert.ObjectsAreEqualValues(v, args[j]) {
				w.Errorf("UniqueValues FAILED — values considered equal:\n"+
					"  Position %d: %#v\n"+
					"  Position %d: %#v\n"+
					"  Full args: [%s]",
					i, v, j, args[j], formatArgs(args...))
				return false
			}
		}
	}
	return true
}

// EqualValues checks if all values are equal using assert.ObjectsAreEqualValues
func EqualValues(w Logger, args ...any) bool {
	if len(args) <= 1 {
		return true
	}
	for i := 1; i < len(args); i++ {
		if !assert.ObjectsAreEqualValues(args[0], args[i]) {
			w.Errorf("EqualValues FAILED — not all equal:\n"+
				"  Expected all like position 0: %#v\n"+
				"  Found different at position %d: %#v\n"+
				"  Full args: [%s]",
				args[0], i, args[i], formatArgs(args...))
			return false
		}
	}
	return true
}

// Greater returns a function that returns true when a value is greater than a threshold.
func Greater[T cmp.Ordered](t T) func(T) bool {
	return func(s T) bool {
		return t > s
	}
}

// GreaterOrEqual returns a function that returns true when a value is greater than or equal to a threshold.
func GreaterOrEqual[T cmp.Ordered](t T) func(T) bool {
	return func(s T) bool {
		return t >= s
	}
}
