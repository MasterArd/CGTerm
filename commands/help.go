package commands

/*
#cgo CFLAGS: -I${SRCDIR}/c/lib
#include "sout.h"
#include "c/help.c"


void printHelp();
*/
import "C"


func helpPrinter(args []string) {
    C.printHelp()
}


func init() {
    Register("help", helpPrinter);
}

