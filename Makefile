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
	# https://github.com/jteeuwen/go-bindata
	go-bindata -pkg speaker -o bindata/speaker/speaker.go chimes/
	go-bindata -pkg display -o bindata/display/display.go etc/roboto/

client:
	go build -ldflags="-s -w" dspa-client.go
	upx -q dspa-client${EXT}

display:
	go build -ldflags="-s -w" dspa-display.go
	upx -q dspa-display${EXT}

speaker:
	go build -ldflags="-s -w" dspa-speaker.go
	upx -q dspa-speaker${EXT}

broadcaster:
	go build -ldflags="-s -w" dspa-broadcaster.go
	upx -q dspa-broadcaster${EXT}

dist: client broadcaster speaker display
	mkdir -p dist && \
		mv dspa-client${EXT} dist/
		mv dspa-broadcaster${EXT} dist/
		mv dspa-speaker${EXT} dist/
		mv dspa-display${EXT} dist/

update:
# dep init created vendor directory!
	dep ensure -update

install:
	go install ./...

fmt:
	go fmt ./...
