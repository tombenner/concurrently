package concurrently

import (
	"math"
	"testing"
	"time"
)

func TestFilter(t *testing.T) {
	numbers := makeSequence(10, 1)
	evenNumbers := makeSequence(5, 2)

	result := Filter(numbers, isEvenNumber, 4).([]int)

	if !slicesIntEqual(result, evenNumbers) {
		t.Errorf("Receved: %q\nExpected: %q", result, evenNumbers)
	}
}

func TestFilterSerial(t *testing.T) {
	numbers := makeSequence(10, 1)
	evenNumbers := makeSequence(5, 2)

	result := FilterSerial(numbers, isEvenNumber).([]int)

	if !slicesIntEqual(result, evenNumbers) {
		t.Errorf("Receved: %q\nExpected: %q", result, evenNumbers)
	}
}

func benchmarkFilter(count int, b *testing.B) {
	numbers := makeSequence(100, 1)

	for n := 0; n < b.N; n++ {
		Filter(numbers, slowIsEvenNumber, count)
	}
}

func benchmarkFilterSerial(b *testing.B) {
	numbers := makeSequence(100, 1)

	for n := 0; n < b.N; n++ {
		FilterSerial(numbers, slowIsEvenNumber)
	}
}

func benchmarkFilterOriginal(b *testing.B) {
	numbers := makeSequence(100, 1)

	for n := 0; n < b.N; n++ {
		for _, number := range numbers {
			slowIsEvenNumber(number)
		}
	}
}

func isEvenNumber(number int) bool {
	return int(math.Mod(float64(number), float64(2))) == 0
}

func slowIsEvenNumber(number int) bool {
	time.Sleep(1 * time.Millisecond)
	return isEvenNumber(number)
}

func BenchmarkFilter001(b *testing.B)      { benchmarkFilter(1, b) }
func BenchmarkFilter010(b *testing.B)      { benchmarkFilter(10, b) }
func BenchmarkFilter100(b *testing.B)      { benchmarkFilter(100, b) }
func BenchmarkFilterSerial(b *testing.B)   { benchmarkFilterSerial(b) }
func BenchmarkFilterOriginal(b *testing.B) { benchmarkFilterOriginal(b) }
