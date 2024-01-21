package main

import "math"

const (
	gravity  = 0.98
	friction = 0.98
	spring   = 0.4
	bounce   = 0.3
)

type Calc struct {
	World World
}

func (u *Calc) Fruits(fruits []*Fruit) {
	u.move(fruits)
	u.hitTest(fruits)
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

func (u *Calc) hitTest(fruits []*Fruit) {
	l := len(fruits)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			f := fruits[i]
			g := fruits[j]
			dx := g.X - f.X
			dy := g.Y - f.Y
			d := math.Sqrt(dx*dx + dy*dy)
			minD := f.Radius + g.Radius
			if d < minD {
				// collision
				angle := math.Atan2(dy, dx)
				tx := f.X + math.Cos(angle)*minD
				ty := f.Y + math.Sin(angle)*minD
				ax := (tx - g.X) * spring
				ay := (ty - g.Y) * spring
				f.VX -= ax
				f.VY -= ay
				g.VX += ax
				g.VY += ay

				f.X = f.X - math.Cos(angle)*(minD-d)/2
				f.Y = f.Y - math.Sin(angle)*(minD-d)/2
				g.X = g.X + math.Cos(angle)*(minD-d)/2
				g.Y = g.Y + math.Sin(angle)*(minD-d)/2
			}
		}
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
