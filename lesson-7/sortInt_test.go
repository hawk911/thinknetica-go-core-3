package lesson_7

import (
	"reflect"
	"testing"
)

func TestSortInt(t *testing.T) {
	i := []int{1, 5, 2, 0}
	SortInt(i)
	got := i
	want := []int{0, 1, 2, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}
