package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Fastfetch(args []string) {
	cmd := exec.Command("fastfetch", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err, "is it installed?")
	}
}

func CGTermff(args []string) {
	home, err := os.UserHomeDir()
	var content []byte
	content, err = os.ReadFile(filepath.Join(home, ".CGTerm_init"))
	if err != nil {
		log.Fatal("Could not read file:", err)
	}
	output := fmt.Sprintf("version %s", content)
	fmt.Println("┌─┐┌─┐┌┬┐┌─┐┬─┐┌┬┐")
	fmt.Println("│  │ ┬ │ ├┤ ├┬┘│││ ")
	fmt.Println("└─┘└─┘ ┴ └─┘┴└─┴ ┴")

	padding := (20 - len(output)) / 2
	fmt.Printf("%*s%s\n", padding, "", output)
}

func init() {
	Register("fastfetch", Fastfetch)
	Register("cgtermff", CGTermff)
}
