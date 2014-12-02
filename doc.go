// Copyright 2014 Tom Benner

// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

/*
Package concurrently provides simple concurrent higher-order functions, like Map and Filter.

For example, if you have a `urls` array and a `getHttpResponse` function, you can distribute the requests across 16 goroutines with a single line:

  responses := concurrently.Map(urls, getHttpResponse, 16).([]http.Response)

You could also concurrently filter a collection across 64 goroutines:

  activeUsers := concurrently.Filter(users, isUserActive, 64).([]User)

Or call a function without requiring a return value:

  concurrently.Each(images, resizeImage, 16)

For IO-bound operations, the execution time will be inversely proportional to the number of goroutines. If an operation averages 117 ms with 1 goroutine, it might average 11.8 ms with 10 goroutines or 1.51 ms with 100 goroutines:

  BenchmarkMap001       20   117447904 ns/op
  BenchmarkMap010      100   11803551 ns/op
  BenchmarkMap100     2000   1511264 ns/op
  BenchmarkMapSerial    20   116438290 ns/op

Map

For example, to square every int in a slice of ints:

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

The function should have one parameter (of any type) and one return value (of any type, not necessarily the same type as the parameter).

Filter

For example, to retrieve the even ints from a slice of ints:

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

The function should have one parameter (of any type) and return a bool.

Each

For example, to log a number of messages:

  // Create an array of messages
  messages := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

  // Define the function
  func logMessage(message string) {
    log.Print(message)
  }

  // Log the messages
  concurrently.Each(messages, logMessage, 4)

*/
package concurrently
