package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in go lang")


	days :=[]string{"a","b","c","d","e","f"}
	
	// for i:=0;i<len(days);i++{
	// 	fmt.Println(days[i])
	// }

	// for i:=range days{
	// 	fmt.Println(days[i])
	// }

	for _,day:=range days{
		fmt.Printf("value %v \n",day)
	}


	var cal int=1
	for cal <10{ 
		if cal == 5 {
			goto lco
			
		} 
		cal++
	}

	lco :
	fmt.Printf("lco \n")
}