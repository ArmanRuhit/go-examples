package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)

	go func ()  {
		<- timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// time.Sleep(time.Second) // and the will prevent the timer2 to stop abruptly
	stop2 := timer2.Stop() // this will stop timer2 abruptly
	if stop2 {
		fmt.Println("Timer 2 Stopped")
	}

	time.Sleep(2 * time.Second)
}