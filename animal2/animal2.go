package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal interface {
	eat()
	move()
	speak()
}

//cow
type Cow struct {
	name string
}
func (c Cow) eat() {
	fmt.Println("grass")
}
func (c Cow) move() {
	fmt.Println("walk")
}
func (c Cow) speak() {
	fmt.Println("moo")
}

//bird
type Bird struct {
	name string
}
func (b Bird) eat() {
	fmt.Println("worm")
}
func (b Bird) move() {
	fmt.Println("fly")
}
func (b Bird) speak() {
	fmt.Println("peep")
}

//snake
type Snake struct {
	name string
}
func (s Snake) eat() {
	fmt.Println("mice")
}
func (s Snake) move() {
	fmt.Println("slither")
}
func (s Snake) speak() {
	fmt.Println("hsss")
}

// Function to process query
func Query(a Animal, query string) {
	if (query == "eat") {
		a.eat()
	} else if (query == "move") {
		a.move()
	} else if (query == "speak") {
		a.speak()
	} else {
		fmt.Println("Please enter valid query")
	}
}

//Function to create an animal
func CreateAnimal(t, name string, cows []Cow, birds []Bird, snakes []Snake) (bool, []Cow, []Bird, []Snake) {
	var f int32
	if (t == "cow") {
		var cow Cow
		cow.name = name
		cows = append(cows, cow)
		fmt.Println("Created it!")
		f=1
	} else if (t == "bird") {
		var bird Bird
		bird.name = name
		birds = append(birds, bird)
		fmt.Println("Created it!")
		f=1
	} else if (t == "snake") {
		var snake Snake
		snake.name = name
		snakes = append(snakes, snake)
		fmt.Println("Created it!")
		f=1
	} else {
		fmt.Println("Type of animal does not exist")
		f=0
	}
	if (f == 1){
		return true, cows, birds, snakes
	} else {
		return false, cows, birds, snakes
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// list of animals
	cows := make([]Cow, 0, 0)
	birds := make([]Bird, 0, 0)
	snakes := make([]Snake, 0, 0)
	names := make([]string, 0, 0)

	fmt.Println("Type 'exit' to exit")
	var command string
	var name string
	var x string
	var c int32
	var f bool

	//Infinite loop
	for i:= 0; i>=0; i++{
		fmt.Print(">")
		// Input to take only the first three strings in the line
		re,_,err := reader.ReadLine()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
		}
		r1 := string(re)
		r := strings.NewReader(r1)
		fmt.Fscanln(r, &command, &name, &x)
		//newanimal command
		if (command == "exit"){
			break
		}
		if (command == "newanimal"){
			c = 0
			for j:=0; j<len(names); j++{
				if (names[j] == name){
					fmt.Println("An animal with that name already exists")
					c = 1
				}
			}
			if (c == 0){
				f,cows,birds,snakes = CreateAnimal(x, name, cows, birds, snakes)
				if (f == true){
					names = append(names, name)
				}
			}
		//query command
		} else if (command == "query") {
			c=0;
			//checking if the animal is a cow
			for j:=0; j<len(cows); j++{
				if (cows[j].name == name){
					Query(cows[j], x)
					c = 1
					break
				}
			}
			if (c == 1){
				continue
			}
			//checking if the animal is a bird
			for j:=0; j<len(birds); j++{
				if (birds[j].name == name){
					Query(birds[j], x)
					c = 1
					break
				}
			}
			if (c == 1){
				continue
			}
			//checking if the animal is a snake
			for j:=0; j<len(snakes); j++{
				if (snakes[j].name == name){
					Query(snakes[j], x)
					c = 1
					break
				}
			}
			if (c == 0) {
				fmt.Println("No animal exists with the given name ", name)
			}
		} else {
			fmt.Println("Please enter a valid command!")
		}
	}
}


/* Sample newanimal command: newanimal tan cow
	Output: Created it!
   Sample query command: query tan speak
	Output: moo */