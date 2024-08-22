package arrays_and_slices

/*
Arrays have a fixed capacity which you define when you declare the variable.
We can initialize an array in two ways:

[N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}

[...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5}


You may be thinking it's quite cumbersome that arrays have a fixed length, and most of the time you probably won't be using them!

Go has slices which do not encode the size of the collection and instead can have any size.


*/

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}

func SumAll(lists ...[]int) []int {
	length := len(lists)
	sums := make([]int, length)

	for i, numbers := range lists {
		sums[i] = Sum(numbers)
	}

	return sums
}
