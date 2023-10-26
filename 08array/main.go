package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in GoLang")

	var fruitList [4]string
	fruitList[0]="Apple"
	fruitList[1]="Mango"
	fruitList[3]="Banana"

	fmt.Println("Display of FruitListArray ::",fruitList)
	fmt.Println("len of fruitList : ",len(fruitList))

	var namelist=[5]string{"Ayush","Aman","Shivam"}

	fmt.Println("namelist len :",len(namelist))


}