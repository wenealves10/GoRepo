package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:MySql2019!@/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	transtion, _ := db.Begin()

	stmt, _ := transtion.Prepare("insert into users(id, name) values(?,?)")
	stmt.Exec(100, "Nadson")
	stmt.Exec(200, "Frank")

	_, err = stmt.Exec(110, "Amilson")
	if err != nil {
		transtion.Rollback()
		log.Fatal(err)
	}

	transtion.Commit()

}
