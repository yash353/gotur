package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type input struct {
	name     string
	mobileno int
	email    string
	password string
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/user")

	if err != nil {
		log.Fatal(err)
	}

	var ins []string
	ins = append(ins, "(?,?)")
	st := fmt.Sprint("insert into test(name, mobileno, email,password) values('shree',1213535342,'shree@gmail.com','12345')")
	insert, err := db.Query(st)

	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Println("successfully insert")
	}

	defer db.Close()
	defer insert.Close()

}
