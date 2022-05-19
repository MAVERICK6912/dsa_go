package main

import "fmt"

type customString string

func main() {
	var v customString
	v = "some random string"
	v.replaceSpace()
}

func (s customString) replaceSpace() customString {
	for _, char := range s {
		fmt.Println(string(char))
	}
	return s
}
