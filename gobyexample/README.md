# Go by Example

- [Go by Example](https://gobyexample.com)

## Run

```bash
go run 01-hello-world.go
go build 01-hello-world.go
./01-hello-world
```

## Rerun Go program when file modified

```bash
go install github.com/hanksudo/gorerun@latest
gorerun -f 01-hello-world.go
```

## Make filename's number up to date

```bash
# 1. update list.txt

# 2. run script
go run . rename-files.go
```