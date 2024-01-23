package main

import (
	"math/rand"
	"time"
)

type Dropper struct {
	World World
	next  *Fruit
}

func NewDropper(world World) *Dropper {
	d := &Dropper{World: world}
	d.Drop()
	d.next.X = world.Width / 2
	return d
}

func (d *Dropper) MoveLeft() {
	if d.next == nil {
		return
	}
	d.next.X -= 2
	if d.next.X-d.next.Radius < 0 {
		d.next.X = d.next.Radius
	}
}

func (d *Dropper) MoveRight() {
	if d.next == nil {
		return
	}
	d.next.X += 2
	if d.World.Width-d.next.Radius < d.next.X {
		d.next.X = d.World.Width - d.next.Radius
	}
}

func (d *Dropper) Next() *Fruit {
	return d.next
}

func (d *Dropper) Tick() {

}

func (d *Dropper) Drop() *Fruit {
	var x float64
	var y float64
	if d.next != nil {
		x = d.next.X
		y = d.next.Y
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	f := r.Float64()
	var next *Fruit
	if f < 0.5 {
		next = NewApple(x, y)
	} else if f < 0.75 {
		next = NewOrange(x, y)
	} else {
		next = NewGrape(x, y)
	}

	ret := d.next
	d.next = next

	return ret
}
