/*Race conditions occur when outcome depends on non-deterministic ordering(interleavings) i.e. 
outcome of one thread depends on the outcome of another thread, if there was no communication between
the two threads the race condition would not occur.
Basically, race condition makes the program unreliable because it would return different outcome
everytime you run the program



In the below program, running it, prints different values everytime, as the loop in the add thread is
running for a significant amount of time and we cannot know at which time will the three print threads 
be executed*/





package main

import (
	"fmt"
	"sync"
)

func main(){
	var wee sync.WaitGroup
	wee.Add(4)
	var x int=1
	go add(&wee, &x)
	go print(&wee, &x)
	go print(&wee, &x)
	go print(&wee, &x)
	wee.Wait()
}

func add(wee *sync.WaitGroup, x *int){
	for i:=0; i<100000; i++{
		*x++
	}
	wee.Done()
}

func print(wee *sync.WaitGroup, x *int){
	fmt.Println(*x)
	wee.Done()
}