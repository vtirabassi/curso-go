package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	dns := "root:root@123@/cursogo"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id > ?", 1)
	defer rows.Close() //fecha o result set

	for rows.Next() { //go nao existe while, use o for pra fazer um laço indeterminado
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}
