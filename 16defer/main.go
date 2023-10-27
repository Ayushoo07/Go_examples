package main

import "fmt"
func main() {
	fmt.Println("Defer")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("tthree")
	fmt.Println("hello world")
}