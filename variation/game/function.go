package game

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func status(army Army, money int) { /*show status : money and army*/

	for i := 0; i < len(army); { //I've seen once somebody's population is 0 but still
		if army[i].Number <= 0 { //show on list, I don't know how to fix this bug before
			army = append(army[:i], army[i+1:]...) //show status, so I write this check code
		} else {
			i++
		}
	}

	time.Sleep(200 * time.Millisecond)
	fmt.Println("--------status--------")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Your money: ", money)
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Your soldiers:")
	for _, s := range army {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("name:%s  population:%d  rank:%s  ATK:%d  HP:%d  lv:%d\n", s.Name, s.Number, s.Rank, s.ATK, s.HP, s.LV)
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Println("-----------------------")
}

func Start() (Army, int) {
	fmt.Println("----------------------------------Game started----------------------------------")

	time.Sleep(200 * time.Millisecond)
	fmt.Println("Strengthen your army and defeat monsters")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("You have three ways to do it:")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Producing: Spend 10 to produce a soldier")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Upgrading: Spend 10 to get a soldier ATK and HP increased by 10 percent")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Merging: Spend 10 to merge two different soldiers into a new and stronger one")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("-------------------------------------------------------------------------------")

	army, money := initialize()
	time.Sleep(200 * time.Millisecond)
	fmt.Println("You get two new kinds of soildiers")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("The first one is")
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("population:%d  rank:%s  ATK:%d  HP:%d  lv:%d\n", army[0].Number, army[0].Rank, army[0].ATK, army[0].HP, army[0].LV)
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Name it:")
	fmt.Scanln(&army[0].Name)
	time.Sleep(200 * time.Millisecond)
	fmt.Println("The second one is")
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("population:%d  rank:%s  ATK:%d  HP:%d  lv:%d\n", army[1].Number, army[1].Rank, army[1].ATK, army[1].HP, army[1].LV)
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Name it:")
	fmt.Scanln(&army[1].Name)

	return army, money
}

func initialize() (Army, int) { /*create two soldiers first, give some money*/

	rand.Seed(time.Now().Unix())

	var army Army
	for i := 0; i < 2; i++ {
		var s Soldier
		s.born()
		s.Number = 3
		army = append(army, s)
	}

	var money int = 100

	return army, money
}

func Option(army Army, money *int) Army {
	for {
		status(army, *money)

		time.Sleep(200 * time.Millisecond)
		fmt.Println("What do you wannna do?")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("A: Produce   B: Upgrade   C: Merge   D: Done")
		var choice string
		fmt.Scanln(&choice)

		if choice == "A" || choice == "a" {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("Input the name and number of the soldiers you wanna produce. Separate with space key")
			var name string
			var num int
			fmt.Scanln(&name, &num)
			if num*10 > *money {
				time.Sleep(200 * time.Millisecond)
				fmt.Println("You don't have enough money")
			} else {
				for i, _ := range army { //here,I use i,not the latter one. bucause the latter
					if army[i].Name == name { //one is a copy is army[i], can't to change army[i]
						army[i].produce(num)
						*money = *money - num*10
					}
				}
			}

		} else if choice == "B" || choice == "b" {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("Input the name of soldier you wannna upgrade")
			var name string
			fmt.Scanln(&name)
			if 10 > *money {
				time.Sleep(200 * time.Millisecond)
				fmt.Println("You don't have enough money")
			} else {
				for i, _ := range army {
					if army[i].Name == name {
						army[i].upgrade()
						*money = *money - 10
					}
				}
			}

		} else if choice == "C" || choice == "c" {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("Input the name of two soldiers you want to merge.Separate with space key")
			var s1, s2 string
			fmt.Scanln(&s1, &s2)
			if *money < 10 {
				time.Sleep(200 * time.Millisecond)
				fmt.Println("You don't have enough money.")
			} else {
				var j, k int
				j, k = -1, -1
				for i, _ := range army {
					if army[i].Name == s1 {
						j = i
					}
					if army[i].Name == s2 {
						k = i
					}
				}
				if j != -1 && k != -1 && j != k {
					army = army.merge(j, k)
					*money = *money - 10
					time.Sleep(200 * time.Millisecond)
					fmt.Println("You get a new soldier:")
					time.Sleep(200 * time.Millisecond)
					fmt.Printf("rank:%s  ATK:%d   HP:%d\n", army[len(army)-1].Rank, army[len(army)-1].ATK, army[len(army)-1].HP)
					time.Sleep(200 * time.Millisecond)
					fmt.Println("Name it:")
					var name string
					fmt.Scanln(&name)
					army[len(army)-1].Name = name
				}
				if j == k {
					fmt.Println("I said two fucking different soldiers")
				}
			}

		} else {
			break
		}
	}
	return army
}

func War(army Army, money *int, stage int) (Army, bool) {
	mon := CreateMonster(stage) //create monsters
	mon_info(mon)               //separate funcs which have print, in case I write gui in the future
	war_start(stage)
	army, mons := war_process(army, mon) //I can't just put mon.number in the left
	mon.number = mons
	war_over()
	if mon.number > 0 {
		lose(mon.number)
		return army, true
	} else {
		win(money, stage)
		return army, false
	}
}

func CreateMonster(stage int) Monster {
	var mon Monster

	if stage < 15 {
		mon.number = stage*2 + rand.Intn(5) - 1 //these rand is to make data not that Changeless
		mon.ATK = stage*5 + rand.Intn(9) - 4
		mon.HP = stage*50 + rand.Intn(91) - 45
		mon.LV = mon.ATK/5 + mon.HP/100
	} else { //after 15 round, monster becomes stronger fast
		t := float64(stage)
		n := int(math.Exp(t))
		mon.ATK = n + rand.Intn(10) - 5
		mon.HP = n*10 + rand.Intn(100) - 50
		mon.LV = mon.ATK/5 + mon.HP/100
		mon.number = stage*2 + rand.Intn(5) - 1
	}

	return mon
}

func mon_info(mon Monster) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("----------Monsters information----------")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("")
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("ATK:%d   HP:%d   lv:%d   number:%d\n", mon.ATK, mon.HP, mon.LV, mon.number)
}

func war_start(stage int) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("----------------War started----------------")
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("                  stage:%d\n", stage)
}
func war_over() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("--------------The war is over--------------")
}

func war_process(army Army, mon Monster) (Army, int) { //return the number of monster afer war, 0 means win

	var a Army //separate the members in group
	var m []Monster
	for _, s := range army { //different type soldier
		for i := 0; i < s.Number; i++ { //same type soldier, translate its number
			a = append(a, s)
			a[len(a)-1].Number = 1
		}
	}
	for i := 0; i < mon.number; i++ {
		m = append(m, mon)
		m[len(m)-1].number = 1
	}

	for { //endless loop until one side cleared
		for i := 0; i < mix(len(a), len(m)); i++ { //all the people attack a random enemy
			m[rand.Intn(len(m))].HP -= a[i].ATK
			a[rand.Intn(len(a))].HP -= m[i].ATK
		}
		if mix(len(a), len(m)) == len(a) {
			for i := len(a); i < len(m); i++ {
				a[rand.Intn(len(a))].HP -= m[i].ATK
			}
		} else {
			for i := len(m); i < len(a); i++ {
				m[rand.Intn(len(m))].HP -= a[i].ATK
			}

		}
		//after one round attack, delete died person
		for i := 0; i < len(a); { //using for_range_loop will lead panic: index out of range
			if a[i].HP <= 0 {
				soldier_died(a[i]) //print died soldier
				a = append(a[:i], a[i+1:]...)
			} else {
				i++
			}
		}
		for i := 0; i < len(m); {
			if m[i].HP <= 0 {
				monster_died() //print died monster
				m = append(m[:i], m[i+1:]...)
			} else {
				i++
			}
		}

		if len(a) == 0 || len(m) == 0 {
			break
		}
	}
	//reset the number of army and monster
	for i, _ := range army {
		army[i].Number = 0
		for _, s := range a {
			if s.Name == army[i].Name {
				army[i].Number++
			}
		}
	}
	for i := 0; i < len(army); { //delete 0 number soldier
		if army[i].Number <= 0 {
			army = append(army[:i], army[i+1:]...)
		} else {
			i++
		}
	}

	mon.number = len(m)

	return army, mon.number
}

func mix(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func soldier_died(s Soldier) {
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("one %s died\n", s.Name)
}

func monster_died() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("One monster has been killed")
}

func lose(num int) {
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("There are %d monster(s) left and you lose all your army\n", num)
}

func win(money *int, stage int) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("All monsters have been killed")
	num := stage*5 + rand.Intn(10) - 5
	*money += num
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("You won %d money\n", num)
}

func End() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("----------------------------------The game is over----------------------------------")
}
