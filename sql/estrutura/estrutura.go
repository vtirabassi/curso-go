package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err) //para o programa
	}
	return result
}

func main() {
	dns := "root:root@123@/"
	db, err := sql.Open("mysql", dns) //string mysql usa o drive my sql, passo user e senha
	if err != nil {
		panic(err)
	}
	defer db.Close() //no final da funcao ela sera finalizada, fechada

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
		)`)
}
