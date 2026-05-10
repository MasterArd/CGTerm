package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func Host(args []string) {
	name, err := os.Hostname()
	if err != nil {
		fmt.Println("Error retrieving hostname:", err)
		return
	}

	fmt.Println("hostname:", name)
}

func Initscreen(args []string) {
	const horizontalCharacter = "|"
	const vertialCharacter = "-"
	const screenSizeHorizontal = "40"
	const screenSizeVertical = "60"
	fmt.Println(horizontalCharacter, vertialCharacter, screenSizeHorizontal, screenSizeVertical)
}


func Exit(args []string) {
	os.Exit(0)

}

func Save_settings(args []string) {
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

func Read_test(args []string) {
	// 1. Read the file into a byte slice
	data, err := os.ReadFile("gpkg_settings.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 2. Unmarshal the bytes into your variable
	var config map[string]any
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("Parsed Data:", config)

}

/*
func Help(config map[string]any) {
	fmt.Println(config)
}
*/

func Lsd(args []string) {
	entries, err := os.ReadDir(".") // . = current dir
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("--> ")
	for _, entry := range entries {
		if entry.IsDir() {

			color.New(color.FgHiGreen).Fprint(os.Stdout, entry.Name())
			fmt.Print("  ")
		}
	}
	fmt.Println(" ")
}
func Lsf(args []string) {
	entries, err := os.ReadDir(".") // . = current dir
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("--> ")
	for _, entry := range entries {
		if !entry.IsDir() {

			color.New(color.FgHiBlue).Fprint(os.Stdout, entry.Name())
			fmt.Print("  ")
		}
	}
	fmt.Println(" ")
}

func Lsa(args []string) {
	entries, err := os.ReadDir(".") // . = current dir
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("--> ")
	for _, entry := range entries {
		if !entry.IsDir() {

			color.New(color.FgHiBlue).Fprint(os.Stdout, entry.Name())
			fmt.Print("  ")
		} else if entry.IsDir() {
			color.New(color.FgHiGreen).Fprint(os.Stdout, entry.Name())
			fmt.Print("  ")
		}
	}
	fmt.Println(" ")
}

// move cursor to top left corner: fmt.Print("\033[H")
func init() {
	Register("hello", Hello)
	Register("host", Host)
	Register("initscreen", Initscreen)
	Register("exit", Exit)
	Register("save_settings", Save_settings)
	Register("read_test", Read_test)
	//Register("help", Help) <-- still needs to be fixed
	Register("lsd", Lsd)
	Register("lsf", Lsf)
	Register("lsa", Lsa)

}
