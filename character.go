package main

import ("fmt"
	"math/rand")

func print_items(of character) string{
	fmt.Println("ITEMS\n-----")
	res := ""
	for i := 0; i < len(of.inventory); i++{
		res += fmt.Sprintf("%v\t%v : hp bounsu = %v attack bonus = %v defence bonus = %v toughness bonus = %v power bonus = %v\n",i,of.inventory[i].name,of.inventory[i].hp_bonus,of.inventory[i].atk_bonus,of.inventory[i].def_bonus,of.inventory[i].t_bonus,of.inventory[i].power_bonus)
	}
	return res
}


func print_subject(subject character) string{
	return fmt.Sprintf("%v\tHP:%v ATK:%v DEF:%v",subject.name,subject.hp,subject.atk,subject.def)
}


func print_enemies(of character) string{
	res := ""
	for i := 0; i < len(of.attack_targets); i++{
		if of.attack_targets[i].hp > 0{
			res += fmt.Sprintf("%v\t%v\n",i,print_subject(of.attack_targets[i]))
		}
	}
	return res
}

type character struct{
	name string
	apparel [4]apparel
	weapon [2]weapon
	hp int
	atk int
	def int
	power int
	toughness int
	stamina int
	inventory[10]consumable
	attack_targets []character
}

func get_weakest_point(target character) int{

	min := target.apparel[0].def_bonus + target.apparel[0].t_bonus

	for i := 1; i < len(target.apparel); i++{
		if (target.apparel[i].def_bonus + target.apparel[i].t_bonus) < min{
			min = target.apparel[i].def_bonus + target.apparel[i].t_bonus
		}
	}
	return min
}

func naked_char(name string, hp int, atk int, def int, power int, toughness int, stamina int)character{
	return character{

		name : name,
		apparel : [4]apparel{emp_apparel,emp_apparel,emp_apparel,emp_apparel}, //to be changed
		weapon : [2]weapon{emp_weapon,emp_weapon}, //to be changed
		hp : hp,
		atk : atk,
		def : def,
		power : power,
		toughness : toughness,
		stamina : stamina,

	}
}

//works
func attack(attacker character, target character, hand int,target_part int) int{
	target_sum_def := target.apparel[target_part].def_bonus + target.def
	atacker_sum_atk := attacker.weapon[hand].atk_bonus + attacker.atk
	if target_sum_def < atacker_sum_atk{
		target_sum_t := target.apparel[target_part].t_bonus + target.toughness
		atacker_sum_p := attacker.weapon[hand].power_bonus + attacker.power
		if atacker_sum_p > target_sum_t{
			return atacker_sum_p - target_sum_t
		}
	}
	return 0
}

func get_if_healer_in_invent(invent [10]consumable) int{
	for i := 0; i < len(invent); i++{
		if invent[i].hp_bonus > 0{
			return i
		}
	}
	return -1
}

//does not work
func  act(subject *character) int{
	dmg := 0
	if subject.hp < 50 && get_if_healer_in_invent(subject.inventory) > -1{
		use_consumeable(subject,subject.inventory[get_if_healer_in_invent(subject.inventory)])
	} else{
		target_index := rand.Int() % len(subject.attack_targets)
		target_part := get_weakest_point(subject.attack_targets[target_index])
		hand := rand.Int() % 2;

		dmg = attack(*subject,subject.attack_targets[target_index],hand,target_part)
		subject.attack_targets[target_index].hp -= dmg
	}
	return dmg
}

func player_action(player *character){
	fmt.Println(print_enemies(*player))
	player_action := ""
	fmt.Println("what do you want to do?(attack/consume)")
	fmt.Scan(&player_action)
	if player_action == "attack" {
		var id_q = "who do you want to attack?"
		var bp_q = "which part of the body?"
		var h_q = "which hand?"
		id  := 0
		bp  := 0
		h  := 0

		fmt.Println(id_q)
		fmt.Scan(&id)
		fmt.Println(bp_q)
		fmt.Scan(&bp)
		fmt.Println(h_q)
		fmt.Scan(&h)
		player.attack_targets[id].hp -= attack(*player,player.attack_targets[id],h,bp)
	} else if player_action == "consume"{
		fmt.Println("ITEMS\n-----")
		fmt.Println(print_items(*player))
		item_index := 0
		fmt.Println("which one do you want to use")
		fmt.Scan(&item_index)
		use_consumeable(player,player.inventory[item_index])
	} else{
		fmt.Println("in your confusion you do nothing...")
	}
}
