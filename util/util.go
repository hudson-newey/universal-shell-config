package util

// TODO: remove this package (I don't like util files)

func ValueInArray[T comparable](value T, array []T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
