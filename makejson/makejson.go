package main

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)

func main(){
	var name string
	var addr string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a name: ")
	line1,_,err := reader.ReadLine()
	name = string(line1)
	fmt.Printf("Enter an address: ")
	line2,_,err := reader.ReadLine()
	addr = string(line2)
	var map1 map[string]string
	map1 = make(map[string]string)
	map1["name"] = name
	map1["address"] = addr
	// map1 := map[string]string{"name":name, "address": addr} 
	barr,err := json.Marshal(map1)
	fmt.Println("Json Object: ", string(barr))
	fmt.Println(err)
}