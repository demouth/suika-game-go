package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	appleImage      *ebiten.Image
	grapeImage      *ebiten.Image
	orangeImage     *ebiten.Image
	pineappleImage  *ebiten.Image
	melonImage      *ebiten.Image
	watermelonImage *ebiten.Image

	//go:embed assets/apple.png
	apple_png []byte
	//go:embed assets/grape.png
	grape_png []byte
	//go:embed assets/orange.png
	orange_png []byte
	//go:embed assets/pineapple.png
	pineapple_png []byte
	//go:embed assets/melon.png
	melon_png []byte
	//go:embed assets/watermelon.png
	watermelon_png []byte
)

type Draw struct {
	op ebiten.DrawImageOptions
}

func init() {
	appleImage = loadImage(apple_png)
	orangeImage = loadImage(orange_png)
	grapeImage = loadImage(grape_png)
	pineappleImage = loadImage(pineapple_png)
	melonImage = loadImage(melon_png)
	watermelonImage = loadImage(watermelon_png)
}

func loadImage(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	origImage := ebiten.NewImageFromImage(img)

	s := origImage.Bounds().Size()
	ebitenImage := ebiten.NewImage(s.X, s.Y)
	op := &ebiten.DrawImageOptions{}
	ebitenImage.DrawImage(origImage, op)
	return ebitenImage
}

func (d *Draw) World(screen *ebiten.Image, world World) {
	vector.DrawFilledRect(
		screen,
		float32(world.X), float32(world.Y), float32(world.Width), float32(world.Height),
		color.RGBA{0x66, 0x66, 0x66, 0xff},
		false,
	)
}

func (d *Draw) Fruit(screen *ebiten.Image, world World, f *Fruit) {
	var img *ebiten.Image
	switch {
	case f.Type == APPLE:
		img = appleImage
	case f.Type == ORANGE:
		img = orangeImage
	case f.Type == GRAPE:
		img = grapeImage
	case f.Type == PINEAPPLE:
		img = pineappleImage
	case f.Type == MELON:
		img = melonImage
	case f.Type == WATERMELON:
		img = watermelonImage
	default:
		img = appleImage
	}

	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	d.op.Filter = ebiten.FilterLinear
	d.op.GeoM.Reset()
	d.op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	d.op.GeoM.Scale(f.Radius/float64(w)*2, f.Radius/float64(h)*2)
	d.op.GeoM.Translate(float64(world.X), float64(world.Y))
	d.op.GeoM.Translate(float64(f.X), float64(f.Y))
	screen.DrawImage(img, &d.op)
}

func (d *Draw) Fruits(screen *ebiten.Image, world World, fruits []*Fruit) {
	l := len(fruits)
	for i := 0; i < l; i++ {
		f := fruits[i]
		d.Fruit(screen, world, f)
	}
}
