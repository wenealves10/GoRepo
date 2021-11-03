package main

import "math"

//Point represents a coordinate in Cartesian plane
type Ponteir struct {
	x float64
	y float64
}

func peccaries(a, b Ponteir) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)

	return
}

//Distance is responsible for calculating the linear line between two points
func Distance(a, b Ponteir) float64 {
	cx, cy := peccaries(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
