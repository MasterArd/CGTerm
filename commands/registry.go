// commands/registry.go
package commands
import (
    "fmt"
)

type Command func([]string)

var Registry = map[string]Command{}

func Register(name string, cmd Command) {
    Registry[name] = cmd
}

func executeLine(line string) error {
	cmd, args := parseLine(line)

	if cmd == "" {
		return nil
	}

	if handler, ok := Registry[cmd]; ok {
		handler(args)
		return nil
	}

	return fmt.Errorf("unknown command: %s", cmd)
}