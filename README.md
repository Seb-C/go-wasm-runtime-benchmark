# Golang WASM runtimes benchmark

This repository is a benchmark of all the wasm (as far as I know) runtimes in Golang.

It compares the time required to call a single function via wasm.

The purpose of this benchmark is to compare the low level calls, meaning that it does not include any of the higher-level logic that would be needed for more advanced applications.

Especially, this does not include WASI or WASIX.

## Result

- `add` benchmarks a function that adds two numbers
- `fibonacci` is self-explanatory
- `lifecycle` benchmarks the startup and shutdown (together) times of the runtime

```
cpu: AMD Ryzen 5 5600 6-Core Processor              

Benchmark/add/wasmedge-12                1266534              1014 ns/op
Benchmark/add/wasmer-12                  1317753               858.0 ns/op
Benchmark/add/wasmtime-12                 601758              2043 ns/op
Benchmark/add/wazero-12                 20905545                57.21 ns/op

Benchmark/fibonacci/wasmedge-12           223616              5096 ns/op
Benchmark/fibonacci/wasmer-12            1525473               785.5 ns/op
Benchmark/fibonacci/wasmtime-12           690450              1931 ns/op
Benchmark/fibonacci/wazero-12           11451154               102.2 ns/op

Benchmark/lifecycle/wazero-12                423           2818693 ns/op
Benchmark/lifecycle/wasmedge-12             2967            388854 ns/op
Benchmark/lifecycle/wasmer-12                309           3820975 ns/op
Benchmark/lifecycle/wasmtime-12              309           3772255 ns/op
```

## How to run it?

First, build the wasm app:

```
docker run --rm -v "$PWD":/app -it $(docker build -q -f ./docker/build-wasm.Dockerfile .)
```

Second, run the benchmark:

```
docker build --tag go-wasm-runtime-benchmark -f ./docker/benchmark.Dockerfile . && docker run go-wasm-runtime-benchmark
```

## Why Rust instead of Go in the wasm file?

Because at the moment, there is no Go compiler that can output pure wasm and that supports both importing and exporting.
- The native compiler has a `go:wasmimport` directive since v1.21, but no `go:wasmimport`, ans until now it focused on giving access to specific JS functions for a browser environment
- Tinygo currently does not allow to output pure WASM, and actually always include WASI functions. Since my goal with this repository is to compare things without overhead and with a completely sandboxed code, this is not appropriate.

## TODO

- Benchmark imported functions
- Benchmark memory accesses
