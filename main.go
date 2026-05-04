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
	const loop = 1

	for loop < 5 {
		fmt.Print("--> ")

		reader := bufio.NewReader(os.Stdin)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Split(input, " ")
		name := parts[0]
		args := parts[1:]
		cmd, exists := commands.Registry[name]
		if !exists {
			fmt.Print(color.RedString("[-] "), "error in input: command not found: ", input)
			fmt.Println(" ")
			continue
		}

		cmd(args)
	}

}
