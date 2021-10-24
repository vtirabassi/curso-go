package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta e determina o tamanho com base nos valores definidos

	for i, numero := range numeros { //i -> inidice, numero -> vslor do indice
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { // nao estou usando o indice, ai no caso ignoro usanso o caracter _
		fmt.Println(num)
	}
}
