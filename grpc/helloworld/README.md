# Hello World

## Generate proto and stub

```bash
brew install prototool

cd proto
prototool config init
prototool create helloworld/helloworld.proto
prototool generate
prototool lint
```

## Execute

```bash
go run server/main.go
go run client/main.go
```

## Run server on Docker

```bash
docker build . -t helloworld
docker run -it -p 50051:50051 helloworld
```

## Test by gRPC cli

[gRPC cli](https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md#grpc-cli)

```bash
./grpc_cli ls localhost:50051
```