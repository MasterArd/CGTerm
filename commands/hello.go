package commands

import "fmt"

func Hello(args []string) {
	fmt.Println("Hello, world!")
}

func init() {
	Register("hello", Hello)
}
