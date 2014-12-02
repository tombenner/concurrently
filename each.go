package concurrently

import (
	"reflect"
	"sync"
)

// Each executes `fn` for each element of `slice`.
// `workersCount` is the number of goroutines that the work will be distributed across.
func Each(slice interface{}, fn interface{}, workersCount int) {
	sliceValue := reflect.ValueOf(slice)
	fnValue := reflect.ValueOf(fn)

	validateSliceAndFnValues(sliceValue, fnValue)

	maps := makeWorkerMaps(sliceValue, workersCount)

	var wg sync.WaitGroup
	wg.Add(workersCount)

	for i := 0; i < workersCount; i++ {
		go goEach(maps[i], fnValue, &wg)
	}

	wg.Wait()
}

// EachSerial is nearly identical to Each, but it performs the work serially rather than concurrently.
// (It doesn't use goroutines.)
func EachSerial(slice interface{}, fn interface{}) {
	sliceValue := reflect.ValueOf(slice)
	fnValue := reflect.ValueOf(fn)

	validateSliceAndFnValues(sliceValue, fnValue)

	for i := 0; i < sliceValue.Len(); i++ {
		fnValue.Call([]reflect.Value{sliceValue.Index(i)})
	}
}

func goEach(inMap map[int]interface{}, fnValue reflect.Value, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, value := range inMap {
		fnValue.Call([]reflect.Value{reflect.ValueOf(value)})
	}
}
