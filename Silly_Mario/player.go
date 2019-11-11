package main

import (
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Player struct {
	score      int
	x, y       float64 //location
	jumping    int     //0 or 1
	jump_speed float64 //relate to roadblock move speed
	imgs       [player_imgs]*ebiten.Image
	img_index  int // which pic will be drew
	blood      int //blood = 0 means player died, game ends
}

func (player *Player) init() {
	player.x = (screen_width - player_width) / 2
	player.y = screen_height - player_height
	player.blood = 1

	for i, _ := range player.imgs { //load every pic
		player.imgs[i] = get_img("player_" + strconv.Itoa(i))
	}
}

func (player *Player) change_pic() {
	//running, change leg
	if player.score%2 == 0 {
		player.img_index = 0
	} else {
		player.img_index = 1
	}
}

func (player *Player) if_jump() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		go player.jump()
	}
}

func (player *Player) jump() {
	if player.jumping == 0 {
		player.jumping = 1 //lock, refuse more signal

		player.jump_speed = rb.speed * speed_ratio
		t := jump_height / player.jump_speed

		for i := 0; i < int(t); i++ {
			player.y -= player.jump_speed //up
			time.Sleep(time.Second / 60)  //60: cuz the game is 60 frames per sec
		}
		for i := 0; i < int(t); i++ {
			player.y += player.jump_speed //down
			time.Sleep(time.Second / 60)
		}

		player.jumping = 0 //unlock
	}
}
