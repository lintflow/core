
buff_build:
	protoc --go_out=plugins=grpc:. proto/*.proto

build:
	go build ./...
