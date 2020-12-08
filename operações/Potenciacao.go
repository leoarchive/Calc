package calc

import "math"

func Potenciacao(primeiro, segundo float64) float64 {
	resultado := math.Pow(float64(primeiro), float64(segundo))
	return resultado
}