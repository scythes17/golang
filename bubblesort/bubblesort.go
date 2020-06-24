package main

import "fmt"

func main(){
	sli:= []int{}
	var size int
	var ele int
	fmt.Print("Enter the number of elements to be inserted: ")
	fmt.Scan(&size)
	fmt.Print("Enter the array elements: ")
	for i:= 0; i<size; i++{
		fmt.Scan(&ele)
		sli = append(sli, ele)
	}
	BubbleSort(sli)
	for i:=0; i<size; i++{
		fmt.Print(sli[i], " ")
	}
}

func BubbleSort(sli[] int){
	for i:=0; i<len(sli); i++{
		for j:=0; j<len(sli)-i-1; j++{
			if(sli[j]>sli[j+1]){
				Swap(sli, j)
			}
		}
	}
}

func Swap(sli[] int, i int){
	temp:= sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}