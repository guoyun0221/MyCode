package main

import (
	"image/png"
	"os"
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
	death      bool
}

type point struct {
	x, y float64
}

type Roadblocks struct {
	point []point // position of every roadblock//in fact I don't need a slice here,
	//cuz there is always one roadblock in the screen, but I'm too lazy to modify
	speed float64 //pixs per frame
	img   *ebiten.Image
}

func (player *Player) init() {
	player.x = (screen_width - player_width) / 2
	player.y = screen_height - player_height

	for i, _ := range player.imgs { //load every pic
		player.imgs[i] = get_img("player_" + strconv.Itoa(i))
	}
}

func (rb *Roadblocks) init() {
	rb.img = get_img("roadblock")
	rb.speed = init_speed
}

func get_img(name string) *ebiten.Image {

	file, err := os.Open("pics/" + name + ".png") //read file
	handle_error(err)
	defer file.Close()
	//decode file to image.Image
	src_img, err := png.Decode(file)
	handle_error(err)
	// turn image.Image into *ebiten.Image
	img, _ := ebiten.NewImageFromImage(src_img, ebiten.FilterDefault)
	//this function always return nil error, so I Ignore it
	return img
}

func (rb *Roadblocks) create() {
	if distance%screen_width <= int(rb.speed) { //one roadblock gone, new one comes

		var new_rb point
		new_rb.x = screen_width
		new_rb.y = screen_height - roadblock_height
		rb.point = append(rb.point, new_rb)
		rb.speed += speed_growth //faster
	}
}

func (rb *Roadblocks) move() { //roadblock move left
	for i := 0; i < len(rb.point); i++ {
		rb.point[i].x -= rb.speed
		if rb.point[i].x < (-roadblock_width) {
			rb.point = append(rb.point[:i], rb.point[i+1:]...)
		}
	}
	distance += int(rb.speed)
}

func (player *Player) change_pic() {
	if player.jumping == 1 {
		player.img_index = 2 //jump pic
	} else {
		//running, change leg
		if player.score%2 == 0 {
			player.img_index = 0
		} else {
			player.img_index = 1
		}
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
