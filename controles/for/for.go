package main

import "fmt"

func main() {
	i := 10
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
}
