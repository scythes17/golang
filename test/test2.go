//Program to solve the 5 philosopher and chopstick issue
package main

import (
	"fmt"
	"sync"
)

//type Communication int

const (
	isFree     int = 0
	notFree        = 1
	yesFree        = 2
	doneEating     = 3
	IAmFull        = 4
)

type Message struct {
	communication int
	index         int
}

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	sendChannel     *chan Message
	receiveChannel  *chan Message
	leftCS, rightCS *ChopS
	index           int
}

//eat function
func (p Philo) eat() {
	for i := 0; i < 3; {
		*p.sendChannel <- Message{
			communication: isFree,
			index:         p.index,
		}
		per := <-*p.receiveChannel
		if per.communication == yesFree && per.index == p.index {
			p.leftCS.Lock()
			p.rightCS.Lock()
			fmt.Printf("starting to eat %d\n", p.index)
			fmt.Printf("finishing eating %d\n", p.index)
			p.rightCS.Unlock()
			p.leftCS.Unlock()
			i++
			*p.sendChannel <- Message{
				communication: doneEating,
				index:         p.index,
			}
		} else {
			continue
		}
	}
	*p.sendChannel <- Message{
		communication: IAmFull,
		index:         p.index,
	}
}

type Host struct {
	sendChannel    *chan Message
	receiveChannel *chan Message
}

// Host to control dinning
func (h Host) host(wg *sync.WaitGroup) {
	defer wg.Done()
	done := 0
	eating := make(map[int]bool)
	for done < 5 {
		message := <-*h.receiveChannel
		leftIndex := getLeftIndex(message.index)
		rightIndex := getRightIndex(message.index)
		switch com := message.communication; com {
		case isFree:
			if len(eating) >= 2 {
				*h.sendChannel <- Message{
					notFree,
					message.index,
				}
			} else if !eating[leftIndex] && !eating[rightIndex] {
				eating[message.index] = true
				*h.sendChannel <- Message{
					communication: yesFree,
					index:         message.index,
				}
			} else {
				*h.sendChannel <- Message{
					communication: notFree,
					index:         message.index,
				}
			}
		case doneEating:
			delete(eating, message.index)
		case IAmFull:
			done += 1
		}
	}
}

func getRightIndex(index int) int {
	return (index + 1) % 5
}

func getLeftIndex(index int) int {
	leftIndex := index - 1
	if leftIndex == 0 {
		leftIndex = 5
	}
	return leftIndex
}

// Main function
func main() {
	philoToHost := make(chan Message)
	hostToPhilo := make(chan Message)
	cSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		cSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{
			sendChannel:    &philoToHost,
			receiveChannel: &hostToPhilo,
			leftCS:         cSticks[i],
			rightCS:        cSticks[(i+1)%5],
			index:          i + 1,
		}
	}
	host := Host{
		sendChannel:    &hostToPhilo,
		receiveChannel: &philoToHost,
	}
	for _, philo := range philos {
		go philo.eat()
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go host.host(&wg)
	wg.Wait()
}