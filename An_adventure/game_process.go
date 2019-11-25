package main

type Process struct {
	//game process related
	distance  float64 //how far has player gone. it influences stage
	stage     int     //it's about game difficulty, monsters level and something like that
	frame_cnt int     //timer, to help to create some duration. it plus one every frame
	//background pics
	monster_group_cnt int
	pause             bool
}

func (process *Process) Update() {
	const stage_duration = 3000
	process.frame_cnt++
	if process.frame_cnt > 2000000000 {
		process.frame_cnt = 0 //avoid overflow
	}
	process.stage = int(process.distance / stage_duration)
}
