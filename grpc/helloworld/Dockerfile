FROM golang:1.12

RUN go get google.golang.org/grpc

COPY . /go/src/app
WORKDIR /go/src/app

CMD ["go", "run", "server/main.go"]