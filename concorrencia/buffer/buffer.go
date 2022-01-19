package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Executou!")
	ch <- 6
}

func main() {
	//channel com tamanho, voce consegue enviar dados ate o tamanho estabelecido e so ai que ele vai ficar bloqueante
	//ate o channel ser ligo
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch) // lendo os dados do canal
}
