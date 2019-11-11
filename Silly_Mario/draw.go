package main

import (
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func draw(screen *ebiten.Image) {
	var op ebiten.DrawImageOptions
	//draw background
	screen.DrawImage(background, &op)
	//draw roadblocks and blood
	for i := 0; i < len(rb.point); i++ {
		op.GeoM.Translate(rb.point[i].x, rb.point[i].y)
		screen.DrawImage(rb.img, &op)
		op.GeoM.Reset()
	}

	for i := 0; i < len(blood); i++ {
		op.GeoM.Translate(blood[i].x, blood[i].y)
		screen.DrawImage(blood[i].img, &op)
		op.GeoM.Reset()
	}
	//draw player
	op.GeoM.Translate(player.x, player.y)
	screen.DrawImage(player.imgs[player.img_index], &op)
	op.GeoM.Reset()
	//draw blood cnt
	for i := 0; i < player.blood; i++ {
		op.GeoM.Translate(screen_width-float64((i+1)*blood_size), 0)
		screen.DrawImage(blood_img, &op)
		op.GeoM.Reset()
	}
	//print score
	ebitenutil.DebugPrint(screen, "Score "+strconv.Itoa(player.score))
}
