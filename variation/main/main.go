package main

import "variation/game"

func main() {
	army, money := game.Start()
	var stage int = 1
	var ending bool = false
	for {
		army = game.Option(army, &money)             //in merge method, new army was created
		army, ending = game.War(army, &money, stage) //army is never the old one,
		//It will lead bug if I don't return army,I don't really know what happened
		if ending == true {
			break
		}
		stage++
	}
	game.End()
}
