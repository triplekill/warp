package utils

func NumberRange(min, max int) []int {

	count := max - min
	numbers := make([]int, count)
	n := min
	for i := 0; i < count; i++ {
		numbers[i] = n
		n++
	}

	return numbers
}
