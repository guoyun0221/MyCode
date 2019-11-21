package main

const (
	//game window size//also background size
	screen_width  = 1200
	screen_height = 600
	//ground height and width. its width equals with screen width
	ground_width  = 1200
	ground_height = 100
	//player pic size
	player_width  = 108
	player_height = 150
	//weapon pic size.put weapon in the same height with player
	weapon_size = 150
	//player highest jump height
	player_jump_height = 200
	//spell pic
	spell_U_size   = 36
	spell_I_width  = 200
	spell_I_height = 100
)

var (
	process    Process
	background Background
	player     Player
)
