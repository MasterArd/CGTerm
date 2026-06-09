package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/chzyer/readline"
	"github.com/fatih/color"

	"cgterm/commands"
	customconfig "cgterm/customConfig"
)

func main() {
	customconfig.Startup()
	Firstlaunch()

	signal.Notify(make(chan os.Signal, 1), syscall.SIGINT, syscall.SIGTERM)

	rl, err := readline.NewEx(&readline.Config{
		HistoryFile:     "/tmp/cgterm.history",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})

	if err != nil {
		panic(err)
	}
	defer rl.Close()

	fmt.Println("Welcome to CGTerm! Type 'exit' or press Ctrl+D to quit.")

	for {
		cwd, err := os.Getwd()
		if err != nil {
			cwd = "?"
		}

		
		home, _ := os.UserHomeDir()
		if strings.HasPrefix(cwd, home) {
			cwd = "[" + color.CyanString("~") + strings.TrimPrefix(cwd, home) + "]"
		}

		rl.SetPrompt(fmt.Sprintf("%s $> ", cwd))
		line, err := rl.Readline()

		if err != nil {
			if err == readline.ErrInterrupt {
				continue
			}
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			break
		}

		input := strings.TrimSpace(line)
		commandsList := strings.Split(input, "&&")

		for _, cmdStr := range commandsList {
			cmdStr = strings.TrimSpace(cmdStr)

			if cmdStr == "" {
				continue
			}

			parts := strings.Fields(cmdStr)
			if len(parts) == 0 {
				continue
			}

			name := parts[0]
			args := parts[1:]

			// internal commands
			if cmdFunc, ok := commands.Registry[name]; ok {
				cmdFunc(args)
				continue
			}
			

			cmd := exec.Command(name, args...)

			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			cmd.SysProcAttr = &syscall.SysProcAttr{
				Setpgid: false,
			}

			if err := cmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, color.RedString("[-] ")+"%s\n", err)

				break
			}
		}
	}
}

func Firstlaunch() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(color.RedString("could not get home directory"))
		return
	}

	marker := filepath.Join(home, ".CGTerm_init")

	
	if _, err := os.Stat(marker); os.IsNotExist(err) {
		fmt.Println(color.CyanString("--FIRST-RUN--"))
		fmt.Println(color.GreenString("This is probably your first time using CGTerm"))
		fmt.Println(color.GreenString("Support this project on github: https://github.com/MasterArd/CGTerm/"))
		fmt.Println(color.GreenString("This message will only show once"))

		
		f, err := os.Create(marker)
		if err != nil {
			fmt.Println("could not create marker file:", err)
			return
		}
		defer f.Close()

		
		_, err = f.WriteString("1.3.3")
		if err != nil {
			fmt.Println("could not write version to marker file:", err)
			return
		}
	}
}
