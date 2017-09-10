package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 320
)

const (
	landHeight    = 12
	characterSize = 8
	grassHeight   = 4
	lowerBound    = screenHeight - (landHeight + characterSize + grassHeight)
)

var onGround = true

// Color Palette
var sky color.Color

// Ebiten images
var character *ebiten.Image
var landmass *ebiten.Image
var grass *ebiten.Image

// Characters positions - saved here as global state variables
var posX float64
var posY = float64(lowerBound)
var velocityY float64
var gravity = 0.5

func main() {
	err := ebiten.Run(update, screenWidth, screenHeight, 2, "Gopher")
	if err != nil {
		panic(err)
	}
}

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
