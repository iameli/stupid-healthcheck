package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello. I am up.\n")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	listen := fmt.Sprintf(":%s", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(listen, nil)
}
