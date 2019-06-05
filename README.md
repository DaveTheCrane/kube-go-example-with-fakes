# Load Balanced 'Hello World' App in MiniKube

This project contains two sub-projects, each with an extensive README file:

* [hello-go/README.md] A simple REST API app, written in Go, that exposes a "hello world" echo style API. The associated Makefile allows for easy creation of a docker image, and exposng that docker image to the minikube repo

* [kube-deploy/README.md] A second Go application (command-line, so suitable for CI) that administers the minikube, allowing setup, live update and teardown of a load-balanced service on a predictable port number, exposing the REST API to the outside world
