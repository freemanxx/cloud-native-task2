run:
	echo "runing http server"
	go run main.go
release:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

.PHONY: release