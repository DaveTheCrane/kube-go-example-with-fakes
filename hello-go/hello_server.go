package main

import (
	"fmt"
	"net/http"
)

type HelloServer struct {
	greeting string
}

func (serv *HelloServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	var name string
	if (len(path) > len("/hello/")) {
		name = path[len("/hello/"):]
	} else {
		name = ""
	}

	switch req.Method {
		case http.MethodGet:
			serv.greet(resp, name)
	}
}

func (serv *HelloServer) greet(resp http.ResponseWriter, name string) {
	resp.WriteHeader(http.StatusOK)
	var msg string
	if (name != "") {
		msg = fmt.Sprintf("Hello, %s", name)
	} else {
		msg = "Hello"
	}
	
	fmt.Print(fmt.Sprintf("serving greeting - %s\n", msg))
	resp.Write([]byte(msg))
}
