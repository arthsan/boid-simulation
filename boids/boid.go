package boids

import (
	"image/color"
	"math/rand"
	"time"
)

const (
	ScreenWidth, ScreenHeight = 640, 360
	BoidCount                 = 500
)

var (
	Green = color.RGBA{10, 255, 50, 255}
	Boids [BoidCount]*Boid
)

type Boid struct {
	Position Vector2D
	Velocity Vector2D
	Id       int
}

func (b *Boid) moveOne() {
	b.Position = b.Position.Add(b.Velocity)
	next := b.Position.Add(b.Velocity)

	if next.X >= ScreenWidth || next.X < 0 {
		b.Velocity = Vector2D{X: -b.Velocity.X, Y: b.Velocity.Y}
	}

	if next.Y >= ScreenHeight || next.Y < 0 {
		b.Velocity = Vector2D{X: b.Velocity.X, Y: -b.Velocity.Y}
	}
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func CreateBoid(bid int) {
	b := Boid{
		Position: Vector2D{X: rand.Float64() * ScreenWidth, Y: rand.Float64() * ScreenHeight},
		Velocity: Vector2D{X: (rand.Float64() * 2) - 1, Y: (rand.Float64() * 2) - 1},
		Id:       bid,
	}

	Boids[bid] = &b
	go b.start()
}
