package main

import (
	"CGTerm/commands"
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/fatih/color"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to CGTerm! Type 'exit' to quit;")

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}
		
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		parts := strings.Split(input, " ")
		name := parts[0]
		args := parts[1:]

		cmd, exists := commands.Registry[name]
		if !exists {
			fmt.Printf("%s error: command not found: %s\n", color.RedString("[-]"), name)
			continue
		}

		cmd(args)
	}
}
