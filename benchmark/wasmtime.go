package benchmark

import (
	"github.com/bytecodealliance/wasmtime-go/v15"
	"testing"
)

func initWasmTime(tb testing.TB, wasmFile []byte) (
	add func(x, y int64) int64,
	fibonacci func(x int64) int64,
	onClose func(),
) {
	engine := wasmtime.NewEngine()
	store := wasmtime.NewStore(engine)
	linker := wasmtime.NewLinker(engine)

	module, err := wasmtime.NewModule(engine, wasmFile)
	if err != nil {
		tb.Error("Failed to create wasmtime module:", err)
	}

	if err := linker.DefineWasi(); err != nil {
		tb.Error("Failed to define wasi:", err)
	}

	instance, err := linker.Instantiate(store, module)
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
