package main

import (
	"fmt"
	"math/rand"
	
)

func main(){
	fmt.Println("welcome to switch cases")

	var num int=rand.Intn(6)+1

	switch num {
	case 1 : fmt.Println("you can jump 1 space")
	case 2 : fmt.Println("you can jump 2 space")
	case 3 : fmt.Println("you can jump 3 space")
	case 4 : fmt.Println("you can jump 4 space")
	case 5 : fmt.Println("you can jump 5 space")
	case 6 : fmt.Println("you can jump 6 space")

	}
}