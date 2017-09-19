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

var brown, _ = colorful.Hex("#895C22")
var green, _ = colorful.Hex("#53D46B")

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

func drawEachLandmass(s *ebiten.Image, sizeX int, sizeY int, coordsX int, coordsY int) {
	mass, _ := ebiten.NewImage(sizeX, sizeY, ebiten.FilterNearest)
	mass.Fill(brown)
	foliage, _ := ebiten.NewImage(sizeX, sizeY/3, ebiten.FilterNearest)
	foliage.Fill(green)
	massOpts := &ebiten.DrawImageOptions{}
	massOpts.GeoM.Translate(float64(coordsX), float64(coordsY))
	foliageOpts := &ebiten.DrawImageOptions{}
	foliageOpts.GeoM.Translate(float64(coordsX), float64(coordsY-1))
	s.DrawImage(mass, massOpts)
	s.DrawImage(foliage, foliageOpts)
}

func drawLand(s *ebiten.Image) {
	if landmass == nil {
		landmass, _ = ebiten.NewImage(screenWidth, landHeight, ebiten.FilterNearest)
		grass, _ = ebiten.NewImage(screenWidth, grassHeight, ebiten.FilterNearest)
	}

	drawEachLandmass(s, 10, 100, screenWidth/3, screenHeight-10)
	drawEachLandmass(s, 125, 10, int(islandTwoX), int(islandTwoY))
	drawEachLandmass(s, 200, 10, int(islandOneX), int(islandOneY-10))
	// TODO loop through all the maps and call drawEachLandmass for each
	x := map[string]int{"sizeX": 10, "sizeY": 100, "coordsX": screenWidth / 3, "coordsY": screenHeight - 10}
	y := map[string]int{"sizeX": 125, "sizeY": 10, "coordsX": int(islandTwoX), "coordsY": int(islandTwoY)}
	z := map[string]int{"sizeX": 200, "sizeY": 10, "coordsX": int(islandOneX), "coordsY": int(islandOneY)}
	log.Println(x, y, z)

	grass.Fill(green)
	landmass.Fill(brown)
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
		cloudX = 0
	}
	oneOpts.GeoM.Translate(cloudX, 5)
	twoOpts.GeoM.Translate(cloudX/2, 100)
	s.DrawImage(one, oneOpts)
	s.DrawImage(two, twoOpts)
	logError(err)
}
