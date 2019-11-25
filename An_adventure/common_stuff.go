package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
)

type Point struct {
	X, Y float64
}

type Spell struct {
	Name     string
	ATK_rate int // =  player_akt * atk_rate
	MP_cost  int
	Point
	direction string
	Pic       *ebiten.Image
	mark      int //to mark frame_count
}

type Hit_Pics struct {
	Points []Point
	Pic    *ebiten.Image
}

func handle_error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func get_img(name string) *ebiten.Image {
	file, err := os.Open("resources/pics/" + name) //read file
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

func overlap(a, b *ebiten.Image, pa, pb Point) bool {
	if (pa.X < pb.X+float64(b.Bounds().Dx())) && (pa.X+float64(a.Bounds().Dx()) > pb.X) {
		if (pa.Y < pb.Y+float64(b.Bounds().Dy())) && (pa.Y+float64(a.Bounds().Dy()) > pb.Y) {
			add_hit_pic(a, b, pa, pb)
			return true
		}
	}
	return false
}

func add_hit_pic(a, b *ebiten.Image, pa, pb Point) {
	var p Point
	center_a_x := (pa.X + pa.X + float64(a.Bounds().Dx())) / 2
	center_a_y := (pa.Y + pa.Y + float64(a.Bounds().Dy())) / 2
	center_b_x := (pb.X + pb.X + float64(b.Bounds().Dx())) / 2
	center_b_y := (pb.Y + pb.Y + float64(b.Bounds().Dy())) / 2
	p.X = (center_a_x+center_b_x)/2 - hit_pic_size/2
	p.Y = (center_a_y+center_b_y)/2 - hit_pic_size/2
	hit_pics.Points = append(hit_pics.Points, p)
}
