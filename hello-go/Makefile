DOCKER_IMAGE_TAG := 0.0.1

clean:
	@rm -f hello-go
	@rm -f hello-go-static

init:
	@go get github.com/stretchr/testify

build: clean
	@go build

buildstatic:
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-go-static .

test: build
	@go test

docker: buildstatic
	@docker build -t hello-go:$(DOCKER_IMAGE_TAG) .

kube: buildstatic
	@eval $$(minikube docker-env); \
	env | grep DOCKER; \
	docker build -t hello-go:$(DOCKER_IMAGE_TAG) .
