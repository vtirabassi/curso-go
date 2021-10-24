package main

import "fmt"

func media(numeros ...float64) float64 { //recebe x quantidades de parametros
	total := 0.0

	for _, num := range numeros { //ignorando o indice
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.9))
}
