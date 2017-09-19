package main

import (
	"github.com/hajimehoshi/ebiten"
)

func jump() {
	if onGround && posY == contact {
		velocityY = -7.0 //this value determines the height of the jump to be executed
		onGround = false
	}
	// time.AfterFunc(300*time.Millisecond, endJump)
}

func endJump() {
	if velocityY < -3.0 {
		velocityY = -3.0
	}
}

func handleInput() {
	// If space is pressed execute a jump
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		jump()
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if !onGround {
			posY += 2
		}
	}
	// When the "down arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if onGround {
			posY -= 2
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
		if posX >= 0 || posX <= screenWidth {
			posX += 4
		} else {
			posX = 0
		}
	}
}
