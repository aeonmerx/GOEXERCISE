package main

import (
	"fmt"
	"log"
	"net/http"
)

func holaFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola nena como estás!")
}

func main() {
	http.HandleFunc("/", holaFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
