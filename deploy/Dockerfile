# Golang Server's Dockerfile
FROM golang:1.14.6-alpine

## installing git
RUN apk add --no-cache git

# move to workdir
WORKDIR /codementordev

# copy current source to container working dir
ADD ../ /codementordev

# build golang executable
RUN cd /codementordev/cmd/codementordev-hello-johnhckuo-server && go build -o server

# expost port for api server
EXPOSE 8080

# excute the executable we just built
ENTRYPOINT ./cmd/codementordev-hello-johnhckuo-server/server