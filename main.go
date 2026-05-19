package main

import (
	"CGTerm/commands"
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func main() {
	Firstlaunch()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to CGTerm! Type 'exit' or press Ctrl+D to quit.")

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err.Error())
			break
		}

		input := strings.TrimSpace(line)
		if input == "" {
			continue
		}
		if input == "exit" {
			break
		}

		parts := strings.Fields(input)
		name := parts[0]
		args := parts[1:]

		cmdFunc, exists := commands.Registry[name]
		if exists {
			cmdFunc(args)
			continue
		}

		c := exec.Command(name, args...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Stdin = os.Stdin

		if err := c.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "%s error: %s\n", color.RedString("[-]"), err)
			continue
		}

		c.Wait()
	}
}

func Firstlaunch() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("could not get home directory")
		return
	}
	marker := filepath.Join(home, ".CGTerm_init")

	// if file does  !exist > first run
	if _, err := os.Stat(marker); os.IsNotExist(err) {

		fmt.Println(color.RedString("WARNING:"))
		fmt.Println(color.RedString("Sudo command is currently bugged and may crash."))
		fmt.Println(color.RedString("Avoid using it until fixed."))

		// create marker
		file, err := os.Create(marker)
		if err != nil {
			fmt.Println("could not create marker file:", err)
			return
		}
		file.Close()
	}
}
