package main

import (
	"fmt"
	"time"
)

func main() {
	
	fmt.Println("Synchronous program in go");
	now:=time.Now()	

	go task1()
	go task2()
	go task3()
	go task4()
	time.Sleep(1000*time.Millisecond)
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
