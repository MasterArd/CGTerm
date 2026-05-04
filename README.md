# CGTerm

**CGTerm** is a minimal terminal-like interface written in Go. It provides a simple interactive command loop with a small set of built-in commands for system interaction and basic screen control.

The project is intentionally lightweight and serves as a foundation for experimenting with terminal behavior, command handling, and modular Go packages.

## Features
* **Interactive prompt:** Simple `-->` interface.
* **Built-in commands:** common commands to help.
* **Modular package structure:** Own commands can be easily added. To do this read [How to make custom commands](./custom-commands.md) 
* **Basic terminal control:** Functions for clearing the screen and listing directories.

## Requirements
- Go install via:https://go.dev/dl/
- 1mb storage

## Building & Running
Building can be done via make or building direct via go
1. clone this repository: `https://github.com/MasterArd/CGTerm.git`
2. Run `Go build .` or `make` followed by `make run` or `./CGTerm`

> [!NOTE]
> **binaries can be found at https://github.com/MasterArd/CGTerm/releases/tag/v1**

## Available Commands

| Command | Description |
| :--- | :--- |
| `host` | Prints the system hostname  |
| `initscreen` | Displays basic screen configuration values |
| `clear` | Clears the terminal screen |
| `exit` | Exits the program |
| `save_settings` | Creates or overwrites a settings file |
| `whoami` | Same as `host` |
| `lsa` | List all files and directories |
| `lsd` | List all directories but not files |
| `lsf` | List all files but not directories |

> If an unknown command is entered, the program will return an error message.

## Project Structure
```text
/
├── main.go          # Entry point and command loop
├── Makefile         # Build and run automation
└── commands/
    └── pkg.go       # standard commands

```

## Contributing
contribution can be done by forking this repository and making a pull request.
> [!NOTE]
> **Contribution might not always be accepted.**

---

### known issues:
- `clear` displaying rogue `[`

## dev notes:
this is an improvement over `UAC`. *(old archive)*: https://masterard.github.io/blue-inft/News.html
