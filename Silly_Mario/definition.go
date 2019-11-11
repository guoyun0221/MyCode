package main

import "github.com/hajimehoshi/ebiten"

const (
	screen_width     = 1000
	screen_height    = 500
	player_width     = 135
	player_height    = 165
	roadblock_width  = 90
	roadblock_height = 60
	blood_size       = 50
	init_speed       = 7   //roadblock move speed at the start of the game
	player_imgs      = 2   //total imgs of player
	duration         = 3   //every (duration) frames, score++
	jump_height      = 250 //highest point where player could jump
	safe_width       = 35  //for the clear area of player img
	speed_growth     = 1   //new roadblock speed grows 1 pix
	speed_ratio      = 1.2 // = player_jump_speed/roadblock_move_speed
	blood_chance     = 3   //there is 1/3 chance that new blood appear
)

var background *ebiten.Image = get_img("background")
var blood_img *ebiten.Image = get_img("blood")
var player Player
var rb Roadblocks
var blood []Blood
var distance int
var timer int
