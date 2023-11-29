FROM golang:1.21

WORKDIR /app

RUN apt-get install gcc
RUN curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.13.4

COPY ./target ./target
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go mod download

COPY ./benchmark ./benchmark

CMD . $HOME/.wasmedge/env && go test -bench=. ./...
