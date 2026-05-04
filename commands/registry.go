// commands/registry.go
package commands

type Command func([]string)

var Registry = map[string]Command{}

func Register(name string, cmd Command) {
    Registry[name] = cmd
}