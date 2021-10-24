package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0 //add elemento no map
	delete(funcsESalarios, "Inexistente")   //nao vai dar erro no programa

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
