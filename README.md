# Load Balanced 'Hello World' App in MiniKube

## Purpose
The code in this repo represents my teaching myself how to drive Kubernetes programmatically, specifically how to drive it with the [official Go client library](https://github.com/kubernetes/client-go).

It's possible to script K8s using KubeCtl, but I wanted to look at management of a K8s
system from a developer's perspective, paying attention to good dev practices such as
automated testing. Looking at the available Go clients, I was struck by the seeming richness of the Go
client, in comparison to the Python or Ruby clients for example. There were two stand-outs:

- full object model of the K8s internals, as opposed to a method that took yaml as input
- the "fake" subpackages, that allow a comprehensive mock library for unit testing (suggesting that the library authors are thinking about good test coverage from the outset)   

I was very much poking about here. Poking about the Go client to see how it worked, and poking about the fakes package to see how it worked. If you dig into the history, you'll see
that I didn't develop it TDD style, as I was figuring out lots of the basics that would allow me to do so in future. Specifically, you'll see
that between commits de1f279 and 9e8c1b6, I refactored the production code to make it testable.
 
## Lessons Learned

I did this exercise to try to answer questions abut how the Go client codebase worked, and also about the merits of controlling a K8s system
programmatically versus more traditional "infrastructure as code" options.  

### Coding with the Go Client

There are two big take-homes for me from this exercise

- use an interface to supply any of the CRUD subsystems in the client library, and write your production code against this interface (yes, this is pretty much standard advice for using any test mocks)
- access fakes via the top-level ClientSet, which is the only fake you will need to explicitly create. All other fakes can be retrieved through that,
once you've got hold of it. See `kube-deploy/k8s_system_test.go` for examples - I used the deployments, services and ingresses CRUD interfaces, others will work similarly. The good news
is that it keeps the interface quite simple. Once you've got the clientset, just write your production code
as you would anyway.

### Merits of this approach

I'm on the fence, to be honest. 

On the plus side, this gives me all the testability that I want for my code,
and does so nice and low down the test pyramid. (i.e. I'm not having to deploy a real minikube or anything expensive 
in order to test the business logic). As a deployment scales up, this could save me big time
on the test suite execution speed, which would be a huge win.  

On the minus side, I'm exporting something more complex and alien to my customers, who now need to know
how to run a command-line executable. If they're a traditional operations team, they may be more in their
comfort zone to be receiving yaml from me. The alternative, would be to write, and unit test, something that
generates yaml from a more structured model, which I can pass on to the customer. Much depends here on non-technical
issues, such as whether the customer is internal or external, and whether we see ourselves as on a journey together 
or just getting stuff done (i.e. are they open to considering improvements in their processes).


Anyway, if you stumble across this, hope it proves useful to you in exploring Go / Kubernetes integration. I wish I'd been able
to read this when I started out looking at this stuff.
  

## Setup
This project contains two sub-projects (a trendy mini-`monorepo` ?!), each with an extensive README file covering setup, running etc.:

* [hello-go/README.md] A simple REST API app, written in Go, that exposes a "hello world" echo style API. The associated Makefile allows for easy creation of a docker image, and exposng that docker image to the minikube repo

* [kube-deploy/README.md] A second Go application (command-line, so suitable for CI) that administers the minikube, allowing setup, live update and teardown of a load-balanced service on a predictable port number, exposing the REST API to the outside world

