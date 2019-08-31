package game

import "math/rand"

type Soldier struct {
	Name   string
	Number int //this type soldier's population
	Rank   string
	ATK    int
	HP     int
	LV     int
}

type Monster struct {
	ATK    int
	HP     int
	LV     int
	number int
}

type Army []Soldier

func (s *Soldier) born() {
	var score int = rand.Intn(100)
	if score < 50 {
		s.Rank = "D" //ATK: 5-9      HP: 50-90
		s.ATK = rand.Intn(4) + 5
		s.HP = rand.Intn(40) + 50
	} else if score < 80 {
		s.Rank = "C" //ATK: 10-14    HP: 100-140
		s.ATK = rand.Intn(4) + 10
		s.HP = rand.Intn(40) + 100
	} else if score < 95 {
		s.Rank = "B" //ATK: 20-28    HP: 200-280
		s.ATK = rand.Intn(8) + 20
		s.HP = rand.Intn(80) + 200
	} else {
		s.Rank = "A" //ATK: 50-90    HP: 500-900
		s.ATK = rand.Intn(40) + 50
		s.HP = rand.Intn(400) + 500
	}
	s.Number = 1
	s.LV = 1
}

func (s *Soldier) produce(num int) {
	s.Number += num
}

func (s *Soldier) upgrade() {
	s.LV++
	s.ATK = int(float32(s.ATK) * (1 + 0.1))
	s.HP = int(float32(s.HP) * (1 + 0.1))
}

func (army Army) merge(i, j int) Army {
	army[i].Number--
	army[j].Number--
	var oldi Soldier = army[i] // to save the old data after deleted
	var oldj Soldier = army[j]
	if army[i].Number <= 0 {
		army = append(army[:i], army[i+1:]...)
		if i < j { //j should change
			j--
		}
	}
	if army[j].Number <= 0 {
		army = append(army[:j], army[j+1:]...)
	}

	var s_old Soldier //select the best attribute of the two old soldier

	if oldi.Rank == oldj.Rank {
		s_old.Rank = oldi.Rank
		if oldi.ATK > oldj.ATK {
			s_old.ATK = oldi.ATK
		} else {
			s_old.ATK = oldj.ATK
		}
		if oldi.HP > oldj.HP {
			s_old.HP = oldi.HP
		} else {
			s_old.HP = oldj.HP
		}
	} else if oldi.Rank < oldj.Rank { //rank doesn't equal, select the higher one
		s_old = oldi //"A"<"B" true
	} else {
		s_old = oldj
	}
	var s Soldier //create a random soldier, compare it with the old one
	s.born()
	if s.Rank >= s_old.Rank {
		s = s_old
		s.ATK += rand.Intn(3)
		s.HP += rand.Intn(30)
	}
	s.LV = 1
	s.Number = 1
	army = append(army, s) //start a new array,the army here is not the old
	return army            //one, so I need to return it to cover the old one
}
