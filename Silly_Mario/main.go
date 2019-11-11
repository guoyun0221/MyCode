package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	//init stuff
	rand.Seed(time.Now().Unix())
	player.init()
	rb.init()
	//run the game
	err := ebiten.Run(game, screen_width, screen_height, 1, "Silly Mario")
	handle_error(err)
}

func game(screen *ebiten.Image) error {
	if player.blood >= 0 { //game not end

		//new roadblock? new blood? roadblocks and blood move left
		create()
		move()
		//player jump? player location, player pic
		player.if_jump()
		player.change_pic()
		//gain blood?
		gain_blood()
		//if crash
		crash()
		//every (duration)frames, score++
		timer++
		if timer%duration == 0 {
			player.score++
		}

	}
	//draw stuff
	draw(screen)

	return nil
}
