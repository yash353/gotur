package main

import (
	"database/sql"
	"fmt"
	"log"

	//	"github.com/go-sql-driver/mysql"
	//"github.com/go-sql-driver/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type fdb struct {
	sno      int    `json:"sno"`
	name     string `json:"name"`
	mobileno int    `json:"mobileno"`
	email    string `json:"email"`
	password string `json:"password"`
}

func main() {

	fmt.Println("data-base connection")

	// cfg := sql.Config{
	// 	User: os.Getenv("root"),

	// 	DBName: "user",
	// }
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/user")

	if err != nil {

		fmt.Println("connection  error")
	}

	fmt.Println("connected")

	defer db.Close()

	row, err := db.Query("SELECT * from test")
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		var fdb fdb
		err := row.Scan(&fdb.sno, &fdb.name, &fdb.mobileno, &fdb.email, &fdb.password)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(fdb.sno)
		fmt.Println(fdb.name)
		fmt.Println(fdb.mobileno)
		fmt.Println(fdb.email)
		fmt.Println(fdb.password)

	}

}
