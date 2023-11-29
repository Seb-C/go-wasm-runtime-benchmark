package benchmark

import (
	wasmer "github.com/wasmerio/wasmer-go/wasmer"
	"testing"
)

func BenchmarkWasmer(b *testing.B) {
	wasmFile := getWasmFile(b)

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)
	defer store.Close()

	module, err := wasmer.NewModule(store, wasmFile)
	if err != nil {
		b.Error("Failed to create module:", err)
	}
	defer module.Close()

	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		b.Error("Failed to create instance:", err)
	}
	defer instance.Close()

	function, err := instance.Exports.GetFunction(functionName)
	if err != nil {
		b.Error("Could not find function:", err)
	}

	b.Run("wasmer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := function(1, 2)
			if err != nil {
				b.Error("Failed to call function:", err)
			}
		}
	})
}
