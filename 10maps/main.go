package main

import (
	"fmt"

)

func main(){
	fmt.Println("Maps in GoLang")

	language:=make(map[string]string)

	language["py"]="python"
	language["rub"]="ruby"

	fmt.Println(language)
	fmt.Println(language["py"])
	delete(language,"py")
	fmt.Println(language)

	// loops

	for key,value :=range language{
		fmt.Printf("for the key : %v , value is %v\n",key,value)
	}

}