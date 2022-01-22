package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dns := "root:root@123@/cursogo"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(1, "Thiago") // chave duplicada no banco

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
