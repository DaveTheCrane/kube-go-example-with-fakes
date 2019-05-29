package main

import (
	"fmt"
	"net/http"
	"strings"
)

type HelloServer struct {
	greeting string
}

func (serv *HelloServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	resp.Write([]byte(fmt.Sprintf("[[[request :: %s :: ]]]\n", req.URL.String())))
	var name string
	if (strings.HasPrefix(path, "/hello")) {
		if (len(path) > len("/hello/")) {
			name = path[len("/hello/"):]
		} else {
			name = req.URL.String()
		}
		switch req.Method {
		case http.MethodGet:
			serv.greet(resp, name)
		}
	} else {
		serv.notFound(resp)
	}
}

func (serv *HelloServer) greet(resp http.ResponseWriter, name string) {
	resp.WriteHeader(http.StatusOK)
	var msg string
	if (name != "") {
		msg = fmt.Sprintf("Hello, %s\n", name)
	} else {
		msg = "Hello\n"
	}
	
	fmt.Print(fmt.Sprintf("serving greeting - %s\n", msg))
	resp.Write([]byte(msg))
}

func (serv *HelloServer) notFound(resp http.ResponseWriter) {
	resp.WriteHeader(http.StatusNotFound)
	resp.Write([]byte("Not Found\n"))
}