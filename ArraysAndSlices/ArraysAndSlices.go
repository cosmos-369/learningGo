package main

func ArraySum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, i := range numbersToSum {
		sums = append(sums, ArraySum(i))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, i := range numbersToSum {
		if len(i) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, ArraySum(i[1:]))
		}
	}

	return sums
}
