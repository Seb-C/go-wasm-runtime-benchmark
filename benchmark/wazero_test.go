package benchmark

import (
	"context"
	"github.com/tetratelabs/wazero"
	"testing"
)

func initWazero(b *testing.B) (
	add func(x, y int64) int64,
	fibonacci func(x int64) int64,
	onClose func(),
) {
	wasmFile := getWasmFile(b)

	ctx := context.Background()
	runtime := wazero.NewRuntime(ctx)

	module, err := runtime.Instantiate(ctx, wasmFile)
	if err != nil {
		b.Error("Failed to instantiate module:", err)
	}

	addFunction := module.ExportedFunction(addFunctionName)
	fibonacciFunction := module.ExportedFunction(fibonacciFunctionName)

	add = func(x, y int64) int64 {
		result, err := addFunction.Call(ctx, uint64(x), uint64(y))
		if err != nil {
			b.Error("Failed to call add function:", err)
		}

		if len(result) != 1 {
			b.Errorf("Expected 1 return param, got %d: %s", len(result), err)
		}

		return int64(result[0])
	}
	fibonacci = func(x int64) int64 {
		result, err := fibonacciFunction.Call(ctx, uint64(x))
		if err != nil {
			b.Error("Failed to call fibonacci function:", err)
		}

		if len(result) != 1 {
			b.Errorf("Expected 1 return param, got %d: %s", len(result), err)
		}

		return int64(result[0])
	}
	onClose = func() {
		runtime.Close(ctx)
	}

	return add, fibonacci, onClose
}
