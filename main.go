package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 480
	screenHeight = 640
)

type Game struct {
	touchIDs []ebiten.TouchID
}

var (
	fruits = []*Fruit{}
	world  = World{X: 120, Y: 100, Width: 240, Height: 540}
	next   = NewApple(world.Width/2, 0)

	calc = &Calc{World: world}
	draw = &Draw{}

	isKeyPressed = false
)

func (g *Game) leftTouched() bool {
	for _, id := range g.touchIDs {
		x, _ := ebiten.TouchPosition(id)
		if x < screenWidth/2 {
			return true
		}
	}
	return false
}

func (g *Game) rightTouched() bool {
	for _, id := range g.touchIDs {
		x, _ := ebiten.TouchPosition(id)
		if x >= screenWidth/2 {
			return true
		}
	}
	return false
}

func (g *Game) bottomTouched() bool {
	for _, id := range g.touchIDs {
		_, y := ebiten.TouchPosition(id)
		if y >= screenHeight/2 {
			return true
		}
	}
	return false
}

func (g *Game) Update() error {
	fruits = calc.Fruits(fruits)

	g.touchIDs = ebiten.AppendTouchIDs(g.touchIDs[:0])

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || g.leftTouched() {
		next.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || g.rightTouched() {
		next.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) || g.rightTouched() {
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
	msg := fmt.Sprintf(
		"<-: move left\n->: move right\nspace: drop fruit\nHI-SCORE: %d\nSCORE: %d\nFPS: %0.2f",
		calc.HiScore,
		calc.Score,
		ebiten.ActualFPS(),
	)
	ebitenutil.DebugPrint(screen, msg)
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
