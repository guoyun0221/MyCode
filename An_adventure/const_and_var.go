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
	//spell pic
	spell_U_size   = 36
	spell_I_width  = 200
	spell_I_height = 100
	//monster
	monster_size       = 120
	mosnter_spell_size = 60
	//hit pic size
	hit_pic_size = 20
)

var (
	process    Process
	background Background
	player     Player
	monsters   Monsters
	hit_pics   Hit_Pics
)
