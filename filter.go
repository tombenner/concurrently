package concurrently

import (
	"log"
	"reflect"
	"sync"
)

// Filter returns a slice containing the elements of `slice` for which `fn` returns true.
// `workersCount` is the number of goroutines that the work will be distributed across.
// You'll usually want to cast the returned interface{}: `Filter(numbers, isEvenNumber, 4).([]int)`.
func Filter(slice interface{}, fn interface{}, workersCount int) interface{} {
	sliceValue := reflect.ValueOf(slice)
	fnValue := reflect.ValueOf(fn)

	validateSliceAndFnValues(sliceValue, fnValue)
	validateFilterFuncValue(fnValue)

	maps := makeWorkerMaps(sliceValue, workersCount)

	var wg sync.WaitGroup
	wg.Add(workersCount)
	channel := make(chan map[int]reflect.Value)

	for i := 0; i < workersCount; i++ {
		go goFilter(maps[i], fnValue, &wg, channel)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	outSliceType := reflect.SliceOf(sliceValue.Index(0).Type())
	outSlice := reflect.MakeSlice(outSliceType, 0, sliceValue.Len())

	filteredMap := make(map[int]reflect.Value, sliceValue.Len())
	for workerMap := range channel {
		for i, value := range workerMap {
			filteredMap[i] = value
		}
	}

	for i := 0; i < sliceValue.Len(); i++ {
		if value, ok := filteredMap[i]; ok {
			outSlice = reflect.Append(outSlice, value)
		}
	}

	return outSlice.Interface()
}

// FilterSerial is nearly identical to Filter, but it performs the work serially rather than concurrently.
// (It doesn't use goroutines.)
func FilterSerial(slice interface{}, fn interface{}) interface{} {
	sliceValue := reflect.ValueOf(slice)
	fnValue := reflect.ValueOf(fn)

	validateSliceAndFnValues(sliceValue, fnValue)
	validateFilterFuncValue(fnValue)

	outSliceType := reflect.SliceOf(sliceValue.Index(0).Type())
	outSlice := reflect.MakeSlice(outSliceType, 0, sliceValue.Len())

	for i := 0; i < sliceValue.Len(); i++ {
		if fnValue.Call([]reflect.Value{sliceValue.Index(i)})[0].Bool() {
			outSlice = reflect.Append(outSlice, sliceValue.Index(i))
		}
	}

	return outSlice.Interface()
}

func goFilter(inMap map[int]interface{}, fnValue reflect.Value, wg *sync.WaitGroup, channel chan map[int]reflect.Value) {
	defer wg.Done()

	workerMap := make(map[int]reflect.Value)

	for i, value := range inMap {
		valueValue := reflect.ValueOf(value)
		if fnValue.Call([]reflect.Value{valueValue})[0].Bool() {
			workerMap[i] = valueValue
		}
	}

	channel <- workerMap
}

func validateFilterFuncValue(fnValue reflect.Value) {
	if fnValue.Type().Out(0) != reflect.TypeOf(true) {
		log.Panicf("`fn` should return a bool type, but it return a %d type", fnValue.Type().Out(0))
	}
}
