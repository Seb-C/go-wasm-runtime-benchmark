# Golang WASM runtimes benchmark

## How to run it?

First, build the wasm app:

```
docker run --rm -v "$PWD":/app -it $(docker build -q -f ./docker/build-wasm.Dockerfile .)
```

Second, run the benchmark:

```
docker build --tag go-wasm-runtime-benchmark -f ./docker/benchmark.Dockerfile . && docker run go-wasm-runtime-benchmark
```
