package main

import "math"

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

// Distancia é reponsável por calcular a fistância linear entre dois pontos
func Distancia(a, b Ponto) float32 {
	cx, cy := catetos(a, b)
	return float32(math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2)))
}
