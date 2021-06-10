# go-101

Learn Go Programming Language <https://golang.org>

## Install Go (on Mac)

```bash
brew install go
```

## Folder structures

- tour - [Tour of Go](http://tour.golang.org/)
- effective-go - [Effective Go](http://golang.org/doc/effective_go.html)
- gobyexample - [Go by Example](https://gobyexample.com/)
- gowiki - [Writing Web Applications - The Go Programming Language](https://golang.org/doc/articles/wiki/)
- learn-go-with-tests - [learn go with tests](https://quii.gitbook.io/learn-go-with-tests/)
- [play-with-go.dev](https://play-with-go.dev/)

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

## Rerun Go program when file modified

```bash
go get -u github.com/hanksudo/gowatch
gowatch -f main.go
```

## References

- [Tour of Go - exercise solutions](https://github.com/golang/tour/tree/master/solutions)
- [GitHub - enocom/gopher-reading-list: A curated selection of blog posts on Go](https://github.com/enocom/gopher-reading-list)
- [How to Write Go Code - The Go Programming Language](http://golang.org/doc/code.html)
- [Writing Web Applications - The Go Programming Language](http://golang.org/doc/articles/wiki/)
- [preface · Build web application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/preface.html)
- [Share Memory By Communicating - The Go Programming Language](http://golang.org/doc/codewalk/sharemem/)
- [First-Class Functions in Go - The Go Programming Language](http://golang.org/doc/codewalk/functions/)
- [逆引きGolang](http://ashitani.jp/golangtips/index.html)
- [GitHub - fatih/vim-go-tutorial: Tutorial for vim-go](https://github.com/fatih/vim-go-tutorial)
- [JustForFunc: Programming in Go - YouTube](https://www.youtube.com/channel/UC_BzFbxG2za3bp5NRRRXJSw)
- [Frontmatter | Go Bootcamp |  Softcover.io](http://www.golangbootcamp.com/book/)
- [GopherAcademy - Conferences, Training, and Community](https://www.gopheracademy.com/)
- [Concurrency, Goroutines and GOMAXPROCS Go, (Golang) Programming - Blog - Ardan Labs](https://www.ardanlabs.com/blog/2014/01/concurrency-goroutines-and-gomaxprocs.html)