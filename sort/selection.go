package sort

func SelectionUint8(u []uint8) []uint8 {
	newList := make([]uint8, len(u))
	tempList := make([]uint8, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}

func SelectionUint16(u []uint16) []uint16 {
	newList := make([]uint16, len(u))
	tempList := make([]uint16, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}

func SelectionUint32(u []uint32) []uint32 {
	newList := make([]uint32, len(u))
	tempList := make([]uint32, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}

func SelectionUint64(u []uint64) []uint64 {
	newList := make([]uint64, len(u))
	tempList := make([]uint64, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}

func SelectionInt8(u []int8) []int8 {
	newList := make([]int8, len(u))
	tempList := make([]int8, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}

func SelectionInt16(u []int16) []int16 {
	newList := make([]int16, len(u))
	tempList := make([]int16, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}

func SelectionInt32(u []int32) []int32 {
	newList := make([]int32, len(u))
	tempList := make([]int32, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}

func SelectionInt64(u []int64) []int64 {
	newList := make([]int64, len(u))
	tempList := make([]int64, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}

func SelectionFloat32(u []float32) []float32 {
	newList := make([]float32, len(u))
	tempList := make([]float32, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}

func SelectionFloat64(u []float64) []float64 {
	newList := make([]float64, len(u))
	tempList := make([]float64, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}

func SelectionByte(u []byte) []byte {
	newList := make([]byte, len(u))
	tempList := make([]byte, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}

func SelectionRune(u []rune) []rune {
	newList := make([]rune, len(u))
	tempList := make([]rune, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}
