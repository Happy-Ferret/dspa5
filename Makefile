all: grpc_go

grpc_go:
	protoc -I/usr/local/include -I. \
	-I${GOPATH}/src \
	-Iproto \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:. \
	proto/*.proto

