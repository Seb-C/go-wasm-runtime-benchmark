package benchmark

import (
	"context"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func initWazero(tb testing.TB, wasmFile []byte) (
	add func(x, y int64) int64,
	fibonacci func(x int64) int64,
	onClose func(),
) {
	ctx := context.Background()
	runtime := wazero.NewRuntime(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, runtime)

	module, err := runtime.Instantiate(ctx, wasmFile)
	if err != nil {
		tb.Error("Failed to instantiate module:", err)
	}

	addFunction := module.ExportedFunction(addFunctionName)
	fibonacciFunction := module.ExportedFunction(fibonacciFunctionName)

	add = func(x, y int64) int64 {
		result, err := addFunction.Call(ctx, uint64(x), uint64(y))
		if err != nil {
			tb.Error("Failed to call add function:", err)
		}

		if len(result) != 1 {
			tb.Errorf("Expected 1 return param, got %d: %s", len(result), err)
		}

		return int64(result[0])
	}
	fibonacci = func(x int64) int64 {
		result, err := fibonacciFunction.Call(ctx, uint64(x))
		if err != nil {
			tb.Error("Failed to call fibonacci function:", err)
		}

		if len(result) != 1 {
			tb.Errorf("Expected 1 return param, got %d: %s", len(result), err)
		}

		return int64(result[0])
	}
	onClose = func() {
		runtime.Close(ctx)
	}

	return add, fibonacci, onClose
}
