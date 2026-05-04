package main

import (
	"CGTerm/gpkg"
	"fmt"
	"slices"

	"github.com/fatih/color"
)

func main() {
	commands := []string{"host", "initscreen", "clear", "exit", "whoami", "save_settings", "read_test", "help", "lsd", "lsf", "lsa", ""}
	//commandsfname = []string{"gpkg.Host()", "gpkg.Initscreen()", "gpkg.Clear()", "gpkg.Exit", "gpkg.Clear()", "gpkg.Save_settings()" }
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
			case "whoami":
				gpkg.Host()
			case "read_test":
				gpkg.Read_test()
			case "lsd":
				gpkg.Lsd()
			case "lsf":
				gpkg.Lsf()
			case "lsa":
				gpkg.Lsa()
			//case "help":
			//gpkg.Help(config map[string]any)

			// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ still needs to be fixed....

			case "":
				input = ""
				fmt.Print("")
			}
		} else {
			fmt.Print(color.RedString("[-] "), "error in input: command not found: ", input)
			fmt.Println(" ")
		}
	}

}
