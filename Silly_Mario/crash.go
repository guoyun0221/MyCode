package main

func crash() {
	for i := 0; i < len(rb.point); i++ {
		//if player overlaps roadblock//safe width is for the clear area
		if player.x+safe_width < rb.point[i].x+roadblock_width && player.x+player_width-safe_width > rb.point[i].x {
			if player.y < rb.point[i].y+roadblock_height && player.y+player_height > rb.point[i].y {
				player.blood--
			}
		}
	}

}
