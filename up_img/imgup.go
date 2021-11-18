package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit") // Parse our multipart form, 10 << 20 specifies a maximum
	r.ParseMultipartForm(10 << 20)          // upload of 10 MB files.
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	}
	file, handler, err := r.FormFile("myFile") // FormFile returns the first file for the given key `myFile`
	if err != nil {
		fmt.Println("Error Retrieving the File") // it also returns the FileHeader so we can get the Filename,
		fmt.Println(err)
		return
	}

	defer file.Close() // the Header and the size of the file
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png") // Create a temporary file within our temp-images directory that follows
	if err != nil {
		fmt.Println(err) // a particular naming pattern
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8084", nil)
}

func main() {
	fmt.Println("Hello ........")
	setupRoutes()
}
