package main

import (
	"CGTerm/gpkg"
	"fmt"
	"slices"
)

func main() {
	commands := []string{"host", "initscreen", "clear", "exit", "whoami", "save_settings", ""}

	var input string
	const loop = 1

	for loop < 5 {
		fmt.Print("--> ")
		_, err := fmt.Scanln(&input)
		// if the user press enter Scanln returns error
		// catch err and restart loop
		if err != nil {
			input = ""
			continue
		}
		if slices.Contains(commands, input) {
			switch input {
			case "host":
				gpkg.Host()
			case "initscreen":
				gpkg.Initscreen()
			case "clear":
				gpkg.Clear()
			case "exit":
				gpkg.Exit()
			case "save_settings":
				gpkg.Save_settings()
			case "":
				input = ""
				fmt.Print("")
			}
		} else {
			fmt.Println("[-] error in input: command not found:", input)
		}
	}

}
