package main

/*The code for this chapter is a continuation from arrays and slices (Chapter 4).*/

func Sum(numbers []int) int {
	add := func(acc, x int) int {
		return acc + x
	}
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	addTails := func(acc, x []int) [] int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbersToSum, addTails, []int{})
}

func Reduce[T, X any](collection []T, accumulator func(X, T) X, initialValue X) X {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

func Find[T comparable](collection []T, comparator func(T) bool) (value T, found bool) {
	for _, x := range collection {
		if comparator(x) {
			return x, true
		}
	}
	return
}
