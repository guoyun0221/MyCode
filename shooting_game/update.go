package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func update_player() {
	//moving position
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		player.y -= player_step
	}
	if player.y < 0 {
		player.y = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		player.y += player_step
	}
	if player.y > float64(screen_size-player_size) {
		player.y = float64(screen_size - player_size)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.x -= player_step
	}
	if player.x < 0 {
		player.x = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.x += player_step
	}
	if player.x > float64(screen_size-player_size) {
		player.x = float64(screen_size - player_size)
	}

	//shooting
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		var bullet object //create new bullet
		bullet.init("bullet", player.x+(player_size-bullet_size)/2, player.y)
		bullets = append(bullets, bullet)
	}
}

func update_bullets() {
	for i := 0; i < len(bullets); i++ {
		bullets[i].y -= bullet_step //bullet moving

		if bullets[i].y < -bullet_size { //delete bullets over border
			bullets = append(bullets[:i], bullets[i+1:]...)
		}
	}
}

func update_enemies() {
	// create new enemy
	if rand.Intn(enemy_birth_rate) == 0 { // there are 1/(enemy_birth_rate) chance that new enemy born
		var enemy object
		enemy.init("enemy", float64(rand.Intn(screen_size-enemy_size)), float64(-enemy_size))
		enemies = append(enemies, enemy)
	}
	//update position
	for i := 0; i < len(enemies); i++ {
		enemies[i].y += float64(score/enemy_growth) + 1 //higher score you got, faster enemies moving,
		//every (enemy_growth) enemies killed, enemy speed gains 1pixel/frame.
		//+1 is to handle score=0 case
		//delete enemies out of border
		if enemies[i].y > screen_size {
			enemies = append(enemies[:i], enemies[i+1:]...)
		}
	}
}
