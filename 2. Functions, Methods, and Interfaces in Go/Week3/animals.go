package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, sound string
}

func (v Animal) Eat() {
	fmt.Println(v.food)
}

func (v Animal) Move() {
	fmt.Println(v.locomotion)
}

func (v Animal) Speak() {
	fmt.Println(v.sound)
}

func main() {
	m := map[string]Animal{
		"cow" : Animal{"grass","walk","moo"},
		"bird" : Animal{"worms","fly","peep"},
		"snake" : Animal{"mice","slither","hsss"},
	}
	for{
		fmt.Print(">")
		an:="0"
		ac:="0"
		fmt.Scan(&an,&ac)
		if ac=="eat"{
			m[an].Eat()
		} else if ac=="move"{
			m[an].Move()
		} else if ac=="speak"{
			m[an].Speak()
		}
	}
}