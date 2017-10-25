package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, os.Args[2])
	})
	log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))
}
