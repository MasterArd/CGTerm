## Making custom commands
### Custom Commands Guides

This guide explains how to register custom CGTerm commands by creating and registering your commands. The system uses a modular registration pattern, allowing you to write logic in Go or bridge to native C code via **cgo**.


1. Command Requirements

To ensure your command is valid to CGTerm, your file must adhere to these rules:

- Package Name: Every command file must belong to package commands.

- The function must accept a slice of strings for arguments: func Name(args []string).

- While the function can be private within the package, it is standard practice to name it descriptively before passing it to the registry.

- Self-Registration: Commands must register themselves using the init() function.

##### Implementing Examples
For simple tasks, you can write the logic entirely in Go. Create your new GO file in the commands/ directory:
```txt
CGTerm/
├── commands/
```

```go
package commands

import "fmt"

func helloCMD(args []string) {
    fmt.Println("Hello from the Go layer!")
}

func init() {
    Register("hello", helloCMD)
}
```

If you want to write your new command in C. Use cgo preamble to link your C functions.

- By creating a go file inside of commands/ directory with this template:

```go
package commands

/*
#cgo CFLAGS: -I${SRCDIR}/c/lib
#include "c/clear.c"

void clearScreen();
*/
import "C"

func clearCMD(args []string) {
    // Function call defined from your C file
    C.clearScreen() 
}

func init() {
    // registering command
    Register("clear", clearCMD)

}
```
with your actual c file from c/ directory. External library can be store in c/lib/ with just c and header file for each.

```txt
CGTerm
└── commands/
    └── c/
        └── lib/
```


## Lastly, running your implements.

By using ``make run`` or just ``go run .`` inside of CGTerm root folder
