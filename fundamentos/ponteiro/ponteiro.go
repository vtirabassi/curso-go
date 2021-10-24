package main

import "fmt"

func main() {
	i := 1
	var p *int = nil //criando um ponteiro (ref de memoria)

	p = &i // pegando o endereço da variável
	*p++   // desreferenciando (pegando o valor) e somando 1
	i++

	// Go não tem aritmética de ponteiros
	// p++
	fmt.Println(p, *p, i, &i)
}
