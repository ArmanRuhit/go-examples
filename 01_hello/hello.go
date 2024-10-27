package main
import (
	"fmt"
	"rsc.io/quote"
	"example.com/greetings"
	"log"
)

func main()  {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	callGreetings("Ruhit")
	callGreetings("Ruhit")
	callGreetings("Ruhit")
	names := []string{"Rafid", "Rahin", "Ammu"}
	callMultipleGreetings(names)
	callGreetings("")
}

func callGreetings(name string){
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Request a greeting message
	message, err := greetings.Hello(name);

	// If an error was returned, print it to the console and exit the program
	if err != nil {
		log.Fatal(err);
	}

	//if no error was returned, print the returned message
	// to the console
	fmt.Println(message)

}

func callMultipleGreetings(name []string){
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("multiple greetings: ")
	log.SetFlags(0)

	//Request a greeting message
	messages, err := greetings.Hellos(name);

	// If an error was returned, print it to the console and exit the program
	if err != nil {
		log.Fatal(err);
	}

	//if no error was returned, print the returned message
	// to the console
	fmt.Println(messages)

}