package main

import "fmt"

func main(){
	var num float32
	var num2 int32
	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&num)
	fmt.Printf("Number after trunc operation: ")
	num2 = int32(num)
	fmt.Printf("%d \n", num2)
}