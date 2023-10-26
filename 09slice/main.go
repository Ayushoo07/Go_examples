package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices");
	var fruitlist=[]string{};
	fmt.Printf("type of fruitList is : %T \n",fruitlist)
	fruitlist=append(fruitlist, "Apple","mango")
	fmt.Println("fruitList :",fruitlist)
	fruitlist=append(fruitlist[:1])

	fmt.Println("fruitList :",fruitlist)


	highscores :=make([] int,4)
	highscores[0]=12
	highscores[1]=13
	highscores[2]=14
	highscores[3]=15

	fmt.Println("Highscores :",highscores)

	highscores = append(highscores,789,798,875 )

	fmt.Println(highscores)
	var rmidx int =2

	highscores=append(highscores[:rmidx],highscores[rmidx+1:]... )
	fmt.Println("Highscores :",highscores)




}