package main

import (
	"fmt"
	"log"
	"net/http"
)

type H struct{}

func (*H) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello world!"))
}

func main() {
	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)
	h := &H{}
	http.Handle("/hello", h)

	// Start the server.
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)

	defer func() {

	}()
}
