# making custom commands:
Custom Commands Guide

This guide explains how to create and register custom commands for your Go application. The system uses a specific structure where commands are initialized using the init() function and registered via the Register function.
1. Requirements

To create a valid custom command, your file must adhere to the following rules:

- Package Name: Must be set to package commands.

- Initialization: Commands must be registered inside the init() function.

- Uppercase function: The command function name must start with an uppercase letter (exported) to be accessible.

2. Example: Creating a "Hello" Command

Below is an example of how to implement a custom hello command. Save this file as something like hello.go inside your commands/ directory.
Go
```

package commands

import "fmt"

// Hello is the function that will be executed when the command is called.
func Hello() {
	fmt.Println("Hello, World!")
}

// init automatically registers the command when the package is loaded.
func init() {
	// Register the command name and its associated function.
	Register("hello", Hello)
}

```

``Register("hello",Hello)`` in this "hello" is the command name and Hello is the function name in your file.
