package sliceExt

var ERR_OutOfIndex string = "The inserting index is out of range!"

func Contains[T comparable](slice []T, item T) bool {
	for i := 0; i < len(slice); i++ {
		if item == slice[i] {
			return true
		}
	}
	return false
}

func Add[T any](slice *[]T, item T) (index int) {
	index = len(*slice)
	*slice = append(*slice, item)

	return
}

func IndexOf[T comparable](slice []T, item T) int {
	for index, element := range slice {
		if element == item {
			return index
		}
	}
	return -1
}

func Remove[T comparable](slice *[]T, item T) {
	if removingIndex := IndexOf[T](*slice, item); removingIndex >= 0 {
		RemoveAt[T](slice, removingIndex)
	}
}

func RemoveAt[T any](slice *[]T, index int) {
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}

func RemoveRange[T any](slice *[]T, index int, count int) {
	if index < 0 || index >= len(*slice) {
		panic(ERR_OutOfIndex)
	}

	if index+count > len(*slice) {
		panic(ERR_OutOfIndex)
	}

	*slice = append((*slice)[:index], (*slice)[index+count:]...)
}

func Insert[T any](slice *[]T, index int, item T) {
	InsertRange(slice, index, []T{item})
}

func InsertRange[T any](slice *[]T, index int, collection []T) {
	if index >= len(*slice) {
		panic(ERR_OutOfIndex)
	}

	if freeSpace := cap(*slice) - len(*slice); len(collection) > freeSpace {
		// Increase slice's length
		*slice = append(*slice, make([]T, len(collection)-freeSpace)...)
	}

	copy((*slice)[index+len(collection):], (*slice)[index:])
	copy((*slice)[index:], collection)
}
