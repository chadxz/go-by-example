package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	u, err := url.Parse("postgres://user:pass@host.com:5432/path?mykey=myval#myfrag")
	if err != nil {
		panic(err)
	}

	fmt.Println(u)
	fmt.Println("scheme:", u.Scheme)
	fmt.Println("user:", u.User)
	fmt.Println("username:", u.User.Username())
	p, set := u.User.Password()
	fmt.Println("password:", p, set)
	fmt.Println("host:", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host and port:", host, port)
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	queryValues, _ := url.ParseQuery(u.RawQuery)
	fmt.Printf("%T %v\n", queryValues, queryValues)
}
