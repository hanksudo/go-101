.PHONY: proto

install:
	@go get github.com/golang/protobuf/{proto,protoc-gen-go}
	@go get google.golang.org/grpc
	# protoc --go_out=plugins=grpc:. *.proto
	
proto:
	@echo Compiling proto files...
	# @find **/*.proto -exec protoc --go_out=plugins=grpc:. {} \;
	cd proto && prototool generate
	@echo Done!
