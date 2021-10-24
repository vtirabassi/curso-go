package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(param1, param2 string) {
	fmt.Printf("%s - %s \n", param1, param2)
}

func f3() string {
	return "teste"
}

func f4() (string, string) {
	return "Retorno1", "Retorno2"
}

func main() {
	f1()
	f2("parametro1", "parametro2")
	f3()

	r1, r2 := f4()
	_, r3 := f4()

	fmt.Printf("Retornos teste1: %s, %s \n", r1, r2)
	fmt.Printf("Retornos teste2: %s \n", r3)
}
