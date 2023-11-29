package benchmark

import (
	"github.com/bytecodealliance/wasmtime-go/v15"
	"testing"
)

func initWasmTime(b *testing.B) (
	add func(x, y int64) int64,
	fibonacci func(x int64) int64,
	onClose func(),
) {
	wasmFile := getWasmFile(b)

	store := wasmtime.NewStore(wasmtime.NewEngine())

	module, err := wasmtime.NewModule(store.Engine, wasmFile)
	if err != nil {
		b.Error("Failed to create wasmtime module:", err)
	}

	instance, err := wasmtime.NewInstance(store, module, nil)
	if err != nil {
		b.Error("Failed to create wasmtime instance:", err)
	}

	addFunction := instance.GetFunc(store, addFunctionName)
	if addFunction == nil {
		b.Error("Add function not found")
	}

	fibonacciFunction := instance.GetFunc(store, fibonacciFunctionName)
	if fibonacciFunction == nil {
		b.Error("Fibonacci function not found")
	}

	add = func(x, y int64) int64 {
		result, err := addFunction.Call(store, x, y)
		if err != nil {
			b.Error("Failed to call add function:", err)
		}

		return result.(int64)

	}
	fibonacci = func(x int64) int64 {
		result, err := fibonacciFunction.Call(store, x)
		if err != nil {
			b.Error("Failed to call fibonacci function:", err)
		}

		return result.(int64)
	}
	onClose = func() {}

	return add, fibonacci, onClose
}

func BenchmarkWasmTime(b *testing.B) {
	add, fibonacci, onClose := initWasmTime(b)
	defer onClose()

	b.Run("add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = add(1, 2)
		}
	})
	b.Run("fibonacci", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fibonacci(42)
		}
	})
}
