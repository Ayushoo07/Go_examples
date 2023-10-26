package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please rate my pizza app")
	reader :=bufio.NewReader(os.Stdin)
	input,_:=reader.ReadString('\n')
	fmt.Println("rating given by us",input)

	numrating,err:=strconv.ParseFloat(strings.TrimSpace(input),64)
	if err !=nil{
		fmt.Println("Error is ::::",err)
	} else
	{
	fmt.Println("Num rating is ::",numrating)
	}
	

}
