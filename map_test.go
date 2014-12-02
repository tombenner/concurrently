package concurrently

import (
	"testing"
)

func TestMap(t *testing.T) {
	numbers := makeSequence(10, 1)
	doubledNumbers := makeSequence(10, 2)

	result := Map(numbers, doubleNumber, 4).([]int)

	if !slicesIntEqual(result, doubledNumbers) {
		t.Errorf("Receved: %q\nExpected: %q", result, doubledNumbers)
	}
}

func TestMapSerial(t *testing.T) {
	numbers := makeSequence(10, 1)
	doubledNumbers := makeSequence(10, 2)

	result := MapSerial(numbers, doubleNumber).([]int)

	if !slicesIntEqual(result, doubledNumbers) {
		t.Errorf("Receved: %q\nExpected: %q", result, doubledNumbers)
	}
}

func benchmarkMap(count int, b *testing.B) {
	numbers := makeSequence(100, 1)

	for n := 0; n < b.N; n++ {
		Map(numbers, slowDoubleNumber, count)
	}
}

func benchmarkMapSerial(b *testing.B) {
	numbers := makeSequence(100, 1)

	for n := 0; n < b.N; n++ {
		MapSerial(numbers, slowDoubleNumber)
	}
}

func benchmarkMapOriginal(b *testing.B) {
	numbers := makeSequence(100, 1)

	for n := 0; n < b.N; n++ {
		for _, number := range numbers {
			slowDoubleNumber(number)
		}
	}
}

func BenchmarkMap001(b *testing.B)      { benchmarkMap(1, b) }
func BenchmarkMap010(b *testing.B)      { benchmarkMap(10, b) }
func BenchmarkMap100(b *testing.B)      { benchmarkMap(100, b) }
func BenchmarkMapSerial(b *testing.B)   { benchmarkMapSerial(b) }
func BenchmarkMapOriginal(b *testing.B) { benchmarkMapOriginal(b) }
