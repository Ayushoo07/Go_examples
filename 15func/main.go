package main

import "fmt"

func main() {
	hello()
	res:=add(1,2,3,4,5,6,6)

	fmt.Println(res)
}
func hello() {
	fmt.Println("hello world")
}
func add( val...int) int {
	total:=0
	for _,val :=range val{
		total+=val
	}
	return total
}