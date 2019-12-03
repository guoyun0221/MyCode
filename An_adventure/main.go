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
	if player.HP > 0 {
		//update process thing
		process.Update()
		//update player
		player.Get_Movement()
		//update monsters
		monsters = monsters.Update()
		//shop thing
		shop.Update()
	}
	//draw stuff
	Draw(screen)
	return nil
}
