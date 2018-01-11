package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Printf("usage: http-file-server <port> <file name> [url path]\n")
		os.Exit(1)
	}
	port, fileName := args[0], args[1]
	var path string
	if len(args) > 2 {
		path = args[3]
	} else {
		path = "/"
	}
	http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("%s %s\n", req.Method, req.URL.String())
		http.ServeFile(w, req, fileName)
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
