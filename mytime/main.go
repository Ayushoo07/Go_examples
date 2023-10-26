package main

import (
	"fmt"
	"time"
)

func main() {
	present:=time.Now();
	fmt.Println("Time is :=",present)

	fmt.Println("new Time parsing",present.Format("01-02-2006 15:04:05 Monday"))






}