package boids

import "math"

type Vector2D struct {
	X float64
	Y float64
}

func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func (v1 Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}

func (v1 Vector2D) MultiplY(v2 Vector2D) Vector2D {
	return Vector2D{X: v1.X * v2.X, Y: v1.Y * v2.Y}
}

func (v1 Vector2D) AddValue(d float64) Vector2D {
	return Vector2D{X: v1.X + d, Y: v1.Y + d}
}

func (v1 Vector2D) MultiplYValue(d float64) Vector2D {
	return Vector2D{X: v1.X * d, Y: v1.Y * d}
}

func (v1 Vector2D) DivisionValue(d float64) Vector2D {
	return Vector2D{X: v1.X / d, Y: v1.Y / d}
}

func (v1 Vector2D) limit(lower, upper float64) Vector2D {
	return Vector2D{
		X: math.Min(math.Max(v1.X, lower), upper),
		Y: math.Min(math.Max(v1.Y, lower), upper),
	}
}

func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v1.X-v2.X, 2) + math.Pow(v1.Y-v2.Y, 2))
}
