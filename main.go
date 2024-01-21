package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 480
	screenHeight = 640
)

type Game struct {
}

var (
	fruits = []*Fruit{
		{X: 100, Y: 100, Radius: 25},
		{X: 250, Y: 100, Radius: 50},
	}
	draw = &Draw{}
)

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	draw.Fruits(screen, fruits)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("suika-game-go")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
