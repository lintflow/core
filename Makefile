
buff_build:
	protoc --go_out=plugins=grpc:. proto/*.proto

install:
	go get -u -v ./...
	go install ./cmd/...

build:
	go install ./cmd/...
