# gRPC

## Installation

```bash
# Install protobuf
brew install protobuf

# Install gRPC
go get -u google.golang.org/grpc

# Install protoc plugin for Go
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Generate proto and stub

```bash
cd proto
prototool create helloworld/helloworld.proto
prototool generate
prototool lint
```