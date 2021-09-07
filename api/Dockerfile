FROM golang:alpine AS build

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

EXPOSE 8000

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go build ." --command=./cinemas