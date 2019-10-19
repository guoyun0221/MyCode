package main

func hit_enemy() {
	for i := 0; i < len(enemies); i++ { //traversal enemies
		for j := 0; j < len(bullets); j++ { //traversal bullets
			//test if hit
			if bullets[j].x+float64(bullet_size) > enemies[i].x && bullets[j].x < enemies[i].x+float64(enemy_size) {
				if bullets[j].y+float64(bullet_size) > enemies[i].y && bullets[j].y < enemies[i].y+float64(enemy_size) {
					//delete hit enemy and the bullet,score++
					enemies = append(enemies[:i], enemies[i+1:]...)
					bullets = append(bullets[:j], bullets[j+1:]...)
					score++
					break //break is to jump put of this inner loop, if I don't do this,
					//len(enemies) didn't update, in line 12, may i=len(enemies), index out of range
				}
			}
		}
	}
}

func player_died() bool {
	for _, e := range enemies { //traversal enemies
		//if enemy overlap player
		if player.x < e.x+enemy_size && player.x+player_size > e.x {
			if player.y < e.y+enemy_size && player.y+player_size > e.y {
				return true
			}
		}
	}
	return false
}
