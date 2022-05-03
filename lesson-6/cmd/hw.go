package hw

import (
	"fmt"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.

type Geom struct {
	X1, Y1, X2, Y2 float64
}

func New(X1, Y1, X2, Y2 float64) *Geom {
	g := Geom{
		X1: X1,
		Y1: Y1,
		X2: X2,
		Y2: Y2,
	}
	return &g
}
func (geom *Geom) CalculateDistance() (distance float64) {

	if geom.X1 < 0 || geom.X2 < 0 || geom.Y1 < 0 || geom.Y2 < 0 {
		fmt.Println("Координаты не могут быть меньше нуля")
		return -1
	} else {
		distance = math.Sqrt(math.Pow(geom.X2-geom.X1, 2) + math.Pow(geom.Y2-geom.Y1, 2))
	}

	// возврат расстояния между точками
	return distance
}
