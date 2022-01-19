package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante, necessario o dado ser lido para liberar
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) //canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c)   // deadlock, pois nao tem como chegar dados nesse canal, devido as goroutines terem sido finalizadas
	fmt.Println("Fim") // não será exibido
}
