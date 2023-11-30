package main

func main() {}

//export add
func add(x int64, y int64) int64 {
	return x + y
}

// From https://www.thepolyglotdeveloper.com/2016/12/fibonacci-sequence-printed-golang/
//export fibonacci
func fibonacci(n int64) int64 {
	f := make([]int64, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := int64(2); i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
