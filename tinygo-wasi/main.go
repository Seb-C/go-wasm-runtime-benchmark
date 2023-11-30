package main

import "fmt"

func main() {}

//export add
func add(x int64, y int64) int64 {
	return x + y
}

//export fibonacci
func fibonacci(n int64) int64 {
	if n < 0 {
		panic(fmt.Errorf("%d is negative!", n))
	} else if n == 0 {
		panic("zero is not a right argument to fibonacci()!")
	} else if n == 1 {
		return 1;
	}

	var sum int64 = 0
	var last int64 = 0
	var curr int64 = 1
	var i int64
	for i = 1; i < n; i++ {
		sum = last + curr
		last = curr
		curr = sum
	}

	return sum
}
