package benchmark

import (
	"github.com/bytecodealliance/wasmtime-go/v15"
	"testing"
)

func BenchmarkWasmTime(b *testing.B) {
	wasmFile := getWasmFile(b)

	store := wasmtime.NewStore(wasmtime.NewEngine())

	module, err := wasmtime.NewModule(store.Engine, wasmFile)
	if err != nil {
		b.Error("Failed to create wasmtime module:", err)
		return
	}

	instance, err := wasmtime.NewInstance(store, module, nil)
	if err != nil {
		b.Error("Failed to create wasmtime instance:", err)
		return
	}

	function := instance.GetFunc(store, functionName)
	if function == nil {
		b.Error("Function not found")
		return
	}

	b.Run("wasmtime", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := function.Call(store, 1, 2)
			if err != nil {
				b.Error("Failed to call function:", err)
				return
			}
		}
	})
}
