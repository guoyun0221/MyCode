package main

import (
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
)

const (
	screen_size      = 700
	player_size      = 100
	enemy_size       = 70
	bullet_size      = 30
	player_step      = 7
	bullet_step      = 10
	enemy_birth_rate = 50
	enemy_growth     = 5
)

type object struct {
	name string        //background or player or enemy or bullet
	x, y float64       //position, use float64 cuz ebiten.DrawImageOptions.GeoM.Translate() use it
	img  *ebiten.Image //to draw this obj
}

//some global variables
var score int
var background object
var player object
var enemies []object
var bullets []object
var endgame bool

func (obj *object) init(name string, x, y float64) {
	obj.name = name
	obj.x = x
	obj.y = y
	obj.get_img()
}

func (obj *object) get_img() {
	src_file, err := os.Open("pics/" + obj.name + ".png") //read file
	handle_error(err)
	defer src_file.Close()
	//decode file to image.Image
	img, err := png.Decode(src_file)
	handle_error(err)
	// turn image.Image into *ebiten.Image
	obj.img, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	//this function always return nil error, so I Ignore it
}
