package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

var img *ebiten.Image

type Game struct{}

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("assets/gopher.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(1, 1)
	op.GeoM.Scale(.5, .5)
	screen.DrawImage(img, op)
	ebitenutil.DebugPrint(screen, "Matze!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 240, 160
}

func main() {
	ebiten.SetWindowSize(480, 320)
	ebiten.SetWindowTitle("Matze!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
