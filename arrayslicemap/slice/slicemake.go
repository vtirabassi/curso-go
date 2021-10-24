package main

import "fmt"

func main() {
	s := make([]int, 10) //cria um slice sem base em um array, com isso o GO cria um array interno
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) //faz o GO criar um array interno com 20posicoes
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
