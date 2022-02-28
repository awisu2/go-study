package main

// IndexOf search value and return first match index
//
// return -1 if not match
// slice must be slice type, slice and search must be of the same type.
func IndexOf(slice interface{}, search interface{}, fromIndex int) int {
	switch search.(type) {
	case int:
		return IndexOfInt(slice.([]int), search.(int), fromIndex)
	case int64:
		return IndexOfInt64(slice.([]int64), search.(int64), fromIndex)
	case int16:
		return IndexOfInt(slice.([]int), search.(int), fromIndex)
	case int8:
		return IndexOfInt8(slice.([]int8), search.(int8), fromIndex)
	case string:
		return IndexOfStrinig(slice.([]string), search.(string), fromIndex)
	case float32:
		return IndexOfFloat32(slice.([]float32), search.(float32), fromIndex)
	case float64:
		return IndexOfFloat64(slice.([]float64), search.(float64), fromIndex)
	case byte:
		return IndexOfFloat64(slice.([]float64), search.(float64), fromIndex)
	}
	return -1
}

func Contain(slice interface{}, search interface{}, fromIndex int) bool {
	return IndexOf(slice, search, fromIndex) != -1
}

func IndexOfInt(slice []int, search int, fromIndex int) int {
	for i, item := range slice[fromIndex:] {
		if item == search {
			return i
		}
	}
	return -1
}

func IndexOfInt64(slice []int64, search int64, fromIndex int) int {
	for i, item := range slice[fromIndex:] {
		if item == search {
			return i
		}
	}
	return -1
}

func IndexOfInt16(slice []int16, search int16, fromIndex int) int {
	for i, item := range slice[fromIndex:] {
		if item == search {
			return i
		}
	}
	return -1
}

func IndexOfInt8(slice []int8, search int8, fromIndex int) int {
	for i, item := range slice[fromIndex:] {
		if item == search {
			return i
		}
	}
	return -1
}

func IndexOfStrinig(slice []string, search string, fromIndex int) int {
	for i, item := range slice[fromIndex:] {
		if item == search {
			return i
		}
	}
	return -1
}

func IndexOfFloat32(slice []float32, search float32, fromIndex int) int {
	for i, item := range slice[fromIndex:] {
		if item == search {
			return i
		}
	}
	return -1
}

func IndexOfFloat64(slice []float64, search float64, fromIndex int) int {
	for i, item := range slice[fromIndex:] {
		if item == search {
			return i
		}
	}
	return -1
}

func IndexOfByte(slice []byte, search byte, fromIndex int) int {
	for i, item := range slice[fromIndex:] {
		if item == search {
			return i
		}
	}
	return -1
}
