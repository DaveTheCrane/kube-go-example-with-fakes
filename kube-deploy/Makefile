DOCKER_IMAGE_TAG := 0.0.1

clean:
	@rm -f hello-kube-deploy

init:
	@go get github.com/stretchr/testify
	@go get k8s.io/client-go@v11.0.0
	@go get k8s.io/api@kubernetes-1.14.0
	@go get k8s.io/apimachinery@kubernetes-1.14.0

build: clean
	@go build

test: build
	@go test

