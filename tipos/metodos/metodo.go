package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setPessoa(nomeCompleto string) {
	nomes := strings.Split(nomeCompleto, " ")
	p.nome = nomes[0]
	p.sobrenome = nomes[1]
}

func main() {

	pessoa1 := pessoa{nome: "Vinicius", sobrenome: "Tirabassi"}
	fmt.Printf("O nome completo da pessoa é : %s \n", pessoa1.getNomeCompleto())

	nomeCompleto := "Sthefanie Carvalho"
	pessoa1.setPessoa(nomeCompleto)

	fmt.Printf("O primeiro nome é %s e o segundo nome é %s \n", pessoa1.nome, pessoa1.sobrenome)
}
