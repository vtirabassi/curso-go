package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	fmt.Println(1, 2, 10000)
	fmt.Println("o tipo da variavel é ", reflect.TypeOf(12345))

	//sem sinal, so positivos, uint8, uint16, uint32, uint64
	//com sinal, int8, int16, int32, int64

	maxInt := math.MaxInt64
	fmt.Println(maxInt)

	//numeros float32, float64
	var x float32 = 45.66
	fmt.Println("o tipo da variavel x é", reflect.TypeOf(x))

	//boleand
	y, z := true, false
	fmt.Print(y)
	fmt.Println("o tipo da variavel z é", reflect.TypeOf(z))

	//string
	s := `vincius
	tirabassi`

	fmt.Println("tamanhao da string é:", len(s))
}
