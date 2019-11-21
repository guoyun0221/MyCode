package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Process struct {
	//game process related
	distance  float64 //how far has player gone. it influences stage
	stage     int     //it's about game difficulty, monsters level and something like that
	frame_cnt int     //timer, to help to create some duration. it plus one every frame
	//background pics
}

type Background struct {
	//background pics
	back_pic     *ebiten.Image
	ground_pic   *ebiten.Image
	ground_point []Point //there should almost always two ground pics in the screen
}

func (background *Background) Init() {
	var ground Point
	ground.X = 0
	ground.Y = screen_height - ground_height
	background.ground_point = append(background.ground_point, ground)
}

func (background *Background) Player_Moving() {
	if player.Left {
		//create a new ground to adjoin the old one in the left, like they are whole
		if len(background.ground_point) == 1 {
			//there are 2 ground pic in screen, if only one ground there, create a new one
			var ground Point
			ground.X = background.ground_point[0].X - ground_width
			ground.Y = background.ground_point[0].Y
			background.ground_point = append(background.ground_point, ground)
		}
		//player is walking to left, increase ground'X
		for i, _ := range background.ground_point {
			background.ground_point[i].X += player.Speed
		}
	} else {
		if len(background.ground_point) == 1 {
			var ground Point
			ground.X = background.ground_point[0].X + ground_width
			ground.Y = background.ground_point[0].Y
			background.ground_point = append(background.ground_point, ground)
		}
		for i, _ := range background.ground_point {
			background.ground_point[i].X -= player.Speed
		}
	}
	//delete ground pic out of bounds
	for i := 0; i < len(background.ground_point); i++ {
		if background.ground_point[i].X < (-ground_width) || background.ground_point[i].X > ground_width {
			background.ground_point = append(background.ground_point[:i], background.ground_point[i+1:]...)
		}
	}
}
