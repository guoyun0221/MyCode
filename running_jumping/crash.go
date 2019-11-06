package main

func ifcrash() bool {
	for _, p := range rb.point {
		//if player overlaps roadblock//safe width is for the clear area
		if player.x+safe_width < p.x+roadblock_width && player.x+player_width-safe_width > p.x {
			if player.y < p.y+roadblock_height && player.y+player_height > p.y {
				return true
			}
		}
	}
	return false
}
