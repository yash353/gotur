package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"image/png"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kbinani/screenshot"
)

var db, _ = sql.Open("mysql", "root:@tcp(localhost:3306)/user") //db_conn

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello yash")
}
func screens() {
	t := time.Now()
	formatted := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fmt.Println(formatted) //time func.
	// fmt.Println("Location : ", t.Location(), " Time : ", t) // local time

	conn, _ := net.Dial("udp", "8.8.8.8:80") //local ip address of machine

	defer conn.Close()

	ipAddress := conn.LocalAddr()

	fmt.Println("Machine ip address :", ipAddress) //print ip of local machine

	n := screenshot.NumActiveDisplays() //screenshot
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/user") //db_conn

	// if err != nil {
	// 	log.Fatal(err)
	// }

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}

		fileName := fmt.Sprint("imgfolder/" + formatted + ".png")
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
		// ins = append(ins, "(?,?)")
		st := fmt.Sprintf("insert into test1(imgname , ip) values('%s','%s')", fileName, ipAddress) //insert into db
		fmt.Println(st)
		insert, err := db.Query(st)

		if err != nil {
			log.Fatal(err)
		} else {

			fmt.Println("successfully insert")
		}
		defer insert.Close()
	}

}
func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		fmt.Println("getmetthh")
		t.ExecuteTemplate(w, "index.html", nil)
	} else {
		r.ParseForm()

		fmt.Println("username:", r.Form["usnam"])
		fmt.Println("password:", r.Form["pass"])

		fmt.Println("sucessfully insert user")
		for {
			time.Sleep(15 * time.Second)
			screens()
		}

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
