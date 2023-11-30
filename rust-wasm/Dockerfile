FROM rust:1.73.0

WORKDIR /app

RUN rustup target add wasm32-unknown-unknown

CMD cargo build --release --target wasm32-unknown-unknown
