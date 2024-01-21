package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 480
	screenHeight = 640
)

type Game struct {
}

var (
	fruits = []*Fruit{}
	world  = World{X: 0, Y: 0, Width: screenWidth, Height: screenHeight}
	next   = NewApple(world.Width/2, 0)

	calc = &Calc{World: world}
	draw = &Draw{}

	isKeyPressed = false
)

func (g *Game) Update() error {
	fruits = calc.Fruits(fruits)

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		next.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		next.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		isKeyPressed = true
	} else if isKeyPressed {
		isKeyPressed = false
		fruits = append(fruits, next)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		f := r.Float64()
		if f < 0.5 {
			next = NewApple(next.X, next.Y)
		} else if f < 0.75 {
			next = NewOrange(next.X, next.Y)
		} else {
			next = NewGrape(next.X, next.Y)
		}
	}

	if next.X-next.Radius < 0 {
		next.X = next.Radius
	}
	if world.Width-next.Radius < next.X {
		next.X = world.Width - next.Radius
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	draw.World(screen, world)
	draw.Fruit(screen, world, next)
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
