package main

import "fmt"

const LoginToken ="sajflsaf;ksajf;saf" // public variable first letter capital

func main(){
	var username string="Ayush"
	fmt.Println(username)
	fmt.Printf("Type of the variable is : %T \n", username)

	var is_logged bool=true
	fmt.Println(is_logged)
	fmt.Printf("Type of the variable is : %T \n", is_logged)


	var small_value uint16=45
	fmt.Println(small_value )
	fmt.Printf("Type of the variable is : %T \n", small_value)


	var default_value int
	fmt.Println(default_value )
	fmt.Printf("Type of the variable is : %T \n", default_value)
}