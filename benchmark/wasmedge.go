package benchmark

import (
	"os"
	"testing"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

func initWasmEdge(tb testing.TB, wasmFile []byte) (
	add func(x, y int64) int64,
	fibonacci func(x int64) int64,
	onClose func(),
) {
	wasmedge.SetLogErrorLevel()

	config := wasmedge.NewConfigure(wasmedge.WASI)
	vm := wasmedge.NewVMWithConfig(config)

	var wasi = vm.GetImportModule(wasmedge.WASI)
	wasi.InitWasi(os.Args[1:], os.Environ(), []string{".:."})

	if err := vm.LoadWasmBuffer(wasmFile); err != nil {
		tb.Error("Failed to load wasm file:", err)
	}
	if err := vm.Validate(); err != nil {
		tb.Error("Failed to validate wasm file:", err)
	}
	if err := vm.Instantiate(); err != nil {
		tb.Error("Failed to instantiate wasm file:", err)
	}

	executor := vm.GetExecutor()
	module := vm.GetActiveModule()
	addFunction := module.FindFunction(addFunctionName)
	fibonacciFunction := module.FindFunction(fibonacciFunctionName)

	add = func(x, y int64) int64 {
		result, err := executor.Invoke(addFunction, x, y)
		if err != nil {
			tb.Error("Failed to call add function:", err)
		}

		if len(result) != 1 {
			tb.Errorf("Expected 1 return param, got %d: %s", len(result), err)
		}

		return result[0].(int64)
	}
	fibonacci = func(x int64) int64 {
		result, err := executor.Invoke(fibonacciFunction, x)
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
