package benchmark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var runtimes = map[string]func (tb testing.TB) (
	add func(x, y int64) int64,
	fibonacci func(x int64) int64,
	onClose func(),
) {
	"wasmedge": initWasmEdge,
	"wasmer": initWasmer,
	"wasmtime": initWasmTime,
	"wazero": initWazero,
}

func Test(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		for name, initRuntime := range runtimes {
			add, _, onClose := initRuntime(t)
			t.Run(name, func(t *testing.T) {
				assert.Equal(t, int64(3), add(1, 2))
				assert.Equal(t, int64(-100), add(23, -123))
			})
			onClose()
		}
	})

	t.Run("fibonacci", func(t *testing.T) {
		for name, initRuntime := range runtimes {
			_, fibonacci, onClose := initRuntime(t)
			t.Run(name, func(t *testing.T) {
				assert.Equal(t, int64(1), fibonacci(1))
				assert.Equal(t, int64(55), fibonacci(10))
			})
			onClose()
		}
	})
}

func Benchmark(b *testing.B) {
	b.Run("add", func(b *testing.B) {
		for name, initRuntime := range runtimes {
			add, _, onClose := initRuntime(b)
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = add(1, 2)
				}
			})
			onClose()
		}
	})

	b.Run("fibonacci", func(b *testing.B) {
		for name, initRuntime := range runtimes {
			_, fibonacci, onClose := initRuntime(b)
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = fibonacci(123)
				}
			})
			onClose()
		}
	})

	b.Run("lifecycle", func(b *testing.B) {
		for name, initRuntime := range runtimes {
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _, onClose := initRuntime(b)
					onClose()
				}
			})
		}
	})
}
