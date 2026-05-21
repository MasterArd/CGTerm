## Install
### From Source

1. Clone CGTerm github repository
```
git clone https://github.com/MasterArd/CGTerm.git
cd CGTerm
```

2. Install CGTerm to the binary directory
```
make install
```

CGTerm will finally qualified as your shell and able to interact with your terminal.

Launch ``CGTerm`` to the terminal.
CGTerm has built-in commands. Standard and External:

https://github.com/MasterArd/CGTerm/#available-commands

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
CGTerm
└── commands/
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

- By creating a go file inside of the same commands/ directory with this template:

```go
package commands

/*
#cgo CFLAGS: -I${SRCDIR}/c/lib
#include "printf.h"
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
> [!NOTE]
> Please follow the standard syntax of CGO. Having the preamble or the comment block staying exactly no line breaks between ``import "C"``

`#cgo CFLAGS: -I${SRCDIR}/c/lib` line MUST include if any exterior library is used. As libraries should be a simple `.c` file or header `.h` with C `.c`  file placed in `./commands/c/lib/` directory

```txt
CGTerm
└── commands/
    └── c/
        └── lib/
```

While standard C libraries function correctly without the `#cgo CFLAGS: -I${SRCDIR}/c/lib` directive, it should still be included to prevent potential library resolution issues when implementing custom commands.

#### C program example
All custom C commands MUST exactly store in the `./commands/c/` directory and exporting all of the function to initialize to CGTerm

```c
#include <stdio.h>

int goodbye() {
    puts("Goodbye, World\n");
}

```

###  Lastly, running your implements.

By using `make run` or just `go run .` inside of CGTerm root folder, you'll get a resulted `cgterm` output binary file if compiled successfully. 

