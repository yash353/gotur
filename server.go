package main

import (
	"fmt"
	"log"
	"net/http"
)

func hompag(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "home page")

}
func handlerequest() {
	http.HandleFunc("/", hompag)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handlerequest()
}
