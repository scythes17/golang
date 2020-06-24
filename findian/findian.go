package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	var s string
	fmt.Printf("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	s,_ = reader.ReadString('\n')
	s2 := strings.TrimSpace(strings.ToLower(s))
	if( strings.HasPrefix(s2, "i") && strings.HasSuffix(s2, "n") && strings.Contains(s, "a")){
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}