package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("fork - join model");
	go work() // fork point
	time.Sleep(50*time.Millisecond)
	fmt.Println("no join point")
	
}

func work(){
	time.Sleep(100*time.Millisecond)
	fmt.Println("Printing Some Stuff")
}