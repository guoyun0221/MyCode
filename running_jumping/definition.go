package main

import "github.com/hajimehoshi/ebiten"

const (
	screen_width     = 1000
	screen_height    = 500
	player_width     = 135
	player_height    = 165
	roadblock_width  = 90
	roadblock_height = 60
	init_speed       = 7   //roadblock move speed at the start of the game
	player_imgs      = 3   //total imgs of player
	duration         = 3   //every (duration) frames, score++
	jump_height      = 250 //highest point where player could jump
	safe_width       = 35  //for the clear area of player img
	speed_growth     = 1   //new roadblock speed grows 1 pix
	speed_ratio      = 1.2 // = player_jump_speed/roadblock_move_speed
)

var background *ebiten.Image = get_img("background")
var player Player
var rb Roadblocks
var distance int
var timer int
