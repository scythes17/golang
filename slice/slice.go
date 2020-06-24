package main

import (
	"fmt"
	"strconv"
)

func main(){
	// sli:= make([]int, 3) // results in a slice with 3 0 initialized elements
	sli := make([]int, 0, 3)
	var ele string
	fmt.Println("TO EXIT PRESS X")
	for i:= 0; i>=0; i++{
		fmt.Printf("Enter an integer to be inserted: ")
		fmt.Scan(&ele)
		if(ele == "X"){
			break
		}
		// if you enter anything other than an integer, it converts it and stores that
		num,_ := strconv.Atoi(ele)
		// uncomment the following if you want to check slice with first 3 ele initialized to 0
		/* if(i<3){
			sli[0] = num
		} else{ */ 
		sli = append(sli, num)
		fmt.Println("length of slice is: ", len(sli))
		fmt.Println("capacity of slice is: ", cap(sli))
		// }
		fmt.Println("Slice after sorting is: ")
		for j:=0; j<len(sli); j++{
			for k:=0; k<len(sli)-j-1; k++{
				if(sli[k]>sli[k+1]){
					temp := sli[k]
					sli[k] = sli[k+1]
					sli[k+1] = temp
				}
			}
		}
		for j:=0; j<len(sli); j++{
			fmt.Print(" ",sli[j])
		}
		fmt.Println()
	}
}