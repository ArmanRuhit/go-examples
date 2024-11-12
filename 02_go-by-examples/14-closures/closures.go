// Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.

// In programming, a closure is a function that “remembers” the environment in which it was created. Specifically, a closure is an inner function that has access to variables in the outer function's scope, even after the outer function has finished executing. Closures are powerful because they allow functions to maintain state and are commonly used in functional programming.

package main

import "fmt"

func initSeq() func() int {
	i := 0
	return func () int {
		i++
		return i
	}
}

func main() {
	nextInt := initSeq()
	
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := initSeq()
	fmt.Println(newInts())
}