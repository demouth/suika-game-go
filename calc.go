package main

const (
	gravity  = 0.98
	friction = 0.98
	bounce   = 0.3
)

type Calc struct {
	World World
}

func (u *Calc) Fruits(fruits []*Fruit) {
	u.move(fruits)
	u.screenWrap(fruits)
}

func (u *Calc) move(fruits []*Fruit) {
	l := len(fruits)
	for i := 0; i < l; i++ {
		f := fruits[i]
		f.VX *= friction
		f.VY *= friction
		f.VY += gravity
		f.X += f.VX
		f.Y += f.VY
	}
}

func (u *Calc) screenWrap(fruits []*Fruit) {
	l := len(fruits)
	for i := 0; i < l; i++ {
		f := fruits[i]
		if f.X-f.Radius < 0 {
			f.X = f.Radius
			f.VX *= -bounce
		} else if u.World.Width < f.X+f.Radius {
			f.X = u.World.Width - f.Radius
			f.VX *= -bounce
		}
		if f.Y < 0 {
			// no screen wrap
		} else if u.World.Height < f.Y+f.Radius {
			f.Y = u.World.Height - f.Radius
			f.VY *= -bounce
		}
	}
}
