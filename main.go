package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":3000"
)

var calls = 0

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	calls++
	fmt.Fprintf(w, "Hello, world! This server is served by a container that was built using another container. Everything is contained\n", calls)
}

func init() {
	fmt.Printf("Started server at http://localhost%v.\n", port)
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(port, nil)
}

func main() {}
