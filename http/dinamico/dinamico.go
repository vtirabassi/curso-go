package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	//01 mes
	//02 dia
	//03 hora
	//04 minuto
	//05 segundo
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta) //funcao informada no HandleFunc tem que ter ja um padrao de assinatura
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
