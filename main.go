package main

import (
	"log"

	bd "github.com/arthsan/boid-simulation/boids"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range bd.Boids {
		screen.Set(int(boid.Position.X+1), int(boid.Position.Y), bd.Green)
		screen.Set(int(boid.Position.X-1), int(boid.Position.Y), bd.Green)
		screen.Set(int(boid.Position.X), int(boid.Position.Y+1), bd.Green)
		screen.Set(int(boid.Position.X), int(boid.Position.Y-1), bd.Green)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return bd.ScreenWidth, bd.ScreenHeight
}

func main() {
	for i := 0; i < bd.BoidCount; i++ {
		bd.CreateBoid(i)
	}

	ebiten.SetWindowSize(bd.ScreenWidth*2, bd.ScreenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
