package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func Init() {
	rand.Seed(time.Now().Unix())

	player.Init()
	background.Init()

	load_pics()
}

func load_pics() {
	files, err := ioutil.ReadDir("resources/pics") //get a slice of files info
	handle_error(err)
	for _, pic := range files {
		img := get_img(pic.Name()) //get the img
		//put the img to right place
		if strings.Index(pic.Name(), "player") != -1 { //it's a player pic
			player.Pics[pic.Name()] = img //add it to player pics map
		} else if strings.Index(pic.Name(), "back") != -1 {
			background.back_pic = img
		} else if strings.Index(pic.Name(), "ground") != -1 {
			background.ground_pic = img
		} else if strings.Index(pic.Name(), "weapon") != -1 {
			player.weapon.Pics[pic.Name()] = img
		} else if strings.Index(pic.Name(), "hit") != -1 {
			hit_pics.Pic = img
		}
	}
}
