all: dist

grpc:
	protoc -I/usr/local/include -I. \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:. \
	dspa5/*.proto

assets:
	cd dspa-server && go generate

server: grpc assets
	cd dspa-server && go build

client: grpc assets
	cd dspa-client && go build

pack: client server
	upx dspa-server/dspa-server
	upx dspa-client/dspa-client

dist: pack
	mkdir -p dist && \
		mv dspa-server/dspa-server dist/
		mv dspa-client/dspa-client dist/
