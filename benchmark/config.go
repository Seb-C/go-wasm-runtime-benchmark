package benchmark

import "testing"

const addFunctionName = "add"
const fibonacciFunctionName = "fibonacci"

type runtime func (tb testing.TB, wasmFile []byte) (
	add func(x, y int64) int64,
	fibonacci func(x int64) int64,
	onClose func(),
)

type testCase struct{
	name string
	wasmFilePath string
	runtimes map[string]runtime
}

var testCases = []testCase{
	testCase{
		name: "rust-wasm",
		wasmFilePath: "./rust.wasm",
		runtimes: map[string]runtime {
			"wasmedge": initWasmEdge,
			"wasmer": initWasmer,
			"wasmtime": initWasmTime,
			"wazero": initWazero,
		},
	},
	testCase{
		name: "tinygo-wasi",
		wasmFilePath: "./tinygo.wasi",
		runtimes: map[string]runtime {
			"wasmedge": initWasmEdge,
			// "wasmer": initWasmer, // I could not get wasi to work on this one
			"wasmtime": initWasmTime,
			"wazero": initWazero,
		},
	},
}
