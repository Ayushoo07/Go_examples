package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang")

	Ayush:=User{"Ayush","21","ayush@gmail.com",true}

	fmt.Println(Ayush)
	fmt.Printf("struct is :%+v\n",Ayush)


}

type User struct
{
	Name string
	Age string
	Email string
	Status bool
}

