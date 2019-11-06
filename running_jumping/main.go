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
	err := ebiten.Run(game, screen_width, screen_height, 1, "Running_Game")
	handle_error(err)
}

func game(screen *ebiten.Image) error {
	if !player.death {

		//new roadblock? roadblocks move left
		rb.create()
		rb.move()
		//player jump? player location, player pic
		player.if_jump()
		player.change_pic()
		//if crash
		player.death = ifcrash()
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
