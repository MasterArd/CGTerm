package commands

/*
#cgo CFLAGS: -I${SRCDIR}/c/lib
#include "printf.h"
#include "c/clear.c"

void clearScreen();
*/
import "C"

func clearCMD(args []string) {
    C.clearScreen()
}

func init() {
    Register("clear", clearCMD)
}
