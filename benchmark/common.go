package benchmark

import (
	"os"
	"testing"
)

const functionName = "add"

func getWasmFile(b *testing.B) []byte {
	b.Helper()

	wasmFilePath := "../target/wasm32-unknown-unknown/release/go-wasm-runtime-benchmark.wasm"
	wasmFile, err := os.ReadFile(wasmFilePath)
	if err != nil {
		b.Error("Failed to get wasm file:", err)
	}

	return wasmFile
}
