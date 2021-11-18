package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, " hello.....  yash")
}

func wsocket(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, " hii..... yash")
}

func setupRoutes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/web", wsocket)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	setupRoutes()
}
