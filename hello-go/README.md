# hello.go

The code in this folder implements a simple "hello world" HTTP server, using golang. The server simply parrots back a greeting, using the full URL path as a name.

The code has been developed test first, with unit and integration tests. 

## Setup

`go get github.com/stretchr/testify`

## Run Tests

`go test`

## Run Server Locally (and verify it manually)

`go build`
`./hello-go` 

This will block the shell, until killed with a `Ctrl-C`.

In a separate terminal, type

`curl -X GET http://localhost:8090/hello/Fred`

## Build Docker Image

