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

var (
	onGround = true
	// Color Palette
	sky color.Color
	// Ebiten images
	islandOne *ebiten.Image
	islandTwo *ebiten.Image
	character *ebiten.Image
	landmass  *ebiten.Image
	grass     *ebiten.Image
)

// Characters positions - saved here as global state variables
var (
	posX          float64
	posY                  = float64(lowerBound)
	islandOneX    float64 = screenWidth - 145
	islandOneY    float64 = screenHeight - 60
	islandTwoX    float64
	islandTwoY    float64 = screenHeight - 50
	size          float64 = 10
	islandYLength float64 = 125
	islandXLength float64 = 200
	cloudX        float64
	velocityY     float64
	gravity       = 0.5
)

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
