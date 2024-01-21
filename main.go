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
		{X: 100, Y: 100, VX: -15, VY: 0, Radius: 25},
		{X: 250, Y: 100, VX: 15, VY: 0, Radius: 50},
		{X: 250, Y: -100, VX: 0, VY: 0, Radius: 25},
		{X: 100, Y: -100, VX: 0, VY: 0, Radius: 25},
		{X: 200, Y: -200, VX: 0, VY: 0, Radius: 25},
		{X: 150, Y: -300, VX: 0, VY: 0, Radius: 50},
		{X: 200, Y: -400, VX: 0, VY: 0, Radius: 25},
		{X: 200, Y: -500, VX: 0, VY: 0, Radius: 25},
		{X: 150, Y: -600, VX: 0, VY: 0, Radius: 75},
	}
	world = World{X: 0, Y: 0, Width: screenWidth, Height: screenHeight}

	calc = &Calc{World: world}
	draw = &Draw{}
)

func (g *Game) Update() error {
	calc.Fruits(fruits)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	draw.World(screen, world)
	draw.Fruits(screen, world, fruits)
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
