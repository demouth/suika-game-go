package main

type Calc struct {
}

func (u *Calc) Fruits(fruits []*Fruit) {
	u.move(fruits)
}

func (u *Calc) move(fruits []*Fruit) {
	l := len(fruits)
	for i := 0; i < l; i++ {
		f := fruits[i]
		f.Y += 1
	}
}
