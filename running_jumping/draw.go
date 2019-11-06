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
	//draw roadblocks
	for i := 0; i < len(rb.point); i++ {
		op.GeoM.Translate(rb.point[i].x, rb.point[i].y)
		screen.DrawImage(rb.img, &op)
		op.GeoM.Reset()
	}
	//draw player
	op.GeoM.Translate(player.x, player.y)
	screen.DrawImage(player.imgs[player.img_index], &op)
	//print score
	ebitenutil.DebugPrint(screen, "Score "+strconv.Itoa(player.score)) 
}
