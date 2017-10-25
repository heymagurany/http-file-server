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
		fmt.Printf("usage: http-file-server <port> <file name>\n")
		os.Exit(1)
	}
	port, fileName := args[0], args[1]
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, fileName)
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
