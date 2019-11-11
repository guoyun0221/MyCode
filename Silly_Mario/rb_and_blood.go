package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

type Roadblocks struct {
	point []point // position of every roadblock
	speed float64 //pixs per frame//also is blood move speed
	img   *ebiten.Image
}

type point struct {
	x, y float64
}

type Blood struct {
	x, y float64
	img  *ebiten.Image
}

func (rb *Roadblocks) init() {
	rb.img = get_img("roadblock")
	rb.speed = init_speed
}

func (blood *Blood) init() {
	blood.img = get_img("blood")
}

func create() {
	if distance%screen_width <= int(rb.speed) { //one roadblock gone, new one comes

		var new_rb point
		new_rb.x = screen_width
		new_rb.y = screen_height - roadblock_height
		rb.point = append(rb.point, new_rb)
		rb.speed += speed_growth //faster

		if rand.Intn(blood_chance) == 0 { // if new a blood
			var new_blood Blood //new blood locate above roadblock
			new_blood.x = new_rb.x + (roadblock_width-blood_size)/2
			new_blood.y = screen_height - jump_height - player_height
			new_blood.init()
			blood = append(blood, new_blood)
		}
	}
}

func move() { //roadblock and blood move left
	for i := 0; i < len(rb.point); i++ { //roadblock
		rb.point[i].x -= rb.speed
		if rb.point[i].x < (-roadblock_width) {
			rb.point = append(rb.point[:i], rb.point[i+1:]...)
		}
	}
	for i := 0; i < len(blood); i++ { //blood
		blood[i].x -= rb.speed // sync with rb
		if blood[i].x < (-blood_size) {
			blood = append(blood[:i], blood[i+1:]...)
		}
	}
	distance += int(rb.speed)
}

func gain_blood() {
	for i := 0; i < len(blood); i++ {
		if player.x < blood[i].x+blood_size && player.x+player_width > blood[i].x {
			if player.y < blood[i].y+blood_size && player.y+player_height > blood[i].y {
				player.blood++
				blood = append(blood[:i], blood[i+1:]...)
			}
		}
	}
}

func crash() {
	for i := 0; i < len(rb.point); i++ {
		//if player overlaps roadblock//safe width is for the clear area
		if player.x+safe_width < rb.point[i].x+roadblock_width && player.x+player_width-safe_width > rb.point[i].x {
			if player.y < rb.point[i].y+roadblock_height && player.y+player_height > rb.point[i].y {
				player.blood--
				if player.blood > 0 {
					rb.point = append(rb.point[:i], rb.point[i+1:]...)
				}
			}
		}
	}
}
