FROM golang:1.17 AS build
WORKDIR /http-server/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux GOARCH=amd64 go build -installsuffix cgo -o http-server main.go

FROM alpine:latest
COPY --from=build /http-server/http-server /http-server
ENV VERSION v0.1
EXPOSE 80
WORKDIR /
ENTRYPOINT [ "/http-server" ]