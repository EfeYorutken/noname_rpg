package main

import("fmt")


type weapon struct{
	name string
	atk_bonus int
	power_bonus int
}

type apparel struct{
	name string
	def_bonus int
	t_bonus int//toughness bonus
}

type consumable struct{
	name string
	hp_bonus int
	def_bonus int
	t_bonus int
	atk_bonus int
	power_bonus int
}

var emp_cons consumable = consumable{
	name : "empty",
	hp_bonus : 0,
	def_bonus : 0,
	t_bonus : 0,
	atk_bonus : 0,
	power_bonus : 0,
}

var emp_apparel apparel = apparel{
	name : "none",
	def_bonus : 0,
	t_bonus : 0,
}


var emp_weapon weapon = weapon{
	name : "empty",
	atk_bonus : 0,
	power_bonus : 0,
}

func get_consumable(name string, hp_bonus int, def_bonus int, t_bonus int, atk_bonus int, power_bonus int) consumable{
	return consumable{
		name : name,
		hp_bonus : hp_bonus,
		def_bonus : def_bonus,
		t_bonus : t_bonus,
		atk_bonus : atk_bonus,
		power_bonus : power_bonus,
	}
}

func get_armor(n string, db int, tb int, body string) apparel{
	switch body{
	case "head":
		return apparel{ name : n, def_bonus : db, t_bonus : tb, }
	case "chest":
		return apparel{ name : n, def_bonus : db, t_bonus : tb, }
	case "leg":
		return apparel{ name : n, def_bonus : db, t_bonus : tb, }
	case "feet":
		return apparel{ name : n, def_bonus : db, t_bonus : tb, }
	default:
		return emp_apparel
	}
}

func get_weapon(n string, ab int, pb int) weapon{
	return weapon{
		name : n,
		atk_bonus : ab,
		power_bonus : pb,
	}
}

func use_consumeable(target *character, item consumable) string{
	target.hp += item.hp_bonus
	target.atk += item.atk_bonus
	target.def += item.def_bonus
	target.power += item.power_bonus
	target.toughness += item.t_bonus
	return fmt.Sprintf("%v used %v",target.name,item.name)
}
