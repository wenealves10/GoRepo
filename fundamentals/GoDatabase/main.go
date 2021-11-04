package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:MySql2019!@/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, _ := db.Query("select id, name from users where id > ?", 1)
	defer rows.Close()

	for rows.Next() {
		var u user

		rows.Scan(&u.Id, &u.Name)

		fmt.Println(u)
	}
}
