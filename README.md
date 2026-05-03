# CGTerm

**CGTerm** is a minimal terminal-like interface written in Go. It provides a simple interactive command loop with a small set of built-in commands for system interaction and basic screen control.

The project is intentionally lightweight and serves as a foundation for experimenting with terminal behavior, command handling, and modular Go packages.

## Features
* **Interactive prompt:** Simple `-->` interface.
* **Built-in command handling:** Standardized command processing.
* **Modular package structure:** commands separated into the `gpkg` directory.
* **Basic terminal control:** Functions for clearing the screen and resetting the cursor.
* **Settings handling:** Basic file creation and management.

## Requirements
- Go
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
| `host` | Prints the system hostname *(not usefull)* |
| `initscreen` | Displays basic screen configuration values |
| `clear` | Clears the terminal screen |
| `exit` | Exits the program |
| `save_settings` | Creates or overwrites a settings file |
| `whoami` | *(Defined but not implemented)* |

> If an unknown command is entered, the program will return an error message.

## Project Structure
```text
/
├── main.go          # Entry point and command loop
├── Makefile         # Build and run automation
└── gpkg/
    └── pkg.go       # Command implementations

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
