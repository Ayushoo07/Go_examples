package main

import (
	"fmt"
	"time"
)

func main() {
	
	fmt.Println("Synchronous program in go");
	now:=time.Now()	
	task1()
	task2()
	task3()
	task4()

	fmt.Println("Elapsed time :",time.Since(now))
}

func task1(){

	time.Sleep(100*time.Millisecond)
	fmt.Println("Task 1")
}

func task2(){
	fmt.Println("Task 2")
	
}

func task3(){

	time.Sleep(200*time.Millisecond)
	fmt.Println("Task 3")
	
}

func task4(){

	time.Sleep(100*time.Millisecond)
	fmt.Println("Task 4")
	
}
