package main

import (
	"database/sql"
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kbinani/screenshot"
)

// var fileName string

func main() {
	n := screenshot.NumActiveDisplays()
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/user")

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		t := time.Now()
		formatted := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		// fmt.Println(formatted)
		fileName := fmt.Sprint(formatted + ".png")
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
		// ins = append(ins, "(?,?)")
		st := fmt.Sprintf("insert into test0(imgname) values('%s')", fileName)
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
