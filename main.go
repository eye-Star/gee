package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	_, err := w.Write([]byte("Hello"))
	if err != nil {
		panic(err)
	}
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	_, err := w.Write([]byte("World"))
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	log.Fatal(http.ListenAndServe(":8080", nil))
}