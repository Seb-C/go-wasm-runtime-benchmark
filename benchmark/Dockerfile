FROM golang:1.21

WORKDIR /app

RUN apt-get install gcc
RUN curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.13.4

COPY ./rust-wasm/target/wasm32-unknown-unknown/release/go-wasm-runtime-benchmark.wasm ./rust.wasm

COPY ./benchmark/go.mod ./go.mod
COPY ./benchmark/go.sum ./go.sum

RUN go mod download

COPY ./benchmark/*.go ./

CMD . $HOME/.wasmedge/env && go test -bench=. ./...
