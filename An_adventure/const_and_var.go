package main

const (
	//some pic size
	screen_width       = 1200
	screen_height      = 600
	ground_width       = 1200
	ground_height      = 100
	player_width       = 108
	player_height      = 150
	weapon_size        = 150 //put weapon in the same height with player
	spell_U_size       = 36
	spell_I_width      = 200
	spell_I_height     = 100
	monster_size       = 120
	mosnter_spell_size = 60
	hit_pic_size       = 20
	coin_size          = 50
)

var (
	process    Process
	background Background
	player     Player
	monsters   Monsters
	hit_pics   Hit_Pics
	coins      Coins
	shop       Shop
)
