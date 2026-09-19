package main

import "math"

const (
	gravity  = 0.4
	friction = 0.98
	bounce   = 0.3

	floorFriction = 0.97
	restitution   = 0.2
	restThreshold = 1.0
	iterations    = 6
	contactEps    = 0.01
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
	u.move(fruits)
	for i := 0; i < iterations; i++ {
		u.hitTest(fruits)
		u.clampWalls(fruits)
	}
	u.wallResponse(fruits)
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
	for _, f := range fruits {
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
			if d >= minD {
				continue
			}

			nx, ny := 1.0, 0.0
			if d > 0 {
				nx, ny = dx/d, dy/d
			}

			mf, mg := f.Mass(), g.Mass()
			rf := mg / (mf + mg)
			rg := mf / (mf + mg)

			overlap := minD - d
			f.X -= nx * overlap * rf
			f.Y -= ny * overlap * rf
			g.X += nx * overlap * rg
			g.Y += ny * overlap * rg

			vrel := (g.VX-f.VX)*nx + (g.VY-f.VY)*ny
			if vrel >= 0 {
				continue
			}
			e := restitution
			if -vrel < restThreshold {
				e = 0
			}
			j2 := -(1 + e) * vrel / (1/mf + 1/mg)
			f.VX -= j2 / mf * nx
			f.VY -= j2 / mf * ny
			g.VX += j2 / mg * nx
			g.VY += j2 / mg * ny
		}
	}
}

func (u *Calc) clampWalls(fruits []*Fruit) {
	for _, f := range fruits {
		if f.X-f.Radius < 0 {
			f.X = f.Radius
		} else if u.World.Width < f.X+f.Radius {
			f.X = u.World.Width - f.Radius
		}
		if 0 <= f.Y && u.World.Height < f.Y+f.Radius {
			f.Y = u.World.Height - f.Radius
		}
	}
}

func (u *Calc) wallResponse(fruits []*Fruit) {
	for _, f := range fruits {
		if f.X-f.Radius <= contactEps && f.VX < 0 {
			f.VX = wallBounce(f.VX)
		} else if u.World.Width <= f.X+f.Radius+contactEps && f.VX > 0 {
			f.VX = wallBounce(f.VX)
		}
		if 0 <= f.Y && u.World.Height <= f.Y+f.Radius+contactEps {
			if f.VY > 0 {
				f.VY = wallBounce(f.VY)
			}
			f.VX *= floorFriction
		}
	}
}

func wallBounce(v float64) float64 {
	if math.Abs(v) < restThreshold {
		return 0
	}
	return -v * bounce
}
