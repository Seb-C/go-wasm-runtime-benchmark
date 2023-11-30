package benchmark

import (
	wasmer "github.com/wasmerio/wasmer-go/wasmer"
	"testing"
)

func initWasmer(tb testing.TB) (
	add func(x, y int64) int64,
	fibonacci func(x int64) int64,
	onClose func(),
) {
	wasmFile := getWasmFile(tb)

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	module, err := wasmer.NewModule(store, wasmFile)
	if err != nil {
		tb.Error("Failed to create module:", err)
	}

	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		tb.Error("Failed to create instance:", err)
	}

	addFunction, err := instance.Exports.GetFunction(addFunctionName)
	if err != nil {
		tb.Error("Could not find add function:", err)
	}

	fibonacciFunction, err := instance.Exports.GetFunction(fibonacciFunctionName)
	if err != nil {
		tb.Error("Could not find fibonacci function:", err)
	}

	add = func(x, y int64) int64 {
		result, err := addFunction(x, y)
		if err != nil {
			tb.Error("Failed to call add function:", err)
		}

		return result.(int64)
	}
	fibonacci = func(x int64) int64 {
		result, err := fibonacciFunction(x)
		if err != nil {
			tb.Error("Failed to call fibonacci function:", err)
		}

		return result.(int64)
	}
	onClose = func() {
		defer instance.Close()
		defer module.Close()
		defer store.Close()
	}

	return add, fibonacci, onClose
}
