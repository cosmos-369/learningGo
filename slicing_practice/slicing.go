package slicePractice

func SliceSplitHalf(slice []int) [][]int {
	mid := len(slice) / 2
	firstHalf := slice[:mid]
	secondHalf := slice[mid:]

	return [][]int{firstHalf, secondHalf}

}

func SliceSplitN(slice []int, n int) [][]int {

	slices := [][]int{}

	sliceLen := len(slice) / n
	for i := 0; i < n; i++ {
		newSlice := make([]int, sliceLen)
		if (n+sliceLen*i)+1 > len(slice) {
			copy(newSlice, slice[sliceLen*i:])
		} else {
			copy(newSlice, slice[sliceLen*i:(n+sliceLen*i)+1])
		}

		slices = append(slices, newSlice) // Removed the redeclaration
	}
	return slices
}
