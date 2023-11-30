MAKEFLAGS += --always-make
.DEFAULT_GOAL := all

all: rust-wasm tinygo-wasi benchmark

rust-wasm:
	docker build --tag go-wasm-runtime-benchmark:rust-wasm -f ./rust-wasm/Dockerfile .
	docker run --rm -v "$$PWD/rust-wasm":/app -it go-wasm-runtime-benchmark:rust-wasm

tinygo-wasi:
	docker run --rm -v "$$PWD/tinygo-wasi:/app" -w /app tinygo/tinygo:0.30.0 tinygo build -o ./tinygo.wasi -target=wasi ./main.go

benchmark:
	docker build --tag go-wasm-runtime-benchmark:benchmark -f ./benchmark/Dockerfile .
	docker run --rm go-wasm-runtime-benchmark:benchmark
