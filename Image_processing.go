package main

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println("Something you can do")
	fmt.Println("A: Create a random rect on a image")

	var s string
	fmt.Scanln(&s)

	if s == "A" || s == "a" {
		random_color_rect()
	}

}

func random_color_rect() {
	fmt.Println("Input the name of the original image (only accept .png file)") //read a img file
	var file_name string
	fmt.Scanln(&file_name)
	src_file, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	defer src_file.Close()
	img, err := png.Decode(src_file) //turn the img file into data
	if err != nil {
		fmt.Println(err)
	}

	color_pic := image.NewRGBA(image.Rect(0, 0, 20+rand.Intn(img.Bounds().Dx()), 20+rand.Intn(img.Bounds().Dy())))
	//color size is random and smaller than src img, +20 is to avoid the pic being too small to see
	color := palette.Plan9[rand.Intn(256)]
	//random color
	draw.Draw(color_pic, color_pic.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)
	//fill color//image.ZP is Point{0,0}
	img_draw := image.NewRGBA(img.Bounds())
	draw.Draw(img_draw, img_draw.Bounds(), img, image.ZP, draw.Src)
	//turn image.Image into draw.Image, cuz draw.Draw()'s fist argument has to be draw.Image
	x0 := rand.Intn(img_draw.Bounds().Dx())
	y0 := rand.Intn(img_draw.Bounds().Dy())
	draw.Draw(img_draw, image.Rect(x0, y0, x0+color_pic.Bounds().Dx(), y0+color_pic.Bounds().Dy()), color_pic, image.ZP, draw.Src)
	//in a random place, draw the random color rect

	fmt.Println("Input the name of the new image (only accept .png file)") //create the new file
	fmt.Scanln(&file_name)
	dst_file, err := os.Create(file_name)
	if err != nil {
		fmt.Println(err)
	}
	defer dst_file.Close()
	err = png.Encode(dst_file, img_draw) //turn data into image file
	if err != nil {
		fmt.Println(err)
	}
}
