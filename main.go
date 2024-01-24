package main

import (
	"fmt"
	"log"

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
	world  = World{X: 90, Y: 100, Width: 300, Height: 540}

	dropper = NewDropper(world)
	calc    = &Calc{World: world}
	draw    = &Draw{}

	isKeyPressed = false
)

func (g *Game) leftTouched() bool {
	for _, id := range g.touchIDs {
		x, y := ebiten.TouchPosition(id)
		if y >= screenHeight/2 {
			return false
		}
		if x < screenWidth/2 {
			return true
		}
	}
	return false
}

func (g *Game) rightTouched() bool {
	for _, id := range g.touchIDs {
		x, y := ebiten.TouchPosition(id)
		if y >= screenHeight/2 {
			return false
		}
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
		dropper.MoveLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || g.rightTouched() {
		dropper.MoveRight()
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) || g.bottomTouched() {
		isKeyPressed = true
	} else if isKeyPressed {
		isKeyPressed = false
		if next := dropper.Drop(); next != nil {
			fruits = append(fruits, next)
		}
	}

	dropper.Tick()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	draw.World(screen, world)
	if next := dropper.Next(); next != nil {
		draw.Fruit(screen, world, next)
	}
	draw.Fruits(screen, world, fruits)
	msg := fmt.Sprintf(
		"PC:\n  <- key: move left\n  -> key: move right\n  spacebar: drop fruit\n"+
			"Touch Devices:\n  left: move left\n  right: move right\n  bottom: drop fruit\n"+
			"HI-SCORE: %d\nSCORE: %d\nFPS: %0.2f",
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
