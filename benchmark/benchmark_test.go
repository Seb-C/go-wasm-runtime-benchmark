package benchmark

import (
	"testing"
)

var benchmarks = map[string]func (b *testing.B) (
	add func(x, y int64) int64,
	fibonacci func(x int64) int64,
	onClose func(),
) {
	"wasmedge": initWasmEdge,
	"wasmer": initWasmer,
	"wasmtime": initWasmTime,
	"wazero": initWazero,
}

func Benchmark(b *testing.B) {
	b.Run("add", func(b *testing.B) {
		for benchmarkName, initBenchmark := range benchmarks {
			add, _, onClose := initBenchmark(b)
			b.Run(benchmarkName, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = add(1, 2)
				}
			})
			onClose()
		}
	})

	b.Run("fibonacci", func(b *testing.B) {
		for benchmarkName, initBenchmark := range benchmarks {
			_, fibonacci, onClose := initBenchmark(b)
			b.Run(benchmarkName, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = fibonacci(123)
				}
			})
			onClose()
		}
	})

	b.Run("lifecycle", func(b *testing.B) {
		for benchmarkName, initBenchmark := range benchmarks {
			b.Run(benchmarkName, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _, onClose := initBenchmark(b)
					onClose()
				}
			})
		}
	})
}
