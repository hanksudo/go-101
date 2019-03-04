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