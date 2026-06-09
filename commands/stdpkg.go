package commands

/*
#cgo CFLAGS: -I${SRCDIR}/c/
#
*/

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)


func Version(args []string) {
    home, err := os.UserHomeDir()
    var content []byte
    content, err = os.ReadFile(filepath.Join(home, ".CGTerm_init"))
    if err != nil {
        log.Fatal("Could not read file:", err)
    }
    fmt.Println(string(content))
}

func Print(args []string) {
	if args[0] == "$SHELL" {
		fmt.Println("CGTerm")
	} else {
		fmt.Println(args[0])
	}
}

func Host(args []string) {
	name, err := os.Hostname()
	if err != nil {
		fmt.Println("Error retrieving hostname:", err)
		return
	}

	fmt.Println("hostname:", name)
}

func Exit(args []string) {
	os.Exit(0)

}

func Lsd(args []string) {
	entries, err := os.ReadDir(".") // . = current dir
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("> ")
	for _, entry := range entries {
		if entry.IsDir() {

			color.New(color.FgHiGreen).Fprintln(os.Stdout, entry.Name())
			fmt.Print("  ")
		}
	}
}
func Lsf(args []string) {
	entries, err := os.ReadDir(".") // . = current dir
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("> ")
	for _, entry := range entries {
		if !entry.IsDir() {

			color.New(color.FgHiBlue).Fprintln(os.Stdout, entry.Name())
			fmt.Print("  ")
		}
	}
}

func Lsa(args []string) {
	entries, err := os.ReadDir(".") // . = current dir
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("> ")
	for _, entry := range entries {
		if !entry.IsDir() {

			color.New(color.FgHiBlue).Fprintln(os.Stdout, entry.Name())
			fmt.Print("  ")
		} else if entry.IsDir() {
			color.New(color.FgHiGreen).Fprintln(os.Stdout, entry.Name())
			fmt.Print("  ")
		}
	}
}
func Lse(args []string) {
    entries, err := os.ReadDir(".") // . = current dir
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Print("> ")
    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }
        info, err := entry.Info()
        if err != nil {
            continue
        }
        if info.Mode()&0111 != 0 {
            color.New(color.FgHiRed).Fprintln(os.Stdout, entry.Name())
            fmt.Print("  ")
        }
    }
}
func Cd(args []string) {
	target := ""
	if len(args) == 0 || args[0] == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error finding home directory:", err)
			return
		}
		target = home
	} else {
		target = args[0]
		if strings.HasPrefix(target, "~") {
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Error finding home directory:", err)
				return
			}
			if target == "~" {
				target = home
			} else {
				target = filepath.Join(home, target[1:])
			}
		}
	}

	if err := os.Chdir(target); err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	_, err := os.Getwd()
	if err != nil {
		fmt.Println("Error retrieving current directory:", err)
		return
	}
}

func init() {
	Register("hello", Hello)
	Register("host", Host)
	Register("exit", Exit)
	Register("cd", Cd)
	Register("lsd", Lsd)
	Register("lsf", Lsf)
	Register("lsa", Lsa)
	Register("lse", Lse)
	Register("version", Version)
	Register("print", Print)
}
