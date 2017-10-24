package main

import (
	"net/http"
	"log"
	"os"
)

func ServeFile(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, os.Args[2])
}

func main() {
	http.HandleFunc("/", ServeFile)
	log.Fatal(http.ListenAndServe(":" + os.Args[1], nil))
}
