package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome:="welcome to the user"

	fmt.Println(welcome)

	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating of our pizza :")

	input,_:=reader.ReadString('\n')
	fmt.Println("rhanks for rating,",input)
	fmt.Printf("the type of input %T",input)
}