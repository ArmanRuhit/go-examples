package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("recieved message", msg)
	default:
		fmt.Println("no message recieved")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent messages:", msg)
	default:
		fmt.Println("no messages sent")
	}

	select {
	case msg = <-messages:
		fmt.Println("recieved message", msg)
	case sig := <-signals:
		fmt.Println("recieved signals", sig)
	default:
		fmt.Println("no activity")
	}
}
