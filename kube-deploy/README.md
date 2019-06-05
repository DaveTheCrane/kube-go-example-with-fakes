# kube-deploy.go

The code in this folder implements a go program that uses the kubernetes API to deploy a pod of the docker images
created by the `../hello-go` program. The following artefacts are created:

* a Deployment named hello-deployment, with n replicas
* a Service named hello-service, exposing the deployment within the cluster
* an Ingress named hello-lb, which does a little url-rewriting, and load balancing using the default nginx load balancer

The program can deploy this setup, remove it, and update the replica count or url rewriting of an active setup. The "hello-" 
prefix for internal names of items can be overridden. It also provides a rudimentary listing
mechanism, which is probably less useful than using `kubectl`. 

The code has been developed in an ad hoc fashion, with some tests added as I go along. 

## Setup Go Environment

`make init`

or

```
go get github.com/stretchr/testify
go get k8s.io/client-go@v11.0.0
go get k8s.io/api@kubernetes-1.14.0
go get k8s.io/apimachinery@kubernetes-1.14.0
```

## Build binary

`make build` or `go build`

## Run tests

`make test` or `go test`

## Run in MiniKube 

### Setup

`minikube addons enable ingress`
`minikube start`

Check that we have an empty initial setup...

`kubectl get pod` 

*Output*
`No resources found`

Find out the minikube IP address using `minikube ip`, and add it to your hosts file, e.g.

```
192.168.99.101 mini.kube.io
```

(The app defaults to a hostname of mini.kube.io, but this can be overridden by command-line flags).

Now make sure the hello-go docker image is made available to the minikube registry (see [../hello-go/README])

The binary `hello-kube-deploy` app can be used to create, delete and modify deployments

### Create a deployment

To create a deployment with default settings:

`./hello-kube-deploy -create`

Now check that everything's come up correctly:

`kubectl get deployments`

*Output*
```
NAME               DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
hello-deployment   5         5         5            5           11s
```

`kubectl get pods`

*Output*
```
NAME                                READY     STATUS    RESTARTS   AGE
hello-deployment-57648cd8c4-5dmcn   1/1       Running   0          11s
hello-deployment-57648cd8c4-78qgp   1/1       Running   0          11s
hello-deployment-57648cd8c4-jmnct   1/1       Running   0          11s
hello-deployment-57648cd8c4-lbvb7   1/1       Running   0          11s
hello-deployment-57648cd8c4-lz5bt   1/1       Running   0          11s
```

`kubectl get services`

*Output*
```
NAME            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
hello-service   ClusterIP   10.111.251.39    <none>        8090/TCP   1m
kubernetes      ClusterIP   10.96.0.1        <none>        443/TCP    10d
```

`kubectl get ingress`

*Output*
```
NAME       HOSTS          ADDRESS     PORTS     AGE
hello-lb   mini.kube.io   10.0.2.15   80        1m
```

The url rewriting will map `/hello/**` to `/hi/**` by default, so the load-balanced service can be reached by:

`curl -X GET http://mini.kube.io/hi/dave`
 
*Output*
`Hello, dave`

(Note that the ingress paths are hostname-specific. A curl to the raw IP address will hit the default nginx backend and return a 404)

### Updating the deployment

A running deployment can be updated in two ways: 

#### Change the number of replicas

`./hello-kube-deploy -update -replicas=3`

`kubectl get deployments`

*Output*
```
NAME               DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
hello-deployment   3         3         3            3           12m
```

#### Change the url-rewrite path

Note that this replaces the existing rewrite rule. (The Ingress model supports multiple
rewrite rules, but my program doesn't at the moment.)

`./hello-kube-deploy -update -ext-path=howdy`

`curl -X GET http://mini.kube.io/hi/dave`

*Output*
`default backend - 404`

`curl -X GET http://mini.kube.io/howdy/dave`

*Output*
`Hello, dave`

### Delete Deployment

The deployment can be removed entirely by 

`hello-kube-deploy -delete`

### Command line flags

A number of parameters can be changed via the command line. Run `./hello-kube-deploy -help` to get full usage text:

```
Usage of ./hello-kube-deploy:
  -create
        Create deployment, service and ingress in kubernetes cluster
  -delete
        Delete current deployment, service and ingress in kubernetes cluster
  -ext-path string
        path used by load balancer rewrite rule to access the service (default "hi")
  -host-name string
        hostname via which service is exposed in the load balancer (default "mini.kube.io")
  -image-name string
        docker image name for service to be deployed (default "hello-go")
  -image-tag string
        docker image tag for service to be deployed (default "0.0.1")
  -list
        List current deployment, service and ingress in kubernetes cluster
  -prefix string
        prefix used to identify artefacts in the kubernetes system (default "hello")
  -replicas int
        number of load-balanced replicas to deploy (default 5)
  -update
        Update selected features of current deployment, currently only external path and replica count supported
```

Note that it's possible to deploy multiple setups in the same kube side by side by supplying different values to `prefix`