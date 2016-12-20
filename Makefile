

all:
	go build -v ./pkg/...
	mkdir -p ./bin
	go build -o bin/ttbot ./cmd/ttbot 

