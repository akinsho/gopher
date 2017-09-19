package main

import (
	"github.com/hajimehoshi/ebiten"
	colorful "github.com/lucasb-eyer/go-colorful"
)

var contact float64

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

	contact = detectCollision()
	posY += velocityY
	if posY >= contact {
		posY = contact
		velocityY = 0.0
	}
	return nil
}
