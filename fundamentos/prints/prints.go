package main

import "fmt"

func main() {

	x := 3.1415
	xs := fmt.Sprint(x)

	fmt.Println("o valor de x é " + xs)
	fmt.Printf("o valor de x é %.2f", x)
}
