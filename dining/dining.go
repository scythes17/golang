package main

/* IMPORTANT : host continues running until main completes execution*/

import (
	"fmt"
	"sync"
)

type ChopS struct{sync.Mutex}
type Philo struct{
	leftCS, rightCS *ChopS
	num int
}
var w sync.WaitGroup

func (p Philo) eat(h1 chan int, h2 chan int){
	for i:=0; i<3; i++ {
		y:= <-h1 // taking value from host to check how many people are eating
		y++ 
		h2 <- y // sending value to host to show how many people will be eating if eat is allowed to continue
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("Starting to eat  ",p.num)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		fmt.Println("Finishing eating ",p.num)
		h2 <- y // sending value to host to show that eating by this philosopher is finished
	}
	w.Done()
}

func host( h1 chan int, h2 chan int){
	fmt.Println("Host is running")
	x := 0 // first time value for eat function saying no one is eating
	for{
		if(x<2){ // if it is less than or equal to 2
			h1 <-x // giving the number of philosophers currently eating
			x = <-h2 // taking value from eat to update how many philosophers are eating right now
		} else {
			<-h2 // waiting for one eat function to finish executing
			x-- // a philosopher finished eating so reducing the no of philosophers eating at the moment
		}
	}
}

func main() {
	h1:= make (chan int, 2)
	h2:= make (chan int, 2)

	go host(h1 , h2)
	CSticks := make([]*ChopS, 5)
	for i:=0; i<5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo,5)
	for i:=0; i<5; i++{
		philos[i] = &Philo{ CSticks[i], CSticks[(i+1)%5], i+1 }
	}

	for i:=0; i<5; i++{
		w.Add(1)
		go philos[i].eat(h1, h2)
	}

	w.Wait()
}