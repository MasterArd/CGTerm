package gpkg

import (
	"fmt"
	"os"
)

func Host() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Println("Error retrieving hostname:", err)
		return
	}

	fmt.Println("hostname:", name)
}

func Initscreen() {
	const horizontalCharacter = "|"
	const vertialCharacter = "-"
	const screenSizeHorizontal = "40"
	const screenSizeVertical = "60"
	fmt.Println(horizontalCharacter, vertialCharacter, screenSizeHorizontal, screenSizeVertical)
}

func Clear() {
	fmt.Print("\033[2J", "\033[H]")
}

func Exit() {
	os.Exit(0)

}

func Save_settings() {
	var overwriteInput string
	_, err := os.Stat("gpkg_settings.json")
	if err == nil {
		fmt.Println("settings file already exists")
		fmt.Println("do you want to overwrite it? [y/N]")
		fmt.Scanln(&overwriteInput)
		if overwriteInput == "y" {
			os.Create("gpkg_settings.json")
			fmt.Println("file created: gpkg_settings.json, and old overwritten")
		} else {
			fmt.Println("aborted")
			return
		}
	}
}

//move cursor to top left corner: fmt.Print("\033[H")
