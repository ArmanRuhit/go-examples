package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File{
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintf(f, "data")
}

func closeFile(f *os.File)  {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("./defer.txt")
	defer closeFile(f)
	writeFile(f)
}