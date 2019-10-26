package main

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
)

const (
	canvas_width  = 1300
	canvas_height = 700
	palette_width = 100
	border_width  = 10
	color_height  = 85 // the height of every color in palette area
)

var canvas *ebiten.Image
var palette [8]color.RGBA //there are 8 kinds of color
var pen [8]*ebiten.Image  //8 images of 8 color pens
var pen_color int         //index of pen to mark current pen color

func main() {
	canvas_init()
	pen_init()

	err := ebiten.Run(drawing, canvas_width, canvas_height, 1, "weak painter")
	if err != nil {
		fmt.Println(err)
	}
}

func drawing(screen *ebiten.Image) error {
	var op ebiten.DrawImageOptions
	//draw canvas
	err := screen.DrawImage(canvas, &op)
	if err != nil {
		fmt.Println(err)
	}
	//when mouse pressed
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x < palette_width { //choose color
			pen_color = (y - border_width) / color_height //mark down color that being chosen
		} else if x > palette_width+border_width { //drawing
			op.GeoM.Translate(float64(x)-5, float64(y)-5) //to make cursor in the middle
			canvas.DrawImage(pen[pen_color], &op)         //draw.
		}
	}

	return nil
}

func canvas_init() {
	//get canvas background and color it white
	background := image.NewRGBA(image.Rect(0, 0, canvas_width, canvas_height))
	for x := 0; x < canvas_width; x++ {
		for y := 0; y < canvas_height; y++ {
			background.Set(x, y, color.White)
		}
	}
	//draw border
	for x := 0; x < canvas_width; x++ {
		for y := 0; y < border_width; y++ {
			background.Set(x, y, color.Black)
		}
		for y := canvas_height - border_width; y < canvas_height; y++ {
			background.Set(x, y, color.Black)
		}
	}
	for y := 0; y < canvas_height; y++ {
		for x := palette_width; x < palette_width+border_width; x++ {
			background.Set(x, y, color.Black)
		}
		for x := canvas_width - border_width; x < canvas_width; x++ {
			background.Set(x, y, color.Black)
		}
	}
	//draw palette
	palette_init()
	for i, _ := range palette {
		for x := 0; x < palette_width; x++ {
			for y := border_width + i*color_height; y < border_width+(i+1)*color_height; y++ {
				background.Set(x, y, palette[i])
			}
		}
	}

	canvas, _ = ebiten.NewImageFromImage(background, ebiten.FilterDefault)
}

func palette_init() {
	palette[0] = color.RGBA{0, 0, 0, 255}       //black
	palette[1] = color.RGBA{255, 0, 0, 255}     //red
	palette[2] = color.RGBA{0, 255, 0, 255}     //green
	palette[3] = color.RGBA{0, 0, 255, 255}     //blue
	palette[4] = color.RGBA{255, 255, 0, 255}   //yellow
	palette[5] = color.RGBA{255, 0, 255, 255}   //violet
	palette[6] = color.RGBA{0, 255, 255, 255}   //cyan
	palette[7] = color.RGBA{255, 255, 255, 255} //white,could be used as eraser
}

func pen_init() {
	const pen_radius = 5
	//pen is a img that with the circular color
	for i, _ := range pen {
		pen_img := image.NewRGBA(image.Rect(0, 0, pen_radius*2, pen_radius*2))
		for x := 0; x < pen_img.Bounds().Dx(); x++ {
			for y := 0; y < pen_img.Bounds().Dy(); y++ {
				if math.Sqrt(float64((x-pen_radius)*(x-pen_radius))+float64((y-pen_radius)*(y-pen_radius))) < pen_radius {
					pen_img.Set(x, y, palette[i]) //set the color with palette
				}
			}
		}
		pen[i], _ = ebiten.NewImageFromImage(pen_img, ebiten.FilterDefault)
	}
}
