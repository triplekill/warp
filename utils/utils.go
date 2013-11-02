package utils

import (

	"fmt"
)

func Panicf(format string, args ...interface{}) {

	msg := fmt.Sprintf(format, args...)
	panic(msg)
}

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
