package point

import "math"

/*
Point представляет точку в двумерном пространстве с помощью двух переменных
типа float64
*/
type Point struct {
	x, y float64
}

// NewPoint создает Point с заданными координатами x и y
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// X возвращает координату X точки p
func (p *Point) X() float64 {
	return p.x
}

// Y возвращает координату Y точки p
func (p *Point) Y() float64 {
	return p.y
}

/*
Distance возвращает расстояние между двумя точками p и q, по формуле евклидова
расстояния.
*/
func Distance(p, q *Point) float64 {
	dx := p.X() - q.X()
	dy := p.Y() - q.Y()
	return math.Sqrt(dx*dx + dy*dy)
}
