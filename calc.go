package main

import "math"

const (
	gravity  = 0.98
	friction = 0.98
	spring   = 0.4
	bounce   = 0.3
)

type Calc struct {
	World   World
	Score   int
	HiScore int
}

func (u *Calc) Fruits(fruits []*Fruit) []*Fruit {
	if u.isGameOver(fruits) {
		if u.HiScore < u.Score {
			u.HiScore = u.Score
		}
		u.Score = 0
		return make([]*Fruit, 0)
	}

	fruits = u.combine(fruits)
	u.hitTest(fruits)
	u.move(fruits)
	u.screenWrap(fruits)
	return fruits
}

func (u *Calc) isGameOver(fruits []*Fruit) bool {
	l := len(fruits)
	for i := 0; i < l; i++ {
		f := fruits[i]
		if f.Y < 0 {
			return true
		}
	}
	return false
}

func (u *Calc) combine(fruits []*Fruit) []*Fruit {
	newFruits := make([]*Fruit, 0)

	l := len(fruits)
	for i := 0; i < l; i++ {
		f := fruits[i]
		for j := i + 1; j < l; j++ {
			g := fruits[j]
			if f.Remove || g.Remove {
				continue
			}
			dx := g.X - f.X
			dy := g.Y - f.Y
			d := math.Sqrt(dx*dx + dy*dy)
			minD := f.Radius + g.Radius
			if d < minD && f.Type == g.Type {
				// collision
				f.Remove = true
				g.Remove = true
				var next *Fruit
				if f.Type == APPLE {
					next = NewOrange((f.X+g.X)/2, (f.Y+g.Y)/2)
					u.Score += 10
				} else if f.Type == ORANGE {
					next = NewGrape((f.X+g.X)/2, (f.Y+g.Y)/2)
					u.Score += 20
				} else if f.Type == GRAPE {
					next = NewPineapple((f.X+g.X)/2, (f.Y+g.Y)/2)
					u.Score += 30
				} else if f.Type == PINEAPPLE {
					next = NewMelon((f.X+g.X)/2, (f.Y+g.Y)/2)
					u.Score += 40
				} else if f.Type == MELON {
					next = NewWatermelon((f.X+g.X)/2, (f.Y+g.Y)/2)
					u.Score += 50
				} else if f.Type == WATERMELON {
					u.Score += 60
				}
				if next != nil {
					newFruits = append(newFruits, next)
				}
			}
		}
	}
	for i := 0; i < l; i++ {
		f := fruits[i]
		if !f.Remove {
			newFruits = append(newFruits, f)
		}
	}
	return newFruits
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
