package main
import("math/rand"
				"fmt"
)

var scenes []string = []string{"desert","temple","forest"}


func are_enemies_dead(arr []character) bool{
	for i := 0; i < len(arr); i++{
		if arr[i].hp != 0{
			return false
		}
	}
	return true
}

type scene struct{
	enemies []character
	player character
	terrain string
}

func get_new_scene(enemies []character, player character, terrain string) scene{
	return scene{
		enemies : enemies,
		player : player,
		terrain : terrain,
	}
}


func (s *scene) play_out_scene(arr []character) (scene,bool){
	for !are_enemies_dead(s.enemies) && s.player.hp != 0{
		fmt.Println(print_subject(s.player))
		player_action(&s.player)
		for i := 0; i < len(s.enemies); i++{
			if s.enemies[i].hp > 0{
				dmg := act(&s.enemies[i])
				if dmg > 0{
					s.player.hp -= dmg
					fmt.Printf("%v's attack was succesfull\n",s.enemies[i].name)
				}else{
					fmt.Printf("%v's attack was not succesfull\n",s.enemies[i].name)
				}
			}

		}


	}


	return get_new_scene(arr,s.player,scenes[rand.Int() % len(scenes)]),s.player.hp != 0

}
