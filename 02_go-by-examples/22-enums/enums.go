package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
	StateUnknown
)

var stateName = map[ServerState]string{
	StateIdle: "idle",
	StateConnected: "connected",
	StateError: "error",
	StateRetrying: "retrying",
	StateUnknown: "unknown",
}

func (ss ServerState) String() string{
	return stateName[ss]
}

func transition(s ServerState) ServerState  {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func main() {
	// ns := transition(StateUnknown) //It will cause an error
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}