package main

import (
	"fmt"
	"net/http"
)

type HelloServer struct {
	greeting string
}

func (serv *HelloServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	name := req.URL.Path[len("/hello/"):]

	switch req.Method {
		case http.MethodGet:
			serv.greet(resp, name)
	}
}

func (serv *HelloServer) greet(resp http.ResponseWriter, name string) {
	resp.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Hello, %s", name)
	resp.Write([]byte(msg))
}
