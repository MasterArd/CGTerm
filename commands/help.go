package commands

/*
#cgo CFLAGS: -I${SRCDIR}/c/lib
#include "sout.h"
#include "c/help.c"


void printHelp();
*/
import (
	"C"
    "fmt"
)

func helpPrinter(args []string) {
    fmt.Println("Help is currently being bug fixed,") // this will be removed once the bug is fixed
	//C.printHelp()
}

func init() {
	Register("help", helpPrinter)
}
