package main

import("fmt")

func main(){
	player := naked_char("efe",50,50,50,50,50,50)
	var arr []character

	for i := 0; i < 10; i++{
		new_one := naked_char(fmt.Sprintf("joe # %v",i),i*10,i*10,i*10,i*10,i*10,i*10)
		new_one.attack_targets = []character{player}
		arr = append(arr,new_one)
	}

	player.attack_targets = arr


	keep_playing := true
	s := get_new_scene(arr,player,"desert")
	for keep_playing{
		s,keep_playing = s.play_out_scene(arr)
		fmt.Println(keep_playing)
	}



}
