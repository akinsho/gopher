package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	screenWidth  = 640
	screenHeight = 320
)

const (
	landHeight    = 12
	characterSize = 8
	grassHeight   = 4
)

func main() {
	err := ebiten.Run(update, screenWidth, screenHeight, 2, "Gopher")
	if err != nil {
		panic(err)
	}
}

// Color Palette
var sky color.Color

// Ebiten images
var character *ebiten.Image
var landmass *ebiten.Image
var grass *ebiten.Image

// Characters positions - saved here as global state variables
var posX float64
var posY float64

func update(screen *ebiten.Image) error {

	// x, y := ebiten.CursorPosition()
	if sky == nil {
		sky, _ = colorful.Hex("#5FE8F7")
	}
	//The Screen must first be filled otherwise it will cover everything else
	screen.Fill(sky)

	drawCharacter(screen)
	drawLand(screen)
	handleInput()
	return nil
}

func logError(message string, err error) {
	log.Printf("%s: %v", message, err)
}

func drawCharacter(s *ebiten.Image) {
	//If square does not already exist initialise it - this way a new image is not create it each time
	if character == nil {
		character, _ = ebiten.NewImage(characterSize, characterSize, ebiten.FilterNearest)
	}
	character.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	opts := &ebiten.DrawImageOptions{}
	posOffset := posY + float64(screenHeight-characterSize-landHeight-grassHeight)
	opts.GeoM.Translate(posX, posOffset)
	s.DrawImage(character, opts)
}

func drawLand(s *ebiten.Image) {
	if landmass == nil {
		landmass, _ = ebiten.NewImage(screenWidth, landHeight, ebiten.FilterNearest)
		grass, _ = ebiten.NewImage(screenWidth, grassHeight, ebiten.FilterNearest)
	}
	brown, err := colorful.Hex("#895C22")
	if err != nil {
		logError("Color Error", err)
	}
	green, err := colorful.Hex("#53D46B")
	if err != nil {
		logError("Color Error", err)
	}

	landmass.Fill(brown)
	grass.Fill(green)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, screenHeight-landHeight)
	s.DrawImage(landmass, opts)

	grassOpts := &ebiten.DrawImageOptions{}
	grassOpts.GeoM.Translate(0, screenHeight-landHeight-grassHeight)
	s.DrawImage(grass, grassOpts)
}

func handleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if posY > 0 || posY < screenHeight+landHeight {
			posY -= 2
		} else {
			posY = 0
		}
	}
	// When the "down arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if posY > 0 || posY < screenHeight {
			posY += 2
		} else {
			posX = 0
		}
	}
	// When the "left arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if posX > 0 || posX < screenWidth {
			posX -= 2
		} else {
			posX = 0
		}
	}
	// When the "right arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if posX > 0 || posX < screenWidth {
			posX += 2
		} else {
			posX = 0
		}
	}
}
