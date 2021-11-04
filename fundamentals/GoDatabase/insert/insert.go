package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:MySql2019!@/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("insert into users(name) values(?)")
	stmt.Exec("Wene")
	stmt.Exec("Ismael")

	res, _ := stmt.Exec("Lucas")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	lines, _ := res.RowsAffected()
	fmt.Println(lines)
}
