package concurrently

import (
	"testing"
	"time"
)

func TestEach(t *testing.T) {
	numbers := makeSequence(10, 1)

	Each(numbers, doubleNumber, 4)
}

func TestEachSerial(t *testing.T) {
	numbers := makeSequence(10, 1)

	EachSerial(numbers, doubleNumber)
}

func benchmarkEach(count int, b *testing.B) {
	numbers := makeSequence(100, 1)

	for n := 0; n < b.N; n++ {
		Each(numbers, sleep100, count)
	}
}

func benchmarkEachSerial(b *testing.B) {
	numbers := makeSequence(100, 1)

	for n := 0; n < b.N; n++ {
		EachSerial(numbers, sleep100)
	}
}

func benchmarkEachOriginal(b *testing.B) {
	numbers := makeSequence(100, 1)

	for n := 0; n < b.N; n++ {
		for _, number := range numbers {
			sleep100(number)
		}
	}
}

func sleep100(number int) {
	time.Sleep(1 * time.Millisecond)
}

func BenchmarkEach001(b *testing.B)      { benchmarkEach(1, b) }
func BenchmarkEach010(b *testing.B)      { benchmarkEach(10, b) }
func BenchmarkEach100(b *testing.B)      { benchmarkEach(100, b) }
func BenchmarkEachSerial(b *testing.B)   { benchmarkEachSerial(b) }
func BenchmarkEachOriginal(b *testing.B) { benchmarkEachOriginal(b) }
