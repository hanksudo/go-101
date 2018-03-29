# go-101

My repository to learn Go Language <http://golang.org>

## Install Go (on OSX)

```bash
brew install go
```

## Basic Commands

```bash
# download from remote (download into both ./src and ./bin)
go get github.com/hanksudo/go-101/tour/hello

# run program
go run tour/hello/hello.go
go run tour/functions/functions.go

# build program
cd hello
go build hello.go
./hello

# go install to $GOPATH/pkg
go install newmath

# go install to $GOPATH/bin
go install mathapp

# test code
go test github.com/hanksudo/go-101/tour/stringutil

# format code syntax
go fmt -w

# documentation
godoc builtin
godoc net/http
godoc fmt Printf
godoc -http=:8080

# others
go version
go env
go list
go list all
go list -json all
```

## Coding style

    variableName

## Run go program when file modified

```bash
go get github.com/hanksudo/gowatch
gowatch -f main.go
```

## Reference

- [Tour of Go](http://tour.golang.org/)
- [Tour of Go - exercise solutions](https://github.com/golang/tour/tree/master/solutions)
- [How to Write Go Code - The Go Programming Language](http://golang.org/doc/code.html)
- [Effective Go](http://golang.org/doc/effective_go.html)
- [Writing Web Applications - The Go Programming Language](http://golang.org/doc/articles/wiki/)
- [Share Memory By Communicating - The Go Programming Language](http://golang.org/doc/codewalk/sharemem/)
- [First-Class Functions in Go - The Go Programming Language](http://golang.org/doc/codewalk/functions/)
- [逆引きGolang](http://ashitani.jp/golangtips/index.html)
- [GitHub - fatih/vim-go-tutorial: Tutorial for vim-go](https://github.com/fatih/vim-go-tutorial)
