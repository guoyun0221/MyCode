package main

import (
	"github.com/hajimehoshi/ebiten"
)

func main() {
	Init()
	err := ebiten.Run(Game, screen_width, screen_height, 1, "Game")
	handle_error(err)
}

func Game(screen *ebiten.Image) error {
	player.Get_Movement()
	//draw stuff
	Draw(screen)
	//update timer
	if process.frame_cnt++; process.frame_cnt > 2000000000 {
		process.frame_cnt = 0 //avoid overflow
	}
	return nil
}
