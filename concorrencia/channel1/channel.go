package main

import (
	"fmt"
)

func main() {

	//channel é um tipo da linguagem, como um inteiro, boolean
	// Channel (canal) - é a forma de comunicação entre goroutines
	// Channel é um tipo

	ch := make(chan int, 1)

	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
