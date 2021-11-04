package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:MySql2019!@/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("update users set name = ? where id = ?")
	stmt.Exec("Wene Alves", 1)
	stmt.Exec("Ismael Albert", 2)

	stmt2, _ := db.Prepare("delete from users where id = ?")
	stmt2.Exec(3)
	stmt2.Exec(10)
	stmt2.Exec(11)
	stmt2.Exec(20)
}
