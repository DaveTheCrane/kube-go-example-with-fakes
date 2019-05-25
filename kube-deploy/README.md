# kube-deploy.go

The code in this folder implements a go program that uses the kubernetes API to deploy a pod of the docker images
created by the `../hello-go` program.

The code has been developed test first, with unit and integration tests. 

## Setup

`go get github.com/stretchr/testify`
`go get k8s.io/client-go@v11.0.0`
`go get k8s.io/api@kubernetes-1.14.0`
`go get k8s.io/apimachinery@kubernetes-1.14.0`

## Build binary

`go build`

## Run in MiniKube  (Create Deployment)

`minikube start`
`kubectl get pod` 

`No resources found`

Make sure the hello-go docker image is made available to the minikube registry (see [../hello-go/README])

`./hello-kube-deploy`
`kubectl get pod`

```
NAME                              READY     STATUS         RESTARTS   AGE
demo-deployment-7c9964f64-lw8kn   0/1       ErrImagePull   0          8s
demo-deployment-7c9964f64-skrnq   0/1       ErrImagePull   0          8s
```

`kubectl expose deployment demo-deployment --type=NodePort`

`curl -X GET http://192.168.99.101:30041/hello/maple`

### Delete Deployment (but leave minikube running)

`kubectl delete deployment demo-deployment`
`curl -X GET $(minikube service demo-deployment --url)/hello/kubert` -> `Hello, kubert`