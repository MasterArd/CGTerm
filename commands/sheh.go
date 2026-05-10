package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func Sheh(args []string) {
	fmt.Println(color.RedString("Due to the input handling of sheh standard output will be ignored and you wont be able to execute commands. This will be removed once sheh gets updated."))
	fmt.Println(" ")
	fmt.Println(color.CyanString("Thanks and special credit to Waxory/Waxodium: https://github.com/waxodium/ "))
	// 1. Check for binary
	_, err := exec.LookPath("sheh")
	if err != nil {
		var shehNotInstalled string
		fmt.Print(color.RedString("[-]"), " sheh is not found on your system. Install it? (y/n): ")
		fmt.Scanln(&shehNotInstalled)

		if strings.ToLower(shehNotInstalled) == "y" {
			fmt.Println(color.BlueString("[*]"), "Installing sheh via npm...")
			installCmd := exec.Command("npm", "install", "-g", "@waxory/sheh")
			installCmd.Stdout = os.Stdout
			installCmd.Stderr = os.Stderr

			if err := installCmd.Run(); err != nil {
				fmt.Println(color.RedString("[-]"), "npm failed. Is it installed and in your PATH?")
				return
			}
			fmt.Println(color.GreenString("[+]"), "Installation complete.")
		} else {
			return
		}
	}

	// execute command
	cmd := exec.Command("sheh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// capture error from Run() here
	err = cmd.Run()
	if err != nil {
		fmt.Println(color.RedString("[-]"), "Execution error:", err)
	}

	// 3. Print the trailing arrow
  fmt.Println("-->")

}
func init() {
	Register("sheh", Sheh)
}
