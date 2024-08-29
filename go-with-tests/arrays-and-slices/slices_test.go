package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("should filter even numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		isEven := func(n int) bool {
			return n%2 == 0
		}

		got := Filter(numbers, isEven)
		want := []int{2, 4, 6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("should filter even numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		double := func(n int) int {
			return n * 2
		}

		got := Map(numbers, double)
		want := []int{2, 4, 6, 8, 10, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
