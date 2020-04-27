package main

import "fmt"

type MyFunx func(string) error

var _ MyFunx = New

func New(string) error {
	fmt.Println("qqwqw")
	return nil
}

func main() {
	fmt.Println("started")
}
