package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func Firstlaunch() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("could not get home directory")
		return
	}
	marker := filepath.Join(home, ".CGTerm_init")

	// if file does NOT exist -> first run
	if _, err := os.Stat(marker); os.IsNotExist(err) {

		fmt.Println(color.CyanString("--FIRST-RUN--"))
		fmt.Println(color.GreenString("This is probably your first time using CGTerm"))
		fmt.Println(color.GreenString("Support this project on github: https://github.com/MasterArd/CGTerm/"))
		fmt.Println(color.GreenString("This message will only show once"))

		// create marker file
		file, err := os.Create(marker)
		if err != nil {
			fmt.Println("could not create marker file:", err)
			return
		}
		file.Close()
	}
}
