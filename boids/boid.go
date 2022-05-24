package boids

import (
	"image/color"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	ScreenWidth, ScreenHeight = 640, 360
	BoidCount                 = 500
	ViewRadius                = 13
	AdjustingRate             = 0.03
)

var (
	Green   = color.RGBA{10, 255, 50, 255}
	Boids   [BoidCount]*Boid
	BoidMap [ScreenWidth + 1][ScreenHeight + 1]int
	rWlock  = sync.RWMutex{}
)

type Boid struct {
	Position Vector2D
	Velocity Vector2D
	Id       int
}

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.Position.AddValue(ViewRadius), b.Position.AddValue(-ViewRadius)
	avgPosition, avgVelocity, separation := Vector2D{X: 0, Y: 0}, Vector2D{X: 0, Y: 0}, Vector2D{X: 0, Y: 0}
	count := 0.0

	rWlock.RLock()
	for i := math.Max(lower.X, 0); i <= math.Min(upper.X, ScreenWidth); i++ {
		for j := math.Max(lower.Y, 0); j <= math.Min(upper.Y, ScreenHeight); j++ {
			if otherBoidId := BoidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.Id {
				if dist := Boids[otherBoidId].Position.Distance(b.Position); dist < ViewRadius {
					count++
					avgVelocity = avgVelocity.Add(Boids[otherBoidId].Velocity)
					avgPosition = avgPosition.Add(Boids[otherBoidId].Position)
					separation = separation.Add(b.Position.Subtract(Boids[otherBoidId].Position).DivisionValue(dist))
				}
			}
		}
	}
	rWlock.RUnlock()

	accel := Vector2D{X: b.borderBounce(b.Position.X, ScreenWidth), Y: b.borderBounce(b.Position.Y, ScreenHeight)}
	if count > 0 {
		avgPosition, avgVelocity = avgPosition.DivisionValue(count), avgVelocity.DivisionValue(count)
		accelAlignment := avgVelocity.Subtract(b.Velocity).MultiplYValue(AdjustingRate)
		accelCohesion := avgPosition.Subtract(b.Position).MultiplYValue(AdjustingRate)
		accelSeparation := separation.MultiplYValue(AdjustingRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	}

	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < ViewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-ViewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()
	rWlock.Lock()
	b.Velocity = b.Velocity.Add(acceleration).limit(-1, 1)
	BoidMap[int(b.Position.X)][int(b.Position.Y)] = -1
	b.Position = b.Position.Add(b.Velocity)
	BoidMap[int(b.Position.X)][int(b.Position.Y)] = b.Id
	rWlock.Unlock()
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
	BoidMap[int(b.Position.X)][int(b.Position.Y)] = b.Id
	go b.start()
}
