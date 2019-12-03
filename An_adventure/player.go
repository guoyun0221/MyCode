package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Player struct {
	//attributes
	Level          int
	EXP            int
	Max_EXP        int
	ATK            int
	MAX_HP, MAX_MP int
	HP, MP         int //current HP and MP
	Money          int
	Step           float64 //about walking, jumping and spell speed(pix per frame)
	//to draw player pic
	Point                              //location in screen
	Left      bool                     //false: player faces right; true: player faces left
	Pics      map[string]*ebiten.Image //the index of the map is the name of the image
	Pic_index string                   //the index of the player pic map
	//status
	Jumping   bool
	Attacking bool
	//weapon and spell
	weapon Weapon
	spells []Spell //recently released spells
}

type Weapon struct {
	Point
	Pics      map[string]*ebiten.Image
	Pic_index string
}

func (player *Player) Init() {
	//initial attributes
	player.Level = 1
	player.Max_EXP = 100
	player.ATK = 1
	player.Step = 4
	player.MAX_HP = 100
	player.HP = 100
	player.MAX_MP = 100
	player.MP = 100
	player.Money = 10
	//set initial location, located at left(1/3), upon ground
	player.X = (screen_width - player_width) / 3
	player.Y = screen_height - player_height - ground_height
	//make map
	player.Pics = make(map[string]*ebiten.Image)
	//initial pic is stand
	player.Pic_index = "player_stand.png"
	//init weapon
	player.weapon.X = player.X + player_width
	player.weapon.Y = player.Y
	player.weapon.Pics = make(map[string]*ebiten.Image)
	player.weapon.Pic_index = "weapon_standby.png"
	// HP and MP recovery
	go player.Recovery()
}

func (player *Player) Get_Movement() {
	/*get player input and do the movement*/
	//walk
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyD) {
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			player.Walk(ebiten.KeyA)
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			player.Walk(ebiten.KeyD)
		}
	}
	//jump
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyK) {
		go player.Jump() //use another goroutine to jump
	}
	//attack, update weapon info
	player.Attack()
	//cast spell thing
	if inpututil.IsKeyJustPressed(ebiten.KeyU) {
		player.Cast_Spell_U()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		player.Cast_Spell_I()
	}
	player.Update_spell()
	//if level up
	if player.EXP >= player.Max_EXP {
		player.Upgrade()
	}
}

func (player *Player) Walk(key ebiten.Key) {
	const border_distance = 3
	//update loaction
	if key == ebiten.KeyA {
		player.Left = true
		process.distance -= player.Step
		player.X -= player.Step
		if player.X < player_width*border_distance {
			//player should be at least 3 person'width away from the left border
			player.X = player_width * border_distance
			//when player gets the border, plus others'X, instead minus player'X,
			//like player is still moving
			background.Player_Moving()
			monsters.Player_Moving()
			coins.Player_Moving()
			shop.Player_Moving()
		}
	} else if key == ebiten.KeyD {
		player.Left = false
		process.distance += player.Step
		player.X += player.Step
		if player.X > screen_width-player_width*(border_distance+1) {
			//right border too(*4 cuz drawing point is left-up point )
			player.X = screen_width - player_width*(border_distance+1)
			//right border, same way to deal it
			background.Player_Moving()
			monsters.Player_Moving()
			coins.Player_Moving()
			shop.Player_Moving()
		}
	}
	//change leg pic
	if process.frame_cnt%3 == 0 { //every 3 frame, change leg
		if player.Pic_index == "player_walk1.png" || player.Pic_index == "player_walk1_L.png" {
			if player.Left {
				player.Pic_index = "player_walk2_L.png"
			} else {
				player.Pic_index = "player_walk2.png"
			}
		} else if player.Pic_index == "player_walk2.png" || player.Pic_index == "player_walk2_L.png" {
			if player.Left {
				player.Pic_index = "player_walk1_L.png"
			} else {
				player.Pic_index = "player_walk1.png"
			}
		} else { //player was not walking just now
			//pick a random walk pic
			if rand.Intn(2) == 1 {
				if player.Left {
					player.Pic_index = "player_walk1_L.png"
				} else {
					player.Pic_index = "player_walk1.png"
				}
			} else {
				if player.Left {
					player.Pic_index = "player_walk2_L.png"
				} else {
					player.Pic_index = "player_walk2.png"
				}
			}
		}
	}
	//if player eats coins
	coins = coins.Player_get()
}

func (player *Player) Jump() {
	const player_jump_height = 200 //player highest jump height

	if !player.Jumping {
		player.Jumping = true //lock, refuse more jump signal
		//jumping pic
		if player.Left {
			player.Pic_index = "player_jumping_L.png"
		} else {
			player.Pic_index = "player_jumping.png"
		}
		//jumping
		t := (player_jump_height / player.Step) / 2
		for i := 0; i < int(t); i++ { //up
			player.Y -= player.Step * 2
			time.Sleep(time.Second / 60)
		}
		for i := 0; i < int(t); i++ { //down
			player.Y += player.Step * 2
			time.Sleep(time.Second / 60)
		}

		player.Y = screen_height - player_height - ground_height //make sure player back to ground
		player.Jumping = false                                   //unlock
	}
}

func (player *Player) Attack() {
	//pick player and weapon pic
	if ebiten.IsKeyPressed(ebiten.KeyJ) { //player is attacking
		player.Attacking = true
		if player.Left {
			player.weapon.Pic_index = "weapon_attack_L.png"
			player.Pic_index = "player_stand_L.png"
		} else {
			player.weapon.Pic_index = "weapon_attack.png"
			player.Pic_index = "player_stand.png"
		}
		//if hit monster
		for i := 0; i < len(monsters); i++ {
			if overlap(player.weapon.Pics[player.weapon.Pic_index], monsters[i].Pic[monsters[i].Pic_index], player.weapon.Point, monsters[i].Point) {
				monsters[i].HP -= player.ATK
				//knockback
				if player.Left {
					monsters[i].X -= (monsters[i].Step + 1)
				} else {
					monsters[i].X += (monsters[i].Step + 1)
				}
			}
		}
	} else { //player is not attacking
		player.Attacking = false
		//draw weapon standby pic
		if player.Left {
			player.weapon.Pic_index = "weapon_standby_L.png"
		} else {
			player.weapon.Pic_index = "weapon_standby.png"
		}
	}
	//update weapon location
	if player.Left {
		player.weapon.X = player.X - weapon_size
		player.weapon.Y = player.Y
	} else {
		player.weapon.X = player.X + player_width
		player.weapon.Y = player.Y
	}
}

func (player *Player) Cast_Spell_U() {
	const ATK_rate_U = 3
	const MP_cost = 10
	/* A horizontal forward ball */
	left := player.Left //Mark player's current direction
	//player casting spell pic
	if left {
		player.Pic_index = "player_spell_L.png"
	} else {
		player.Pic_index = "player_spell.png"
	}
	//add new spell to slice
	if player.MP >= MP_cost {
		player.MP -= MP_cost

		var new_spell Spell
		new_spell.Name = "spell_U"
		new_spell.ATK_rate = ATK_rate_U
		new_spell.Pic = get_img(new_spell.Name + ".png")
		new_spell.Y = player.Y + 54 //by my experiment, this height looks well
		if left {
			new_spell.X = player.X - spell_U_size
			new_spell.direction = "left"
		} else {
			new_spell.X = player.X + player_width
			new_spell.direction = "right"
		}
		player.spells = append(player.spells, new_spell)
	}
}

func (player *Player) Cast_Spell_I() {
	/*a stuff goes straight down from the sky*/
	const ATK_rate_I = 10
	const MP_cost = 20
	left := player.Left
	if left {
		player.Pic_index = "player_spell_L.png"
	} else {
		player.Pic_index = "player_spell.png"
	}
	if player.MP >= MP_cost {
		player.MP -= MP_cost

		var new_spell Spell
		new_spell.ATK_rate = ATK_rate_I
		new_spell.Name = "spell_I"
		new_spell.Pic = get_img(new_spell.Name + ".png")
		new_spell.Y = -spell_I_height
		new_spell.direction = "down"
		//spell is two body width away from player
		if left {
			new_spell.X = player.X - 3*player_width
		} else {
			new_spell.X = player.X + 2*player_width
		}
		player.spells = append(player.spells, new_spell)
	}
}

func (player *Player) Update_spell() {
	var spell_step float64 = player.Step + 3

	for i := 0; i < len(player.spells); i++ {
		//update location
		if player.spells[i].direction == "left" {
			player.spells[i].X -= spell_step
		} else if player.spells[i].direction == "right" {
			player.spells[i].X += spell_step
		} else if player.spells[i].direction == "down" {
			player.spells[i].Y += spell_step
		}
		//delete spell out of bounds
		if player.spells[i].X < (-spell_U_size) || player.spells[i].X > (screen_width) {
			player.spells = append(player.spells[:i], player.spells[i+1:]...)
			//to avoid index out of range
			i--
			break
		}
		if player.spells[i].Y > screen_height {
			player.spells = append(player.spells[:i], player.spells[i+1:]...)
			break
		}
		//if hit monster
		for j := 0; j < len(monsters); j++ {
			if overlap(player.spells[i].Pic, monsters[j].Pic[monsters[j].Pic_index], player.spells[i].Point, monsters[j].Point) {
				monsters[j].HP -= player.spells[i].ATK_rate * player.ATK
			}
		}
	}
}

func (player *Player) Upgrade() {
	player.Level++
	player.MAX_HP += 100
	player.MAX_MP += 100
	player.ATK = player.Level
	player.EXP = 0
}

func (player *Player) Recovery() {
	for player.HP > 0 {
		if player.HP += player.MAX_HP / 100; player.HP > player.MAX_HP {
			player.HP = player.MAX_HP
		}
		if player.MP += player.MAX_MP / 100; player.MP > player.MAX_MP {
			player.MP = player.MAX_MP
		}
		time.Sleep(time.Second)
	}
}
