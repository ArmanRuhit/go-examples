package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)

	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("Username: ", u.User.Username())

	p, _ := u.User.Password()
	fmt.Println("Password: ", p)

	fmt.Println("u.Host:", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("Host: ", host)
	fmt.Println("Port: ", port)

	fmt.Println("Path: ",u.Path)
	fmt.Println("Fragment: ", u.Fragment)

	fmt.Println("RawQuery: ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("ParsedQuery:",m)
	fmt.Println(m["k"][0])
}