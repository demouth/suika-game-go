package main

const (
	gravity = 0.4
)

type Calc struct {
	World World
}

func (u *Calc) Fruits(fruits []*Fruit) {
	u.move(fruits)
}

func (u *Calc) move(fruits []*Fruit) {
	l := len(fruits)
	for i := 0; i < l; i++ {
		f := fruits[i]
		f.VY += gravity
		f.X += f.VX
		f.Y += f.VY
	}
}
