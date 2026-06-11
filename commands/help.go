package commands

import (
    "fmt"
)

func Help(args []string) {
	fmt.Printf("  %s\t\t%s\n", "host", "Manage host configuration")
	fmt.Printf("  %s\t\t%s\n", "clear", "Clear the terminal screen")
	fmt.Printf("  %s\t%s\n", "whoami", "Print current user info")
	fmt.Printf("  %s\t\t%s\n", "lsa", "List all files")
	fmt.Printf("  %s\t\t%s\n", "lsd", "List directories only")
	fmt.Printf("  %s\t\t%s\n", "lsf", "List files only")
	fmt.Printf("  %s\t\t%s\n", "lse", "List executables only")
	fmt.Printf("  %s\t\t%s\n", "cd", "Change directory")
	fmt.Printf("  %s\t\t%s\n", "cg", "execute CGTerm shell scripts")
	fmt.Printf("  %s\t%s\n", "version", "Print CGT version")
	fmt.Printf("  %s\t%s\n", "standard shell commands", "")
}

func init() {
	Register("help", Help)
}
