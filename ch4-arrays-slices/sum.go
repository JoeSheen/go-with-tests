package main

/*// part 2.
func Sum(numbers [5]int) int {
	return 0
}*/

/*// part 3.
func Sum(numbers [5]int) int {
	sum := 0

	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}

	return sum
}*/

/*// part 4.
func Sum(numbers [5] int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}*/
/* range lets you iterate over an array.
   On each iteration, range returns two values:
   	the index and the value.
   We are choosing to ignore the index value by using '_' known as the blank identifier.
*/

// part 6.
func Sum(numbers []int) int {
	sum := 0
	
	for _, number := range numbers {
		sum += number
	}
	
	return sum
}

/*// part 8.
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	
	// make allows you to create a slice with a starting capacity of the len of the numbersToSum
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}*/

// part 9.
func SumAll(numbersToSum ...[]int) []int {
	var sums [] int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
		// the append function which takes a slice and a new value, then returns a new slice with all the items in it
	}

	return sums
}

/*// part 11.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		// Slices can be sliced! The syntax is slice[low:high]. 
		// If you omit the value on one of the sides of the ':' it captures everything to that side of it
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}*/

// part 13.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	
	return sums
}
