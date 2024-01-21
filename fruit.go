package main

const (
	APPLE = iota
	GRAPE
	ORANGE
	PINEAPPLE
	MELON
	WATERMELON
)

type Fruit struct {
	X      float64
	Y      float64
	VX     float64
	VY     float64
	Radius float64
	Type   int
	Remove bool
}

func NewApple(x float64, y float64) *Fruit {
	return &Fruit{
		X:      x,
		Y:      y,
		Radius: 20,
		Type:   APPLE,
	}
}

func NewOrange(x float64, y float64) *Fruit {
	return &Fruit{
		X:      x,
		Y:      y,
		Radius: 35,
		Type:   ORANGE,
	}
}

func NewGrape(x float64, y float64) *Fruit {
	return &Fruit{
		X:      x,
		Y:      y,
		Radius: 50,
		Type:   GRAPE,
	}
}
func NewPineapple(x float64, y float64) *Fruit {
	return &Fruit{
		X:      x,
		Y:      y,
		Radius: 65,
		Type:   PINEAPPLE,
	}
}
func NewMelon(x float64, y float64) *Fruit {
	return &Fruit{
		X:      x,
		Y:      y,
		Radius: 80,
		Type:   MELON,
	}
}

func NewWatermelon(x float64, y float64) *Fruit {
	return &Fruit{
		X:      x,
		Y:      y,
		Radius: 95,
		Type:   WATERMELON,
	}
}
