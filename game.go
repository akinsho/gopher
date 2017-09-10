package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/nfnt/resize"
)

func drawCharacter(s *ebiten.Image) {
	//If square does not already exist initialise it - this way a new image is not created each time
	if character == nil {
		character, _ = ebiten.NewImage(characterSize, characterSize, ebiten.FilterNearest)
	}
	character.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(posX, posY)

	s.DrawImage(character, opts)
}

func detectCollision() (b float64) {
	if posX >= islandOneX && posY <= islandOneY {
		onGround = true
		return islandOneY - characterSize
	} else if posX <= islandTwoX+200 && posY <= islandTwoY {
		onGround = true
		return islandTwoY - characterSize
	}
	onGround = true
	return lowerBound
}

func drawLand(s *ebiten.Image) {
	if landmass == nil {
		landmass, _ = ebiten.NewImage(screenWidth, landHeight, ebiten.FilterNearest)
		islandOne, _ = ebiten.NewImage(125, 10, ebiten.FilterNearest)
		islandTwo, _ = ebiten.NewImage(200, 10, ebiten.FilterNearest)
		grass, _ = ebiten.NewImage(screenWidth, grassHeight, ebiten.FilterNearest)
	}
	brown, err := colorful.Hex("#895C22")
	logError(err)
	green, err := colorful.Hex("#53D46B")
	logError(err)

	landmass.Fill(brown)
	islandOne.Fill(brown)
	islandTwo.Fill(brown)
	grass.Fill(green)

	iOneOpts := &ebiten.DrawImageOptions{}
	iOneOpts.GeoM.Translate(islandOneX, islandOneY)
	s.DrawImage(islandOne, iOneOpts)

	iTwoOpts := &ebiten.DrawImageOptions{}
	iTwoOpts.GeoM.Translate(islandTwoX, islandTwoY)
	s.DrawImage(islandTwo, iTwoOpts)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, screenHeight-landHeight)
	s.DrawImage(landmass, opts)

	grassOpts := &ebiten.DrawImageOptions{}
	grassOpts.GeoM.Translate(0, lowerBound+characterSize)
	s.DrawImage(grass, grassOpts)
}

func getImage() (i image.Image) {
	file, err := os.Open("assets/cloud.png")
	defer file.Close()
	logError(err)

	img, err := png.Decode(file)
	logError(err)

	resized := resize.Resize(50, 0, img, resize.Lanczos3)
	return resized
}

func drawClouds(s *ebiten.Image) {
	img := getImage()
	oneOpts := &ebiten.DrawImageOptions{}
	twoOpts := &ebiten.DrawImageOptions{}
	one, err := ebiten.NewImageFromImage(img, ebiten.FilterNearest)
	two, err := ebiten.NewImageFromImage(img, ebiten.FilterNearest)
	if cloudX < screenWidth-20 || cloudX > 0 {
		cloudX += 0.5
	} else {
		log.Println(cloudX)
		cloudX = 0
	}
	oneOpts.GeoM.Translate(cloudX, 5)
	twoOpts.GeoM.Translate(cloudX/2, 100)
	s.DrawImage(one, oneOpts)
	s.DrawImage(two, twoOpts)
	logError(err)
}
