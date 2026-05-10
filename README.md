<p align="center" >
    <img style="display: flex; justify-content: center; align-items: center; height: 250px" src="https://masterard.github.io/blue-inft/CGTerm.png" > 
</p>


<h1 align="center"> CGTerm </h1> 

**CGTerm** is a minimal terminal-like interface written in Go & C. It provides a simple interactive command loop with a small set of built-in commands for system interaction and basic screen control.

The project is intentionally lightweight and serves as a foundation for experimenting with terminal behavior, command handling, and modular Go packages.

## Features
* **Interactive prompt:** Simple `-->` interface.
* **Built-in commands:** common commands to help.
* **Modular package structure:** Own commands can be easily added. To do this read [How to make custom commands](./custom-commands.md) 
* **Basic terminal control:** Functions for clearing the screen and listing directories.

## Requirements
- [Golang](https://go.dev/dl/)
- 5mb+ storage

## Building & Running
Building can be done with make or building directly with go

1. clone CGTerm repository
```
git clone https://github.com/MasterArd/CGTerm.git
cd CGTerm
```

2. Build using make
```
make build
```
**with GO**
```
go build .
```

> [!NOTE]
> **binaries can be found at [Releases](https://github.com/MasterArd/CGTerm/releases/)**

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
    └── c/           # standard commands written in C

```

## Contributing
contribution can be done by forking this repository and making a pull request.
> [!NOTE]
> **Contribution might not always be accepted.**

---

### Issues:
- `clear` displaying rogue `[` (fixed)


## dev notes:
this is an improvement over `UAC`. *(old archive)*: https://masterard.github.io/blue-inft/News.html

