package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func(){
		defer wg.Done()
		work()
	}()
	wg.Wait()
	fmt.Println("Done waiting , main exists")
}

func work() {
	
	time.Sleep(50*time.Millisecond)
	fmt.Println("hey please wait for me")
}