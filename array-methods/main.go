package main

import "fmt"

func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T

	for _, v := range slice {
		response := fn(v)

		if response {
			result = append(result, v)
		}
	}

	return result
}

func Map[T, R any](slice []T, fn func(T) R) []R {
	var result []R

	for _, v := range slice {
		result = append(result, fn(v))
	}

	return result
}

func ForEach[T any](slice []T, fn func(T)) {
	for _, v := range slice {
		fn(v)
	}
}

func Some[T any](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}

	return false
}

func Every[T any](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}

	return true
}

func Contains[T comparable](slice []T, e T) bool {
	for _, v := range slice {
		if v == e {
			return true
		}
	}

	return false
}

// TODO: REDUCE

func main() {
	s := []int{1, 2, 3, 4, 6}
	fmt.Printf("\nINITIAL %v", s)

	// Filter
	fn := func(n int) bool {
		return n%2 == 0
	}

	filtered := Filter(s, fn)

	fmt.Printf("\nFILTER %v", filtered)

	// Map

	f := func(n int) int {
		return n * 2
	}

	updated := Map(s, f)

	fmt.Printf("\nMAP %v", updated)
}
