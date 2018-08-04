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
	cd dspa-speaker && go generate

client:
	cd dspa-client && go build -ldflags="-s -w"
	upx -q  dspa-client/dspa-client${EXT}

speaker:
	cd dspa-speaker && go build -ldflags="-s -w"
	upx -q dspa-speaker/dspa-speaker${EXT}

broadcaster:
	cd dspa-broadcaster && go build -ldflags="-s -w"
	upx -q dspa-broadcaster/dspa-broadcaster${EXT}

dist: client broadcaster speaker
	mkdir -p dist && \
		mv dspa-client/dspa-client${EXT} dist/
		mv dspa-broadcaster/dspa-broadcaster${EXT} dist/
		mv dspa-speaker/dspa-speaker${EXT} dist/

update:
# dep init created vendor directory!
	dep ensure -update

install:
	go install ./...

fmt:
	go fmt ./...
