package main

import (
	"github.com/hajimehoshi/ebiten"
)

func jump() {
	if onGround {
		velocityY = -12.0
		onGround = false
	}
}

func handleInput() {
	// If space is pressed execute a jump
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		jump()
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		// log.Println("pos Y ", posY)
	}
	// When the "down arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if posY > 280 {
			if onGround {
				posY -= 2
			}
		}
	}
	// When the "left arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if posX >= 0 || posX < screenWidth {
			posX -= 4
		} else {
			posX = 0
		}
	}
	// When the "right arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if posX > 0 || posX < screenWidth {
			posX += 4
		} else {
			posX = 0
		}
	}
}
