package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is test route")
	})

	http.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is test1 route")
	})

	log.Fatal(http.ListenAndServe(":7777", nil))
}