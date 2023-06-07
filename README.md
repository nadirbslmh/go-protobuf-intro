# go-protobuf-intro

Simple application using Protocol Buffers in Go.

## How to use

1. Clone this repository.

2. Run the application.

```sh
go run main.go
```

3. If protocol buffer file is modified, generate the Go code.

```sh
protoc -I=protobuf --go_out=./bookpb/ protobuf/book.proto
```
