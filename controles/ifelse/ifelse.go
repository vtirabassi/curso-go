package main

import "fmt"

func imprimirResultado(nota float32) {
	if nota >= 7 {
		fmt.Println("Aluno aprovado com nota", nota)
	} else {
		fmt.Println("Aluno reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(8)
	imprimirResultado(5.1)
}
