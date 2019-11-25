package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

type Monster struct {
	Name string
	//attributes
	Level      int
	ATK        int
	MAX_HP, HP int
	Step       float64 //pix per frame
	// to draw pic
	Point
	Self_Point Point //Take birth point as origin of coordinate system
	Left       bool
	Pic        map[string]*ebiten.Image
	Pic_index  string
	//status
	Attacking bool
	//attack thing
	Spell
}

type Monsters []Monster

func (monsters Monsters) Update() Monsters {
	monsters = monsters.Birth()
	monsters.Walk()
	monsters.If_Attack()
	monsters.Update_Spell()
	return monsters
}

func (monsters Monsters) Birth() Monsters {
	const distance_interval = 1800
	//Every (distance_interval) pixes player walks, a group of monsters will be born
	const birth_zone = 1200 //monster will birth in the zone of [x,x+BirthZone), x is the screen_width
	const max_number = 7    //number of monsters at most in a group

	if int(process.distance/distance_interval) == process.monster_group_cnt {
		//create a group of mosnters
		num := rand.Intn(max_number) + 1
		for i := 0; i < num; i++ {
			//init a monster
			var monster Monster
			monster.Name = "monster_0"
			monster.Step = 2
			monster.X = float64(rand.Intn(birth_zone) + screen_width)
			monster.Y = screen_height - ground_height - monster_size
			monster.Left = (rand.Intn(2) == 1) //randon direction
			monster.Pic = make(map[string]*ebiten.Image)
			monster.Pic[monster.Name+"_L.png"] = get_img(monster.Name + "_L.png")
			monster.Pic[monster.Name+".png"] = get_img(monster.Name + ".png")
			if monster.Left {
				monster.Pic_index = monster.Name + "_L.png"
			} else {
				monster.Pic_index = monster.Name + ".png"
			}
			monster.ATK = (process.stage + 1)
			monster.MAX_HP = (process.stage + 1) * 100
			monster.HP = monster.MAX_HP
			//add it to slice
			monsters = append(monsters, monster)
		}
		process.monster_group_cnt++
	}
	return monsters
}

func (monsters Monsters) Walk() {
	const Max_Distance = 300 //to make monster active zone
	for i := 0; i < len(monsters); i++ {
		if monsters[i].Left {
			monsters[i].X -= monsters[i].Step
			monsters[i].Self_Point.X -= monsters[i].Step
			monsters[i].Pic_index = monsters[i].Name + "_L.png"
			if monsters[i].Self_Point.X < (-Max_Distance) {
				//change direction
				monsters[i].Left = false
			}
		} else {
			monsters[i].X += monsters[i].Step
			monsters[i].Self_Point.X += monsters[i].Step
			monsters[i].Pic_index = monsters[i].Name + ".png"
			if monsters[i].Self_Point.X > Max_Distance {
				monsters[i].Left = true
			}
		}
	}
}

func (monsters Monsters) If_Attack() {
	const attack_chance = 500 // 1/500 chance to attack
	for i, _ := range monsters {
		if rand.Intn(attack_chance) == 0 && !monsters[i].Attacking {
			monsters[i].Attacking = true
			monsters[i].Spell.Pic = get_img("monster_spell.png")
			monsters[i].Spell.mark = process.frame_cnt
		}
	}
}

func (monsters Monsters) Update_Spell() {
	const spell_duration = 60 //(frames)
	const spell_step = 1
	for i, _ := range monsters {
		if monsters[i].Attacking {
			if process.frame_cnt < monsters[i].Spell.mark+spell_duration {
				if monsters[i].Left {
					monsters[i].Spell.X = monsters[i].X - mosnter_spell_size - spell_step*(float64(process.frame_cnt-monsters[i].Spell.mark))
					monsters[i].Spell.Y = monsters[i].Y + (monster_size-mosnter_spell_size)/2
				} else {
					monsters[i].Spell.X = monsters[i].X + monster_size + spell_step*(float64(process.frame_cnt-monsters[i].Spell.mark))
					monsters[i].Spell.Y = monsters[i].Y + (monster_size-mosnter_spell_size)/2
				}
				if overlap(monsters[i].Spell.Pic, player.Pics[player.Pic_index], monsters[i].Spell.Point, player.Point) {
					player.HP -= monsters[i].ATK
				}
			} else { //spell time is over
				monsters[i].Attacking = false

			}
		}
	}
}

func (monsters Monsters) Player_Moving() {
	for i, _ := range monsters {
		if player.Left {
			monsters[i].X += player.Step
		} else {
			monsters[i].X -= player.Step
		}
	}
}
