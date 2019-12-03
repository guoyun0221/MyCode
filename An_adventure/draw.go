package main

import (
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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
	//draw monsters
	for _, monster := range monsters {
		if monster.X > (-monster_size) && monster.X < screen_width { //draw monsters in the screen scope only
			op.GeoM.Translate(monster.X, monster.Y)
			screen.DrawImage(monster.Pic[monster.Pic_index], &op)
			op.GeoM.Reset()
		}
	}
	//draw monster spells(if attacking)
	for _, monster := range monsters {
		if monster.X > (-monster_size) && monster.X < screen_width {
			if monster.Attacking {
				op.GeoM.Translate(monster.Spell.X, monster.Spell.Y)
				screen.DrawImage(monster.Spell.Pic, &op)
				op.GeoM.Reset()
			}
		}
	}
	//draw coins
	for _, coin := range coins {
		op.GeoM.Translate(coin.X, coin.Y)
		screen.DrawImage(coin.img, &op)
		op.GeoM.Reset()
	}
	//draw shop
	op.GeoM.Translate(shop.X, shop.Y)
	screen.DrawImage(shop.img, &op)
	op.GeoM.Reset()
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
	//draw hit pics
	for i := 0; i < len(hit_pics.Points); i++ {
		op.GeoM.Translate(hit_pics.Points[i].X, hit_pics.Points[i].Y)
		screen.DrawImage(hit_pics.Pic, &op)
		op.GeoM.Reset()
	}
	//delete hit pics points
	hit_pics.Points = hit_pics.Points[0:0]
	//print player state
	ebitenutil.DebugPrint(screen, "lv: "+strconv.Itoa(player.Level)+"   money: "+strconv.Itoa(player.Money)+"   ATK: "+strconv.Itoa(player.ATK))
	ebitenutil.DebugPrint(screen, "\nHP: "+strconv.Itoa(player.HP)+"/"+strconv.Itoa(player.MAX_HP)+"   MP "+strconv.Itoa(player.MP)+"/"+strconv.Itoa(player.MAX_MP))
}
