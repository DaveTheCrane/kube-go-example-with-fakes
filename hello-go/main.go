package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &HelloServer{}
	port := 8090

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), server); err != nil {
		log.Fatalf("could not listen on port %d %v", port, err)
	}
}