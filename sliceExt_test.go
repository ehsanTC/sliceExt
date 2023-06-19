package sliceExt

import "testing"

func TestContains(t *testing.T) {
	slice := []int{1}

	// Check 1 exists
	want := true
	got := Contains[int](slice, 1)
	if want != got {
		t.Error("The contains() did not detect the element")
	}

	want = false
	got = Contains[int](slice, 100)
	if want != got {
		t.Error("The contains() detected the wrong element")
	}
}

func TestAddShouldAppendItemToEndOfSlice(t *testing.T) {
	slice := make([]int, 10)

	got := Add(&slice, 100)
	want := 10

	if got != want {
		t.Errorf("The item should be in the position %q, but got %q", want, got)
	}
}

func TestAddShouldIncreaseSliceSize(t *testing.T) {
	size := 10
	slice := make([]int, size)

	Add(&slice, 100)
	got := len(slice)
	want := size + 1

	if got != want {
		t.Errorf("The slice size should be %q, but got %q", want, got)
	}
}

func TestIndexOfShouldReturnIndexOfExistingItem(t *testing.T) {
	names := []string{"Ehsan", "Ali"}

	for index, name := range names {
		want := index
		got := IndexOf[string](names, name)

		if want != got {
			t.Errorf("The index of %q is %q, but got %q", name, want, got)
		}
	}
}

func TestIndexOfShouldReturnMinusOfNonExistingItem(t *testing.T) {
	names := []string{"Ehsan", "Ali"}

	want := -1
	got := IndexOf[string](names, "Mohammad")

	if want != got {
		t.Errorf("The index of %q is %q, but got %q", "Mohammad", want, got)
	}
}

func TestRemoveShouldRemoveItem(t *testing.T) {
	names := []string{"Ehsan", "Ali"}
	removingName := "Ehsan"

	Remove[string](&names, removingName)

	for _, name := range names {
		if name == removingName {
			t.Errorf("The removed item is found in the slice!")
		}
	}
}

func TestRemoveAt(t *testing.T) {
	names := []string{"Ehsan", "Ali"}
	removingIndex := 1
	removingName := names[removingIndex]
	primaryLen := len(names)

	RemoveAt[string](&names, removingIndex)

	if primaryLen != len(names)+1 {
		t.Errorf("The item is not removed from the slice.")
	}

	for _, name := range names {
		if name == removingName {
			t.Errorf("The removed item is found in the slice!")
		}
	}
}

func TestRemoveRange(t *testing.T) {
	slice := make([]int, 10)
	for i := 0; i < 10; i++ {
		slice[i] = i
	}

	primaryLen := len(slice)

	removingSlice := []int{1, 2, 3, 4}
	removingSliceLen := len(removingSlice)

	RemoveRange[int](&slice,
		IndexOf[int](slice, removingSlice[0]),
		removingSliceLen)

	if len(slice) != primaryLen-removingSliceLen {
		t.Errorf("The items are not removed from the slice completely!")
	}

	for _, item := range removingSlice {
		if Contains[int](slice, item) {
			t.Errorf("The removed item is found in the slice!")
		}
	}
}

func TestInsert(t *testing.T) {
	numbers := new([3]int)[0:2]
	Insert(&numbers, 1, 100)

	got := IndexOf(numbers, 100)
	want := 1
	if got != want {
		t.Errorf("The inserted item is not in the correct place")
	}

}

func TestInsertRange(t *testing.T) {
	chars := []string{"x", "y"}
	newChars := []string{"a", "b", "c"}

	charsLen := len(chars)

	insertIndex := 0
	InsertRange(&chars, insertIndex, newChars)

	if len(chars) != len(newChars)+charsLen {
		t.Error("The length of final slice is wrong!")
	}

	for i := 0; i < len(newChars); i++ {
		if chars[insertIndex+i] != newChars[i] {
			t.Errorf("The inserted item (%q) not found in the result.", newChars[i])
		}
	}
}

func TestInsertRangePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not get the panic")
		}
	}()

	numbers := new([3]int)[0:2]
	Insert(&numbers, 3, 100)
}
