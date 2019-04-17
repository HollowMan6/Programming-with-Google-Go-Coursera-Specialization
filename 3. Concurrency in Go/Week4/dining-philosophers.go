package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Philosopher struct {
	name      string
	chopstick chan bool
	neighbor  *Philosopher
	leftHand  bool
	rightHand bool
}

func NewPhilosopher(name string) *Philosopher {
	p := &Philosopher{
		name:      name,
		chopstick: make(chan bool, 1),
		leftHand:  false,
		rightHand: false,
	}
	p.chopstick <- true
	return p
}

func (p *Philosopher) Thinking() {
	time.Sleep(time.Second * 1)
}

func (p *Philosopher) GetChopsticks() {
	for {
		if p.leftHand && p.rightHand {
			return
		}

		select {
		case <-p.chopstick:
			fmt.Printf("%v starting to eat.\n", p.name)
			p.leftHand = true
		case <-p.neighbor.chopstick:
			p.rightHand = true
		case <-time.After(time.Second * 2):
			if p.leftHand {
				p.chopstick <- true
				p.leftHand = false
			}
			if p.rightHand {
				p.neighbor.chopstick <- true
				p.rightHand = false
			}

			// No chopstick for me, try to figuring out the reason
			p.Thinking()
		}
	}
}

func (p *Philosopher) Dining() {
	time.Sleep(time.Second * time.Duration(rand.Int63n(10)))
}

func (p *Philosopher) ReturnChopsticks() {
	p.chopstick <- true
	p.neighbor.chopstick <- true
	p.leftHand, p.rightHand = false, false
}

func (p *Philosopher) Dine(announce chan *Philosopher) {
	p.Thinking()
	p.GetChopsticks()
	p.Dining()
	p.ReturnChopsticks()
	announce <- p
}

func main() {
	// Define a bunch of Philosophers
	philosophers := []*Philosopher{
		NewPhilosopher("1"),
		NewPhilosopher("2"),
		NewPhilosopher("3"),
		NewPhilosopher("4"),
		NewPhilosopher("5"),
	}

	// Initialize Philosophers' neighbors and let them pick their own chopsticks
	i := 0
	for ; i < len(philosophers)-1; i++ {
		philosophers[i].neighbor = philosophers[i+1]
	}
	philosophers[i].neighbor = philosophers[0]

	// Let Philosophers start to dine
	announce := make(chan *Philosopher)
	defer close(announce)
	for _, p := range philosophers {
		go p.Dine(announce)
	}

	// Loop until all Philosophers finish dining
	count := 0
	for {
		if count == len(philosophers) {
			break
		}
		select {
		case p := <-announce:
			fmt.Printf("finishing eating %v \n", p.name)
			count += 1
		}
	}
}