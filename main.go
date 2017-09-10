package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 320
	screenHeight = 420
)

const (
	landHeight    = 12
	characterSize = 8
)

func main() {
	err := ebiten.Run(update, screenWidth, screenHeight, 2, "Gopher")
	if err != nil {
		panic(err)
	}
}

// Ebiten images
var character *ebiten.Image
var landmass *ebiten.Image

// Characters positions - saved here as global state variables
var posX float64
var posY float64

func update(screen *ebiten.Image) error {

	// x, y := ebiten.CursorPosition()
	drawCharacter(screen)
	drawLand(screen)
	handleInput()

	return nil
}

func drawCharacter(s *ebiten.Image) {
	//If square does not already exist initialise it - this way a new image is not create it each time
	if character == nil {
		character, _ = ebiten.NewImage(characterSize, characterSize, ebiten.FilterNearest)
	}
	character.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	opts := &ebiten.DrawImageOptions{}
	// posOffset := posY - float64(screenHeight+characterSize+landHeight)
	opts.GeoM.Translate(posX, posY)
	s.DrawImage(character, opts)
}

func drawLand(s *ebiten.Image) {
	if landmass == nil {
		landmass, _ = ebiten.NewImage(screenWidth, landHeight, ebiten.FilterNearest)
	}
	landmass.Fill(color.White)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, screenHeight-landHeight)
	s.DrawImage(landmass, opts)
}

func handleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if posY > 0 || posY < screenHeight {
			posY -= 10
		} else {
			posY = 0
		}
	}
	// When the "down arrow keposY" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if posY > 0 || posY < screenHeight {
			posY += 10
		} else {
			posX = 0
		}
	}
	// When the "left arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if posX > 0 || posX < screenWidth {
			posX -= 10
		} else {
			posX = 0
		}
	}
	// When the "right arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if posX > 0 || posX < screenWidth {
			posX += 10
		} else {
			posX = 0
		}
	}
}
