package arrays_and_slices

func Filter(s []int, fn func(int) bool) []int {
	var p []int

	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func Map(s []int, fn func(int) int) []int {
	var p []int

	for _, v := range s {
		p = append(p, fn(v))
	}

	return p
}
