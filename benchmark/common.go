package benchmark

import (
	"os"
	"testing"
)

const addFunctionName = "add"
const fibonacciFunctionName = "fibonacci"

func getWasmFile(tb testing.TB) []byte {
	tb.Helper()

	wasmFilePath := "./rust.wasm"
	wasmFile, err := os.ReadFile(wasmFilePath)
	if err != nil {
		tb.Error("Failed to get wasm file:", err)
	}

	return wasmFile
}
