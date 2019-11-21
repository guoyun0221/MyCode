package main

import (
	"github.com/hajimehoshi/ebiten"
)

func Draw(screen *ebiten.Image) {
	var op ebiten.DrawImageOptions
	//draw background
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background.back_pic, &op)
	op.GeoM.Reset()
	for _, ground := range background.ground_point {
		op.GeoM.Translate(ground.X, ground.Y)
		screen.DrawImage(background.ground_pic, &op)
		op.GeoM.Reset()
	}
	//draw player
	op.GeoM.Translate(player.X, player.Y)
	screen.DrawImage(player.Pics[player.Pic_index], &op) //draw player pic according index
	op.GeoM.Reset()
	//draw weapon
	if player.Pic_index != "player_spell.png" && player.Pic_index != "player_spell_L.png" {
		//player is not casting spell,
		op.GeoM.Translate(player.weapon.X, player.weapon.Y)
		screen.DrawImage(player.weapon.Pics[player.weapon.Pic_index], &op)
		op.GeoM.Reset()
	}
	//draw spells
	for _, spell := range player.spells {
		op.GeoM.Translate(spell.X, spell.Y)
		screen.DrawImage(spell.Pic, &op)
		op.GeoM.Reset()
	}
}
