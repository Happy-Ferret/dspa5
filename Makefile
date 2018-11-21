MAKEFLAGS += -j 3
all: dist

# note -ldflags="-s -w" is used to strip debugging tables from executables to reduce file size

ifeq ($(GOOS),windows)
    EXT := .exe
endif

grpc:
	protoc -I/usr/local/include -I. \
	-I${GOPATH}/src \
	--go_out=plugins=grpc:. \
	dspa5/*.proto

assets:
	go generate

client:
	go build -ldflags="-s -w" dspa-client.go

speaker:
	go build -ldflags="-s -w" dspa-speaker.go

broadcaster:
	go build -ldflags="-s -w" dspa-broadcaster.go

dist: client broadcaster speaker
	mkdir -p dist && \
		mv dspa-client${EXT} dist/
		mv dspa-broadcaster${EXT} dist/
		mv dspa-speaker${EXT} dist/
	upx -q dist/dspa-client${EXT}
	upx -q dist/dspa-speaker${EXT}
	upx -q dist/dspa-broadcaster${EXT}

update:
# dep init created vendor directory!
	dep ensure -update

install:
	go install ./...

fmt:
	go fmt ./...
