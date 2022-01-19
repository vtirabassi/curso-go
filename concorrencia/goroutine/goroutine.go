package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// sem concorrência (ele executa primeiro o da Maria e depois que acabar executa do Joao)
	// fale("Maria", "Por que você não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você!", 1)

	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)

	//Ele cria uma goroutine para executar, é como se fosse um "thread", sao funcoes executadas de forma independete
	go fale("Maria", "Entendi!", 10)
	fale("João", "Parabéns", 5)

	//uma goroutine fica executando ate que a thread main seja finalizada.
}
