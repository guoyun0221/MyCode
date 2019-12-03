package main

import (
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Coin struct {
	Point
	value int
	img   *ebiten.Image
}

type Coins []Coin

type Shop struct {
	Created int
	Point
	img        *ebiten.Image
	cost       int
	HP_growth  int
	ATK_growth int
}

func (coins Coins) Player_Moving() {
	for i, _ := range coins {
		if player.Left {
			coins[i].X += player.Step
		} else {
			coins[i].X -= player.Step
		}
	}
}

func (coins Coins) Player_get() Coins {
	for i := 0; i < len(coins); i++ {
		if overlap(player.Pics[player.Pic_index], coins[i].img, player.Point, coins[i].Point) {
			player.Money += coins[i].value
			coins = append(coins[:i], coins[i+1:]...)
		}
	}
	return coins
}

func (shop *Shop) Update() {
	shop.Create()
	shop.Buying()
}

func (shop *Shop) Create() {
	const create_duration = 2000
	if int(process.distance/create_duration) == shop.Created {
		shop.Created++
		//show shop in screen
		shop.Y = screen_height - player_height - ground_height - 100
		shop.X = screen_width
		shop.img = get_img("shop.png")
		//get random cost and growth
		shop.cost = rand.Intn(process.stage * 17)
		shop.HP_growth = rand.Intn(process.stage * 50)
		shop.ATK_growth = rand.Intn(process.stage * 2)

		ebitenutil.DebugPrint(shop.img, "   Money-"+strconv.Itoa(shop.cost))
		ebitenutil.DebugPrint(shop.img, "\n   ATK+"+strconv.Itoa(shop.ATK_growth))
		ebitenutil.DebugPrint(shop.img, "\n\n   MAX_HP+"+strconv.Itoa(shop.HP_growth))
	}
}

func (shop *Shop) Buying() {
	if overlap(player.Pics[player.Pic_index], shop.img, player.Point, shop.Point) {
		if player.Money >= shop.cost {
			player.Money -= shop.cost
			player.MAX_HP += shop.HP_growth
			player.ATK += shop.ATK_growth
			//delete shop then
			shop.Y = -100
		}
	}
}

func (shop *Shop) Player_Moving() {
	if player.Left {
		shop.X += player.Step
	} else {
		shop.X -= player.Step
	}
}
