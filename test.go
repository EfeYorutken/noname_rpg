package main

import("fmt")

func print_enemies(list []character){
	for i := 0; i < len(list); i++{
		if list[i].hp > 0{
			fmt.Printf("%v\tHP:%v ATK:%v DEF:%v\n",list[i].name,list[i].hp,list[i].atk,list[i].def)
		}
	}
}

func main(){
	player := naked_char("efe",50,50,50,50,50,50)
	var arr []character

	id_q := "who do you want to attack?"
	bp_q := "which part of the body?"
	h_q := "which hand?"

	for i := 0; i < 10; i++{
		new_one := naked_char(fmt.Sprintf("joe # %v",i),i*10,i*10,i*10,i*10,i*10,i*10)
		new_one.attack_targets = []character{player}
		arr = append(arr,new_one)
	}


	id  := 0
	bp  := 0
	h  := 0

	fmt.Println(id_q)
	fmt.Scan(&id)
	fmt.Println(bp_q)
	fmt.Scan(&bp)
	fmt.Println(h_q)
	fmt.Scan(&h)

	print_enemies(arr)
	// func attack(attacker character, target *character, hand int,target_part int) bool{
	fmt.Println(attack(player,&arr[id],h,bp))
	fmt.Println("-----")
	print_enemies(arr)

}
