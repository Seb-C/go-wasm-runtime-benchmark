package benchmark

import (
	"github.com/bytecodealliance/wasmtime-go/v15"
	"testing"
)

func initWasmTime(tb testing.TB) (
	add func(x, y int64) int64,
	fibonacci func(x int64) int64,
	onClose func(),
) {
	wasmFile := getWasmFile(tb)

	store := wasmtime.NewStore(wasmtime.NewEngine())

	module, err := wasmtime.NewModule(store.Engine, wasmFile)
	if err != nil {
		tb.Error("Failed to create wasmtime module:", err)
	}

	instance, err := wasmtime.NewInstance(store, module, nil)
	if err != nil {
		tb.Error("Failed to create wasmtime instance:", err)
	}

	addFunction := instance.GetFunc(store, addFunctionName)
	if addFunction == nil {
		tb.Error("Add function not found")
	}

	fibonacciFunction := instance.GetFunc(store, fibonacciFunctionName)
	if fibonacciFunction == nil {
		tb.Error("Fibonacci function not found")
	}

	add = func(x, y int64) int64 {
		result, err := addFunction.Call(store, x, y)
		if err != nil {
			tb.Error("Failed to call add function:", err)
		}

		return result.(int64)

	}
	fibonacci = func(x int64) int64 {
		result, err := fibonacciFunction.Call(store, x)
		if err != nil {
			tb.Error("Failed to call fibonacci function:", err)
		}

		return result.(int64)
	}
	onClose = func() {}

	return add, fibonacci, onClose
}
