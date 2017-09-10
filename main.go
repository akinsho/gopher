package main

// TODO
// 1. clouds should repeat the movements - i.e loop
// 2. clarify collision borders
// 3. complete out of bounds management

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
var islandOne *ebiten.Image
var islandTwo *ebiten.Image
var character *ebiten.Image
var landmass *ebiten.Image
var grass *ebiten.Image

// Characters positions - saved here as global state variables
var posX float64
var posY = float64(lowerBound)
var islandOneX float64 = screenWidth - 145
var islandOneY float64 = screenHeight - 60
var islandTwoX float64
var islandTwoY float64 = screenHeight - 50
var cloudX float64
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
