package concurrently

import (
	"time"
)

func makeSequence(length int, multiple int) []int {
	numbers := make([]int, length)
	for i := range numbers {
		numbers[i] = i * multiple
	}
	return numbers
}

func doubleNumber(number int) int {
	return number * 2
}

func slowDoubleNumber(number int) int {
	time.Sleep(1 * time.Millisecond)
	return doubleNumber(number)
}

func slicesIntEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
