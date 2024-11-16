package main

import (
	"fmt"
	"time"
)

func worker(done chan bool){
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true //necessary to make the goroutine function synchronous
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	fmt.Println(<- done)
}