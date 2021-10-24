package main

func multiplica(x, y int) int {
	return x * y
}

func exec(funcao func(int, int) int, p1 int, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	result := exec(multiplica, 3, 3)

	println(result)
}
