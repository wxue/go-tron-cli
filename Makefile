default: build

pre:
	export PATH=$PATH:$GOPATH/bin

fmt:
	go fmt ./...

run:
	GO111MODULE=on go run ./troncli/main.go

install:
	GO111MODULE=on go install ./troncli

bin:
	mkdir -p bin

build: bin
	GO111MODULE=on go build -o ./bin/troncli ./troncli/main.go

clean:
	rm -rf bin