package customconfig

import (
	"fmt"
	"os"
)

func Startup() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("could not get home dir:", err)
		return
	}

	err = os.Chdir(home)
	if err != nil {
		fmt.Println("could not cd to home:", err)
		return
	}
}
