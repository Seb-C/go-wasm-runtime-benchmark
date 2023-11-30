MAKEFLAGS += --always-make
.DEFAULT_GOAL := all

all: rust-wasm benchmark

rust-wasm:
	docker build --tag go-wasm-runtime-benchmark:rust-wasm -f ./rust-wasm/Dockerfile .
	docker run --rm -v "$$PWD/rust-wasm":/app -it go-wasm-runtime-benchmark:rust-wasm

benchmark:
	docker build --tag go-wasm-runtime-benchmark:benchmark -f ./benchmark/Dockerfile .
	docker run --rm go-wasm-runtime-benchmark:benchmark