package sort

import "github.com/sayden/algorithms-in-go/common"

func InsertionUint8(u []uint8) []uint8 {

	newL := make([]uint8, len(u))
	copy(newL, u)

	isLess := func(a, b uint8) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func InsertionUint16(u []uint16) []uint16 {

	newL := make([]uint16, len(u))
	copy(newL, u)

	isLess := func(a, b uint16) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func InsertionUint32(u []uint32) []uint32 {

	newL := make([]uint32, len(u))
	copy(newL, u)

	isLess := func(a, b uint32) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func InsertionUint64(u []uint64) []uint64 {

	newL := make([]uint64, len(u))
	copy(newL, u)

	isLess := func(a, b uint64) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func InsertionInt8(u []int8) []int8 {

	newL := make([]int8, len(u))
	copy(newL, u)

	isLess := func(a, b int8) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func InsertionInt16(u []int16) []int16 {

	newL := make([]int16, len(u))
	copy(newL, u)

	isLess := func(a, b int16) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func InsertionInt32(u []int32) []int32 {

	newL := make([]int32, len(u))
	copy(newL, u)

	isLess := func(a, b int32) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func InsertionInt64(u []int64) []int64 {

	newL := make([]int64, len(u))
	copy(newL, u)

	isLess := func(a, b int64) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func InsertionFloat32(u []float32) []float32 {

	newL := make([]float32, len(u))
	copy(newL, u)

	isLess := func(a, b float32) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func InsertionFloat64(u []float64) []float64 {

	newL := make([]float64, len(u))
	copy(newL, u)

	isLess := func(a, b float64) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func InsertionByte(u []byte) []byte {

	newL := make([]byte, len(u))
	copy(newL, u)

	isLess := func(a, b byte) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func InsertionRune(u []rune) []rune {

	newL := make([]rune, len(u))
	copy(newL, u)

	isLess := func(a, b rune) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func InsertionUint(u []uint) []uint {

	newL := make([]uint, len(u))
	copy(newL, u)

	isLess := func(a, b uint) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func InsertionInt(u []int) []int {
	newL := make([]int, len(u))
	copy(newL, u)

	isLess := func(a, b int) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			common.SwapInt(newL, j, j-1)
		}
	}

	return newL
}

func InsertionUintptr(u []uintptr) []uintptr {
	newL := make([]uintptr, len(u))
	copy(newL, u)

	isLess := func(a, b uintptr) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

func hello(n int) int {
	return n ^ 7
}
