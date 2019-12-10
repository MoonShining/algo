package main

import "log"

func full(runes []rune, i int){
	if i == len(runes) -1 {
		log.Println(string(runes))
		return
	}else{
		for j := i;j<len(runes);j++{
			runes[i], runes[j] = runes[j], runes[i]
			full(runes, i+1)
			runes[i], runes[j] = runes[j], runes[i]
		}
	}
}


func main() {
	full([]rune("abc"),0)
}
