package main

import "fmt"

func main(){
	var a float64
	var v0 float64
	var s0 float64
	var t float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&v0)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s0)
	fmt.Print("Enter time: ")
	fmt.Scan(&t)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println(fn(t))
}

func GenDisplaceFn(a,v0,s0 float64) func (float64) float64 {
	fn := func (t float64) float64 {
		var s float64
		s = ((a*t*t)/2) +(v0*t) + s0
		return s
	}
	return fn
}