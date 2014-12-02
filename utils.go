package concurrently

import (
	"log"
	"math"
	"reflect"
)

func makeWorkerMaps(sliceValue reflect.Value, workersCount int) []map[int]interface{} {
	maps := make([]map[int]interface{}, workersCount)

	for i := 0; i < workersCount; i++ {
		maps[i] = make(map[int]interface{})
	}

	for i := 0; i < sliceValue.Len(); i++ {
		mapsIndex := int(math.Mod(float64(i), float64(workersCount)))
		maps[mapsIndex][i] = sliceValue.Index(i).Interface()
	}

	return maps
}

func validateSliceAndFnValues(sliceValue reflect.Value, fnValue reflect.Value) {
	fnType := fnValue.Type()
	sliceType := sliceValue.Type()

	if fnType.Kind() != reflect.Func {
		log.Panicf("`fn` should be %s, but got %s", reflect.Func, fnType.Kind())
	}
	if fnType.NumIn() != 1 {
		log.Panicf("`fn` should have 1 parameter, but it has %d parameters", fnType.NumIn())
	}

	if sliceType.Kind() != reflect.Slice {
		log.Panicf("`slice` should be %s, but got %s", reflect.Slice, sliceType.Kind())
	}
	if sliceType.Elem() != fnType.In(0) {
		log.Panicf("type of `fn`'s parameter should be %s, but slice contains %s", fnType.In(0), sliceType.Elem())
	}
}
