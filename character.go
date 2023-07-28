package main


var WHO_TO_ATTACK string = "Which enemies do you want to attack? "
var WHICH_HAND_TO_USE string = "Which hand do you want to use? "
var WHICH_PART_OF_THE_BODY string = "Which part of the body do you want to attack? "

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

func attack(attacker character, target *character, hand int,target_part int) bool{
	target_sum_def := target.apparel[target_part].def_bonus + target.def
	atacker_sum_atk := attacker.weapon[hand].atk_bonus + attacker.atk
	if target_sum_def < atacker_sum_atk{
		target_sum_t := target.apparel[target_part].t_bonus + target.toughness
		atacker_sum_p := attacker.weapon[hand].power_bonus + attacker.power
		if atacker_sum_p > target_sum_t{
			target.hp = target.hp - atacker_sum_p - target_sum_t
			return true
		}
	}
	return false
}
