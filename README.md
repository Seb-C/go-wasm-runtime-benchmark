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

Benchmark/rust-wasm/add/wasmedge-12              1964247               676.7 ns/op
Benchmark/rust-wasm/add/wasmer-12                1341111               835.6 ns/op
Benchmark/rust-wasm/add/wasmtime-12               683862              1903 ns/op
Benchmark/rust-wasm/add/wazero-12               17825775                57.03 ns/op

Benchmark/rust-wasm/fibonacci/wasmedge-12         251000              4651 ns/op
Benchmark/rust-wasm/fibonacci/wasmer-12          1446950               779.5 ns/op
Benchmark/rust-wasm/fibonacci/wasmtime-12         638666              1978 ns/op
Benchmark/rust-wasm/fibonacci/wazero-12         12426373               101.4 ns/op

Benchmark/rust-wasm/lifecycle/wasmedge-12           2926            367483 ns/op
Benchmark/rust-wasm/lifecycle/wasmer-12              309           3842560 ns/op
Benchmark/rust-wasm/lifecycle/wasmtime-12            279           4226500 ns/op
Benchmark/rust-wasm/lifecycle/wazero-12              369           2926465 ns/op

Benchmark/tinygo-wasi/add/wasmedge-12            1741168               704.5 ns/op
Benchmark/tinygo-wasi/add/wasmtime-12             582134              1914 ns/op
Benchmark/tinygo-wasi/add/wazero-12             19349671                58.45 ns/op

Benchmark/tinygo-wasi/fibonacci/wasmedge-12        95944             12645 ns/op
Benchmark/tinygo-wasi/fibonacci/wasmtime-12       694854              1937 ns/op
Benchmark/tinygo-wasi/fibonacci/wazero-12        5007806               212.2 ns/op

Benchmark/tinygo-wasi/lifecycle/wasmedge-12          510           2285191 ns/op
Benchmark/tinygo-wasi/lifecycle/wasmtime-12           27          39080134 ns/op
Benchmark/tinygo-wasi/lifecycle/wazero-12             38          33822371 ns/op
```

## TODO

- Benchmark a function that uses a system call
