
run:
	echo "runing http server"
	go run main.go
release:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/http-server .
docker:
	echo "building docker image"
	docker buildx build --platform linux/amd64 . -t freeman007/cloud-native-http-server:v0.1

docker-push: docker
	echo "pushing docker image to docker.io"
	docker push freeman007/cloud-native-http-server:v0.1

.PHONY: release