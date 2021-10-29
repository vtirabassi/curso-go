package main

import "math"

//se voce iniciar algo com letra maiucusla a visibilidade vai ser public (dentro e fora do pacote)
//se voce inciar algo com letra micuscula a visibilidade vai ser privada (dentro do pacote)

//
//Name -> gera algo public
//name ou _Name gera algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é reponsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
