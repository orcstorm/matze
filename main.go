package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func main() {
	if err := ebiten.Run(update, 240, 160, 2, "Matze!"); err != nil {
		panic(err)
	}
}

func update(screen *ebiten.Image) error {
	ebitenutil.DebugPrint(screen, "Matze!")
	return nil
}
