package compare

import (
	"github.com/stretchr/testify/assert"
)

type Integer interface {
	Signed | Unsigned
}
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func AreUnique[T comparable](args ...T) bool {
	for i, v := range args {
		for j := i + 1; j < len(args); j++ {
			if assert.ObjectsAreEqual(v, args[j]) {
				return false
			}
		}
	}
	return true
}

func AreEqual[T comparable](args ...T) bool {
	for i := 1; i < len(args); i++ {
		if args[0] != args[i] {
			return false
		}
	}
	return true
}

func AreUniqueValues(args ...any) bool {
	for i, v := range args {
		for j := i + 1; j < len(args); j++ {
			if assert.ObjectsAreEqualValues(v, args[j]) {
				return false
			}
		}
	}
	return true
}

func AreEqualValues(args ...any) bool {
	for i, v := range args {
		for j := i + 1; j < len(args); j++ {
			if !assert.ObjectsAreEqualValues(v, args[j]) {
				return false
			}
		}
	}
	return true
}
