package commands

import (
	"fmt"
	"os"
	"strings"
)

func cg(args []string) {
	if len(args) < 1 {
		fmt.Println("usage: script <file>")
		return
	}

	runScript(args[0])
}

func runScript(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("script error:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		line = strings.TrimSpace(line)

		// ignore comments
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}

		err := executeLine(line)
		if err != nil {
			fmt.Printf("error on line %d: %v\n", i+1, err)
			return
		}
	}
}


func parseLine(line string) (string, []string) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return "", nil
	}

	return parts[0], parts[1:]
}

func init() {
	Register("cg", cg)
}