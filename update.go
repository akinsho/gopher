package main

import (
	"github.com/hajimehoshi/ebiten"
	colorful "github.com/lucasb-eyer/go-colorful"
)

func update(screen *ebiten.Image) error {

	if sky == nil {
		sky, _ = colorful.Hex("#5FE8F7")
	}
	//The Screen must first be filled otherwise it will cover everything else
	screen.Fill(sky)
	drawCharacter(screen)
	drawLand(screen)
	drawClouds(screen)
	handleInput()
	velocityY += gravity

	posY += velocityY
	if posY >= lowerBound {
		posY = lowerBound
		onGround = true
	}
	return nil
}
