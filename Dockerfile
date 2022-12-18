FROM alpine:latest
ADD bin/amd64/http-server /http-server

ENV VERSION v0.1
EXPOSE 80
WORKDIR /
ENTRYPOINT [ "/http-server" ]