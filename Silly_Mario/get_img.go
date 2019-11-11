package main

import (
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
)

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
