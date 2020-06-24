package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a *Animal) CreateAnimal(foo, loco, noi string){
	a.food = foo
	a.locomotion = loco
	a.noise = noi
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
} 

func main() {
	reader := bufio.NewReader(os.Stdin)

	var cow Animal
	var bird Animal
	var snake Animal
	cow.CreateAnimal("grass", "walk", "moo")
	bird.CreateAnimal("worm", "fly", "peep")
	snake.CreateAnimal("mice", "slither", "hsss")
	fmt.Println("Enter X to exit")
	var name string
	var info string
	for i:= 0; i>=0; i++{
		fmt.Print(">")
		// Input to take only the first two strings in the line
		re,_,err := reader.ReadLine()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
		}
		r1 := string(re)
		r := strings.NewReader(r1)
		fmt.Fscanln(r, &name, &info)
		if (name == "X" || info == "X") {
			break
		}
		// Check name and info
		if (name == "cow"){
			if (info == "eat") {
				cow.Eat()
			} else if (info == "move") {
				cow.Move()
			} else if (info == "speak") {
				cow.Speak()
			} else {
				fmt.Println("Information does not exist")
			}
		} else if (name == "bird"){
			if (info == "eat") {
				bird.Eat()
			} else if (info == "move") {
				bird.Move()
			} else if (info == "speak") {
				bird.Speak()
			} else {
				fmt.Println("Information does not exist")
			}
		} else if (name == "snake"){
			if (info == "eat") {
				snake.Eat()
			} else if (info == "move") {
				snake.Move()
			} else if (info == "speak") {
				snake.Speak()
			} else {
				fmt.Println("Information does not exist")
			}
		} else {
			fmt.Println("Animal does not exist")
		}
	}
}