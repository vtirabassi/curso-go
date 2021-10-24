package main

import (
	"fmt"
	"time"
)

func notaConceito(n float64) string {
	nota := int(n)

	switch nota {
	case 7, 8:
		return "B"
	case 9:
		fallthrough
	case 10:
		return "A"
	default:
		return "C"
	}
}

func periodoHorario(hour int) string {
	switch true {
	case hour <= 12:
		return "Bom dia"
	case hour > 12 && hour <= 18:
		return "Boa tarde"
	case hour > 18 && hour <= 0:
		return "Boa noite"
	default:
		return "nao foi possivel"
	}
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "nao sei"
	}
}

func main() {

	n := 9.1
	nC := notaConceito(n)

	fmt.Printf("Nota conceito da nota %0.2f é: %s", n, nC)

	t := time.Now()
	fmt.Println(periodoHorario(t.Hour()))

	fmt.Println(tipo(4.99))
}
