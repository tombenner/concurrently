package concurrently

import (
	"log"
	"reflect"
	"sync"
)

// Map applies `fn` to each element of `slice` and returns the result.
// `workersCount` is the number of goroutines that the work will be distributed across.
// You'll usually want to cast the returned interface{}: `Map(numbers, doubleNumber, 4).([]int)`.
func Map(slice interface{}, fn interface{}, workersCount int) interface{} {
	sliceValue := reflect.ValueOf(slice)
	fnValue := reflect.ValueOf(fn)

	validateSliceAndFnValues(sliceValue, fnValue)
	validateMapFuncValue(fnValue)

	maps := makeWorkerMaps(sliceValue, workersCount)

	var wg sync.WaitGroup
	wg.Add(workersCount)
	channel := make(chan map[int]reflect.Value)

	for i := 0; i < workersCount; i++ {
		go goMap(maps[i], fnValue, &wg, channel)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	outSliceType := reflect.SliceOf(fnValue.Type().Out(0))
	outSlice := reflect.MakeSlice(outSliceType, sliceValue.Len(), sliceValue.Len())

	for outMap := range channel {
		for i, value := range outMap {
			outSlice.Index(i).Set(value)
		}
	}

	return outSlice.Interface()
}

// MapSerial is nearly identical to Map, but it performs the work serially rather than concurrently.
// (It doesn't use goroutines.)
func MapSerial(slice interface{}, fn interface{}) interface{} {
	sliceValue := reflect.ValueOf(slice)
	fnValue := reflect.ValueOf(fn)

	validateSliceAndFnValues(sliceValue, fnValue)
	validateMapFuncValue(fnValue)

	outSliceType := reflect.SliceOf(fnValue.Type().Out(0))
	outSlice := reflect.MakeSlice(outSliceType, sliceValue.Len(), sliceValue.Len())

	for i := 0; i < sliceValue.Len(); i++ {
		outElem := fnValue.Call([]reflect.Value{sliceValue.Index(i)})[0]
		outSlice.Index(i).Set(outElem)
	}

	return outSlice.Interface()
}

func goMap(inMap map[int]interface{}, fnValue reflect.Value, wg *sync.WaitGroup, channel chan map[int]reflect.Value) {
	defer wg.Done()

	outMap := make(map[int]reflect.Value)

	for i, value := range inMap {
		outMap[i] = fnValue.Call([]reflect.Value{reflect.ValueOf(value)})[0]
	}

	channel <- outMap
}

func validateMapFuncValue(fnValue reflect.Value) {
	fnType := fnValue.Type()

	if fnType.NumOut() != 1 {
		log.Panicf("`fn` should return 1 value, but it returns %d values", fnType.NumOut())
	}
}
