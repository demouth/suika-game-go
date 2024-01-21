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
		NewApple(100, 0),
		NewApple(110, -100),
		NewOrange(110, -2000),
		NewGrape(110, -3000),
		NewPineapple(140, -4000),
		NewMelon(150, -5000),
		NewWatermelon(100, -6000),
	}
	world = World{X: 0, Y: 0, Width: screenWidth, Height: screenHeight}

	calc = &Calc{World: world}
	draw = &Draw{}
)

func (g *Game) Update() error {
	fruits = calc.Fruits(fruits)
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
