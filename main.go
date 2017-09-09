package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func main() {
	err := ebiten.Run(update, 300, 300, 1, "A Test")
	if err != nil {
		log.Fatal(err)
	}
}

var square *ebiten.Image

//DrawImageOptions is a config option struct
type DrawImageOptions struct{}

func update(screen *ebiten.Image) error {
	if square != nil {
		square, _ = ebiten.NewImage(10, 10, ebiten.FilterNearest)
	}

	square.Fill(color.White)
	ebitenutil.DebugPrint(screen, "A Test")
	opts := &ebiten.DrawImageOptions{}

	screen.DrawImage(square, opts)

	return nil
}
