package main

import (
	"fmt"
	"sync"
    "time"
)
type ChopS struct {
	sync.Mutex
}
type Philos struct {
	num, count int
	leftcs,rightcs *ChopS
}

func (p Philos) eat(c chan *Philos, wg *sync.WaitGroup) {
	for i:=0;i<3;i++ {
			c <- &p
			if (p.count < 3) {
			p.leftcs.Lock()
			p.rightcs.Lock()

			fmt.Println("starting to eat ",p.num)
			p.count = p.count + 1
			fmt.Println("finishing eating",p.num)
			p.rightcs.Unlock()
			p.leftcs.Unlock()
			wg.Done()		
			}

	}
}

func host(c chan *Philos, wg *sync.WaitGroup) {
	for {
		if  (len(c)==2) {
		<- c
		<- c
		//time delay
		time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	var i int
	var wg sync.WaitGroup
	c := make(chan *Philos,2)

	wg.Add(15)

	ChopSticks := make([] *ChopS,5)
	for i=0;i<5;i++ {
		ChopSticks[i] = new(ChopS)
	}

	Philosophers := make([] *Philos,5)
	for i=0;i<5;i++ {
		Philosophers[i] = &Philos{i+1,0,ChopSticks[i],ChopSticks[(i+1)%5]}
	}

	go host(c,&wg)
	for i=0;i<5;i++ {
		go Philosophers[i].eat(c,&wg)
	}
	wg.Wait()
}