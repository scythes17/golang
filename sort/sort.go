 
package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"os"
)

func main() {
	var size int
	var ele int
	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&size)
	if(size<4){
		fmt.Println("Size should be more than 4")
		os.Exit(1)
	}
	fmt.Print("Enter the array elements seperated by enter: ")
	arr := make([]int,0,size)
	for i:= 0; i<size; i++{
		fmt.Scan(&ele)
		arr = append(arr, ele)
	}
	fmt.Println("Complete array: ", arr)

	const par = 4
	n := int(math.Max(math.Ceil(float64(size)/float64(par)), 1))
	/* if we use math.Floor in an example array of size 10 (something which is not divisible by 4)
	the partition will be 2,2,2,4, obviously will have to change the to:= to math.Max, using the current statements
	we get a partition of 3,3,3,1, the partition for size divisible by 4 will be equal in both respects*/
	var w sync.WaitGroup

	for i := 1; i <= par; i++ {
		from := int((i-1)*n)
		to := int(math.Min(float64(i*n), float64(size)))
		// adding 1 each time a go routine is called
		w.Add(1)
		go sortey(arr[from:to], &w, i)
	}
	// waiting for all 4 goroutines to finish
	w.Wait()

	sort.Ints(arr)
	fmt.Println("complete sorted array: ", arr)
}

func sortey (arr [] int, w *sync.WaitGroup, i int) {
	fmt.Println("Goroutine",i ,"to sort array: ", arr)
	sort.Ints(arr)

	fmt.Println("Array after sorting by goroutine", i, " : ", arr)
	w.Done()
}