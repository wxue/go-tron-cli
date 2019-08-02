default: build

pre:
	export PATH=$PATH:$GOPATH/bin

fmt:
	go fmt ./...

run:
	GO111MODULE=on go run ./main.go ./config.go

install:
	GO111MODULE=on go install ./

bin:
	mkdir -p bin

build: bin
	GO111MODULE=on go build -o ./bin/troncli ./main.go  ./config.go

clean:
	rm -rf bin