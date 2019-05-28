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

You'll see a line appear in the output of the server terminal:

`serving greeting - Hello, Fred`

## Build Docker Image

Rather than build inside the official docker image, we're going to build the docker image from a minimal Scratch, and host
a statically linked binary that we've created on the host:

### Build the statically-linked binary
`CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-go-static .`

### Build the docker image
`docker build -t hello-go:0.0.1 .`

NB: We use a tagged image, not `latest`, so that MiniKube can pick it up without trying to fetch it from a central repo (latest defaults to ImagePullPolicy=Always)

### Run the docker image
`docker run -p 8090:8090 hello-go`

Verify using the curl command as per running locally in the previous section.

## Build Docker image for use with MiniKube

MiniKube can't see the local docker registry. So if you want to use the image in minikube without exporting to a public
image registry, you need to set environment variables up to point the local docker image build to minikube

### Start MiniKube

`minikube start`

### Export MiniKube environment to current shell

`eval $(minikube docker-env)`

### Build the Docker Image as before

`docker build .`