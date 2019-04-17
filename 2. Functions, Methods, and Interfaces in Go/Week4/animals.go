package main

import (
	"fmt"
)
type Animal interface {
	Eat()
	Move()
	Speak()
}
type Animaltp struct {
	food, locomotion, sound string
}

func (v Animaltp) Eat() {
	fmt.Println(v.food)
}

func (v Animaltp) Move() {
	fmt.Println(v.locomotion)
}

func (v Animaltp) Speak() {
	fmt.Println(v.sound)
}

func main() {
	m := map[string]Animaltp{
		"cow" : Animaltp{"grass","walk","moo"},
		"bird" : Animaltp{"worms","fly","peep"},
		"snake" : Animaltp{"mice","slither","hsss"},
	}
	for{
		fmt.Print(">")
		cm:="0"
		an:="0"
		ac:="0"
		fmt.Scan(&cm,&an,&ac)
		if cm=="newanimal"{
			m[an]=Animaltp{m[ac].food, m[ac].locomotion, m[ac].sound}
			fmt.Println("Created it!")
		}else if cm=="query" {
			if ac=="eat"{
				m[an].Eat()
			} else if ac=="move"{
				m[an].Move()
			} else if ac=="speak"{
				m[an].Speak()
			}
		}
	}
}