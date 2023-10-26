package main

import "fmt"

func main() {

	var mynum int =34
	ptr:=&mynum
	fmt.Println("Value of address",ptr)
	fmt.Println("Value present at address",*ptr)


}