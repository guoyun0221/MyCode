package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func main() {
	//some init stuff
	rand.Seed(time.Now().Unix())
	background.init("background", 0, 0)
	player.init("player", (screen_size-player_size)/2, screen_size-player_size)

	//run the game
	err := ebiten.Run(game, screen_size, screen_size, 1, "Shooting_Game")
	//function game() is called at every frame(60 times a second)
	handle_error(err)
}

func game(screen *ebiten.Image) error {

	/* get player's operation and update position of objects*/
	if !endgame { //if game ends, nobody moves anymore
		update_player()
		update_bullets()
		update_enemies()
	}

	/*events*/
	hit_enemy()
	endgame = player_died()

	/*draw objects*/
	var op_back ebiten.DrawImageOptions //draw background
	screen.DrawImage(background.img, &op_back)

	var op_player ebiten.DrawImageOptions //draw player,
	//op.GeoM.Translate(x, y) will plus x,y to the old ones instead use x,y directly,
	//so every time I use this function, I need to create a new op
	op_player.GeoM.Translate(player.x, player.y)
	//draw at player.x, player.y point
	screen.DrawImage(player.img, &op_player)

	for i, _ := range bullets { //draw bullets
		var op_bullet ebiten.DrawImageOptions
		op_bullet.GeoM.Translate(bullets[i].x, bullets[i].y)
		screen.DrawImage(bullets[i].img, &op_bullet)
	}

	for i, _ := range enemies { //draw enemies
		var op_enemy ebiten.DrawImageOptions
		op_enemy.GeoM.Translate(enemies[i].x, enemies[i].y)
		screen.DrawImage(enemies[i].img, &op_enemy)
	}

	ebitenutil.DebugPrint(screen, "Score"+strconv.Itoa(score)) //print score

	return nil
}
