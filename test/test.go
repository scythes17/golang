package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
)

type Animal interface {
  Eat()
  Move()
  Speak()
}

type Cow struct {
  food string
  locomotion string
  sound string
}

func (c Cow) Eat() {
  fmt.Println(c.food)
}

func (c Cow) Move() {
  fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
  fmt.Println(c.sound)
}

type Bird struct {
  food string
  locomotion string
  sound string
}

func (c Bird) Eat() {
  fmt.Println(c.food)
}

func (c Bird) Move() {
  fmt.Println(c.locomotion)
}

func (c Bird) Speak() {
  fmt.Println(c.sound)
}

type Snake struct {
  food string
  locomotion string
  sound string
}

func (c Snake) Eat() {
  fmt.Println(c.food)
}

func (c Snake) Move() {
  fmt.Println(c.locomotion)
}

func (c Snake) Speak() {
  fmt.Println(c.sound)
}

type DomesticatedAnimal struct {
  name string
  animal Animal
}

var cow Cow
var bird Bird
var snake Snake
var myAnimals []DomesticatedAnimal

func main(){
 setupData()
 for {
    fmt.Println(">")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()
    stringss := strings.Fields(input)
    command := stringss[0]
    if command == "newanimal" {
      name := stringss[1]
      animalType := stringss[2]
      NewAnimal(name, animalType)
    } 
    if command == "query" {
      animal := stringss[1]
      action := stringss[2]
      Query(animal, action)
    }
 }
}

func setupData(){
  cow = newCow()
  bird = newBird()
  snake = newSnake()
  myAnimals = make([]DomesticatedAnimal, 0)
}

func newCow() Cow {
	return Cow{"grass", "walk", "moo"}
}

func newBird() Bird {
	return Bird{"worms", "fly", "peep"}
}

func newSnake() Snake {
	return Snake{"mice", "slither", "hss"}
}

func NewAnimal(name string, animalType string){
  var newAnimal Animal
  switch animalType {
    case "cow":
        newAnimal = newCow()
    case "snake":
        newAnimal = newSnake()
    case "bird":
        newAnimal = newBird()
  } 

  myAnimals = append(myAnimals, DomesticatedAnimal{name, newAnimal})
  fmt.Println("Created it!")
}

func Query(animal string, action string){
  var choosenAnimal Animal

  switch animal {
	case "cow":
		choosenAnimal = cow
	case "snake":
		choosenAnimal = snake
	case "bird":
		choosenAnimal = bird
  } 
  
  if action == "eat" {
    choosenAnimal.Eat()
  }
  if action == "move" {
    choosenAnimal.Move()
  }
  if action == "speak" {
    choosenAnimal.Speak()
  }
}