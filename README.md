# Golang WASM runtimes benchmark

This repository is a benchmark of all the wasm (that I could find) runtimes in Golang.

It compares the time required to call a single function via wasm.

There is one benchmark with pure WASM written in Rust, and one in Go that includes the WASI layer compiled by TinyGo.

## Result

- `add` benchmarks a function that adds two numbers
- `fibonacci` is self-explanatory
- `lifecycle` benchmarks the startup and shutdown (together) times of the runtime

```
cpu: AMD Ryzen 5 5600 6-Core Processor              

Benchmark/rust-wasm/add/wasmedge-12              1500884               881.3 ns/op
Benchmark/rust-wasm/add/wasmer-12                1000000              1036 ns/op
Benchmark/rust-wasm/add/wasmtime-12               442851              2581 ns/op
Benchmark/rust-wasm/add/wazero-12               12850506                92.53 ns/op

Benchmark/rust-wasm/fibonacci/wasmedge-12         272600              4753 ns/op
Benchmark/rust-wasm/fibonacci/wasmer-12          1000000              1031 ns/op
Benchmark/rust-wasm/fibonacci/wasmtime-12         573050              2122 ns/op
Benchmark/rust-wasm/fibonacci/wazero-12          7266033               145.0 ns/op

Benchmark/rust-wasm/lifecycle/wasmedge-12           3230            363073 ns/op
Benchmark/rust-wasm/lifecycle/wasmer-12              318           3733105 ns/op
Benchmark/rust-wasm/lifecycle/wasmtime-12            270           4231169 ns/op
Benchmark/rust-wasm/lifecycle/wazero-12              388           4055002 ns/op

Benchmark/tinygo-wasi/add/wasmedge-12            1305934               882.9 ns/op
Benchmark/tinygo-wasi/add/wasmtime-12             515164              2590 ns/op
Benchmark/tinygo-wasi/add/wazero-12             12304556                85.12 ns/op

Benchmark/tinygo-wasi/fibonacci/wasmedge-12        10000            129466 ns/op
Benchmark/tinygo-wasi/fibonacci/wasmtime-12       375784              3345 ns/op
Benchmark/tinygo-wasi/fibonacci/wazero-12         136599              8488 ns/op

Benchmark/tinygo-wasi/lifecycle/wasmedge-12         2462            407246 ns/op
Benchmark/tinygo-wasi/lifecycle/wasmtime-12          210           5703182 ns/op
Benchmark/tinygo-wasi/lifecycle/wazero-12            216           5212543 ns/op
```

## TODO

- Benchmark a function that uses a system call
