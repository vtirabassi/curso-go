package main

import (
	f "fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	// forma de criar uma var
	area := PI * math.Pow(raio, 2)

	f.Println(area)

	const (
		a = 1
		b = 2
	)

	f.Println(a, b)

	var tr, fa = true, false

	f.Println(tr, fa)

	x, y, z := true, 2, "teste"

	f.Println(x, y, z)
}
