concurrently
=====
Simple, easy concurrent processing in Go

Overview
--------

`concurrently` makes it easy to concurrently process collections using higher-order functions like Map and Filter.

For example, if you have a `urls` array and a `getHttpResponse` function, you can distribute the requests across 16 goroutines with a single line:

```go
responses := concurrently.Map(urls, getHttpResponse, 16).([]http.Response)
```

You could also concurrently filter a collection across 64 goroutines:
```go
activeUsers := concurrently.Filter(users, isUserActive, 64).([]User)
```

Or call a function without requiring a return value:
```go
concurrently.Each(images, resizeImage, 16)
```

For IO-bound operations, the execution time will be inversely proportional to the number of goroutines. If an operation averages 117 ms with 1 goroutine, it might average 11.8 ms with 10 goroutines or 1.51 ms with 100 goroutines:

```
BenchmarkMap001       20   117447904 ns/op
BenchmarkMap010      100   11803551 ns/op
BenchmarkMap100     2000   1511264 ns/op
BenchmarkMapSerial    20   116438290 ns/op
```

Usage
-----

For full documentation, see the [GoDoc documentation](https://godoc.org/github.com/tombenner/concurrently).

### Map

For example, to square every int in a slice of ints:

```go
// Create a slice of ints
numbers := make([]int, 10)
for i := range numbers {
  numbers[i] = i
}

// Define the function
func square(number int) int {
  return number * number
}

// Perform the mapping, distributing it across 8 goroutines
squaredNumbers := concurrently.Map(numbers, square, 8).([]int)
```

The function should have one parameter (of any type) and one return value (of any type, not necessarily the same type as the parameter).

### Filter

For example, to retrieve the even ints from a slice of ints:

```go
// Create a slice of ints
numbers := make([]int, 10)
for i := range numbers {
  numbers[i] = i
}

// Define the function
func isEvenNumber(number int) bool {
  return int(math.Mod(float64(number), float64(2))) == 0
}

// Perform the filtering, distributing it across 8 goroutines
evenNumbers := concurrently.Filter(numbers, isEvenNumber, 8).([]int)
```

The function should have one parameter (of any type) and return a bool.

### Each

For example, to log a number of messages:

```go
// Create an array of messages
messages := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

// Define the function
func logMessage(message string) {
  log.Print(message)
}

// Log the messages
concurrently.Each(messages, logMessage, 4)
```

License
-------

concurrently is released under the MIT License. Please see the MIT-LICENSE file for details.
