# kube-deploy.go

The code in this folder implements a go program that uses the kubernetes API to deploy a pod of the docker images
created by the `../hello-go` program.

The code has been developed test first, with unit and integration tests. 

## Setup Go Environment

`go get github.com/stretchr/testify`

`go get k8s.io/client-go@v11.0.0`

`go get k8s.io/api@kubernetes-1.14.0`

`go get k8s.io/apimachinery@kubernetes-1.14.0`

## Build binary

`go build`

## Run in MiniKube  (Create Deployment)

`minikube addons enable ingress`
`minikube start`


`kubectl get pod` 

*Output*
`No resources found`

Now make sure the hello-go docker image is made available to the minikube registry (see [../hello-go/README])

`./hello-kube-deploy`
`kubectl get pod`

*Output*
```
NAME                                READY     STATUS    RESTARTS   AGE
hello-deployment-57648cd8c4-46xtx   1/1       Running   1          2h
hello-deployment-57648cd8c4-hpqqd   1/1       Running   1          2h
```

`kubectl get services`

*Output*
```
NAME            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
hello-service   NodePort    10.107.54.138   <none>        8090:30280/TCP   2h
kubernetes      ClusterIP   10.96.0.1       <none>        443/TCP          2d
```

### Test Manually via service URL directly

`curl -X GET $(minikube service demo-deployment --url)/hello/kubert`

*Output*
`Hello, kubert`

### Manually Delete Deployment (but leave minikube running)

`kubectl delete deployment hello-deployment`
`kubectl delee service hello-service`
