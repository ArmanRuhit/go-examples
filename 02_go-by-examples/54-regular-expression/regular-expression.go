package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var Println = fmt.Println

func main() {
	match, _ := regexp.MatchString("(p[a-z]+)ch", "peach")
	Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	Println(r.MatchString("peach"))

	Println(r.FindString("peach punch"))

	Println("idx:", r.FindStringIndex("peach punch"))

	Println(r.FindStringSubmatch("peach punch"))

	Println(r.FindStringSubmatchIndex("peach punch"))

	Println(r.FindAllString("peach punch pinch", -1))

	Println(r.FindAllString("peach punch pinch", 2))

	Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	Println("regexp:", r)

	Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	Println(string(out))
	
}