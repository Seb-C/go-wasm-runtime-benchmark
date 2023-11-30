package benchmark

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			wasmFile, err := os.ReadFile(testCase.wasmFilePath)
			if err != nil {
				t.Error("Failed to get wasm file:", err)
			}

			t.Run("add", func(t *testing.T) {
				for name, initRuntime := range testCase.runtimes {
					add, _, onClose := initRuntime(t, wasmFile)
					t.Run(name, func(t *testing.T) {
						assert.Equal(t, int64(3), add(1, 2))
						assert.Equal(t, int64(-100), add(23, -123))
					})
					onClose()
				}
			})

			t.Run("fibonacci", func(t *testing.T) {
				for name, initRuntime := range testCase.runtimes {
					_, fibonacci, onClose := initRuntime(t, wasmFile)
					t.Run(name, func(t *testing.T) {
						assert.Equal(t, int64(1), fibonacci(1))
						assert.Equal(t, int64(55), fibonacci(10))
					})
					onClose()
				}
			})
		})
	}
}

func Benchmark(b *testing.B) {
	for _, testCase := range testCases {
		b.Run(testCase.name, func(b *testing.B) {
			wasmFile, err := os.ReadFile(testCase.wasmFilePath)
			if err != nil {
				b.Error("Failed to get wasm file:", err)
			}

			b.Run("add", func(b *testing.B) {
				for name, initRuntime := range testCase.runtimes {
					add, _, onClose := initRuntime(b, wasmFile)
					b.Run(name, func(b *testing.B) {
						for i := 0; i < b.N; i++ {
							_ = add(1, 2)
						}
					})
					onClose()
				}
			})

			b.Run("fibonacci", func(b *testing.B) {
				for name, initRuntime := range testCase.runtimes {
					_, fibonacci, onClose := initRuntime(b, wasmFile)
					b.Run(name, func(b *testing.B) {
						for i := 0; i < b.N; i++ {
							_ = fibonacci(123)
						}
					})
					onClose()
				}
			})

			b.Run("lifecycle", func(b *testing.B) {
				for name, initRuntime := range testCase.runtimes {
					b.Run(name, func(b *testing.B) {
						for i := 0; i < b.N; i++ {
							_, _, onClose := initRuntime(b, wasmFile)
							onClose()
						}
					})
				}
			})
		})
	}
}
