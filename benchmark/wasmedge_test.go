package benchmark

import (
	"github.com/second-state/WasmEdge-go/wasmedge"
	"testing"
)

func initWasmEdge(tb testing.TB) (
	add func(x, y int64) int64,
	fibonacci func(x int64) int64,
	onClose func(),
) {
	wasmFile := getWasmFile(tb)

	wasmedge.SetLogErrorLevel()

	config := wasmedge.NewConfigure(wasmedge.WASI)

	vm := wasmedge.NewVMWithConfig(config)

	if err := vm.LoadWasmBuffer(wasmFile); err != nil {
		tb.Error("Failed to load wasm file:", err)
	}
	if err := vm.Validate(); err != nil {
		tb.Error("Failed to validate wasm file:", err)
	}
	if err := vm.Instantiate(); err != nil {
		tb.Error("Failed to instantiate wasm file:", err)
	}

	add = func(x, y int64) int64 {
		result, err := vm.Execute(addFunctionName, x, y)
		if err != nil {
			tb.Error("Failed to call add function:", err)
		}

		if len(result) != 1 {
			tb.Errorf("Expected 1 return param, got %d: %s", len(result), err)
		}

		return result[0].(int64)
	}
	fibonacci = func(x int64) int64 {
		result, err := vm.Execute(fibonacciFunctionName, x)
		if err != nil {
			tb.Error("Failed to call fibonacci function:", err)
		}

		if len(result) != 1 {
			tb.Errorf("Expected 1 return param, got %d: %s", len(result), err)
		}

		return result[0].(int64)
	}
	onClose = func() {
		vm.Release()
		config.Release()
	}

	return add, fibonacci, onClose
}
