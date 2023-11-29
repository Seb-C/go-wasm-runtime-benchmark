package benchmark

import (
	"github.com/second-state/WasmEdge-go/wasmedge"
	"testing"
)

func BenchmarkWasmEdge(b *testing.B) {
	wasmFile := getWasmFile(b)

	wasmedge.SetLogErrorLevel()

	config := wasmedge.NewConfigure(wasmedge.WASI)
	defer config.Release()

	vm := wasmedge.NewVMWithConfig(config)
	defer vm.Release()

	if err := vm.LoadWasmBuffer(wasmFile); err != nil {
		b.Error("Failed to load wasm file:", err)
	}
	if err := vm.Validate(); err != nil {
		b.Error("Failed to validate wasm file:", err)
	}
	if err := vm.Instantiate(); err != nil {
		b.Error("Failed to instantiate wasm file:", err)
	}

	function := vm.GetActiveModule().FindFunction(functionName)
	if function == nil {
		b.Error("Function not found")
		return
	}

	b.Run("wasmedge", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := vm.Execute(functionName, 1, 2)
			if err != nil {
				b.Error("Failed to call function:", err)
			}
		}
	})
}
