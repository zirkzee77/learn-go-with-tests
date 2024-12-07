package arrayandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("Sum of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
  t.Run("One input", func(t *testing.T) {
    got := SumAll([]int{1, 2}, []int{3, 4, 5})
    want := []int{3, 12}
    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })
}
