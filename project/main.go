package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// fmt.Println(r.Form)
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello yash")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		fmt.Println("username:", r.Form["usnam"])
		fmt.Println("password:", r.Form["pass"])
		fmt.Println("img", r.Form["img"])
		// t := time.Now()                                         //time func.
		// fmt.Println("Location : ", t.Location(), " Time : ", t) // local time

		// conn, _ := net.Dial("udp", "8.8.8.8:80") //local ip address of machine

		// defer conn.Close()

		// ipAddress := conn.LocalAddr().(*net.UDPAddr)

		// fmt.Println("Machine ip address :", ipAddress) //print ip of local machine

	}
}

func main() {

	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
