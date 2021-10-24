package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!") // ele atrasa a execuçao ate o ultimo segundo antes de retornar resultado da func
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}

	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
