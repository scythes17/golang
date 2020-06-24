package main

import (
	"fmt"
	"os"
	"strings"
)
type name struct {
	fname string
	lname string
}

func main(){
	var p1 name
	var filename string
	sli := make([]name, 0)
	fmt.Print("Enter the name of the text file: ")
	fmt.Scan(&filename)
	f,err:= os.Open(filename)
	// checking if file exists
	if(err != nil){
		fmt.Println("Enter correct filename!")
		fmt.Println(err)
	}
	// getting the size of the file
	f1,err:= f.Stat()
	if(err != nil){
		fmt.Println(err)
	}
	size := f1.Size()
	// reading the file
	barr:= make([]byte,size)
	nb,err:= f.Read(barr)

	if(err != nil || nb == 0){
		fmt.Println(err)
	}

	// splitting read file into lines
	line:= strings.Split(string(barr), "\n")
	// splitting lines into fname and lname
	for i:= range line{
		if(line[i] == ""){
			fmt.Println("Line is empty")
			break
		}
		name:= strings.Split(line[i], " ")
		
		// reducing the length of first and last name to 20 chars
		if len(name[0]) >= 20 {
			name1 := name[0]
			name[0] = name1[0:20]
		}
		if len(name[1]) >= 20 {
			name2 := name[1]
			name[1] = name2[0:20]
		}

		p1.fname = name[0]
		p1.lname = name[1]
		sli = append(sli, p1)
	}
	f.Close()
	// prints the firstname lastname pair
	for i:= range sli{
		fmt.Println(sli[i])
	}
}