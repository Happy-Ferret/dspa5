MAKEFLAGS += -j 3
all: dist

# note -ldflags="-s -w" is used to strip debugging tables from executables to reduce file size

grpc:
	protoc -I/usr/local/include -I. \
	-I${GOPATH}/src \
	--go_out=plugins=grpc:. \
	dspa5/*.proto

assets:
	cd dspa-speaker && go generate

client:
	cd dspa-client && go build -ldflags="-s -w"
	upx -q  dspa-client/dspa-client

speaker:
	cd dspa-speaker && go build -ldflags="-s -w"
	upx -q dspa-speaker/dspa-speaker

broadcaster:
	cd dspa-broadcaster && go build -ldflags="-s -w"
	upx -q dspa-broadcaster/dspa-broadcaster

dist: client broadcaster speaker
	mkdir -p dist && \
		mv dspa-speaker/dspa-speaker dist/
		mv dspa-client/dspa-client dist/

update:
# dep init created vendor directory!
	dep ensure -update
