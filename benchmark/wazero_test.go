package benchmark

import (
	"context"
	"github.com/tetratelabs/wazero"
	"testing"
)

func BenchmarkWazero(b *testing.B) {
	wasmFile := getWasmFile(b)

	ctx := context.Background()
	runtime := wazero.NewRuntime(ctx)
	defer runtime.Close(ctx)

	module, err := runtime.Instantiate(ctx, wasmFile)
	if err != nil {
		b.Error("Failed to instantiate module:", err)
		return
	}

	function := module.ExportedFunction(functionName)

	b.Run("wazero", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := function.Call(ctx, 1, 2); err != nil {
				b.Error("Failed to call function:", err)
				return
			}
		}
	})
}
