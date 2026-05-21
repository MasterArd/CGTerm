package commands

/*
#cgo CFLAGS: -I${SRCDIR}/c/lib
#include "sout.h"
#include "c/help.c"


void helpHost();
void helpInitscreen();
void helpClear();
void helpWhoami();
void helpLsa();
void helpLsd();
void helpLsf();
void helpCd();

void printHelp();
*/
import "C"
import "fmt"


func helpPrinter(args []string) {
    if len(args) == 0 {
        C.printHelp()
        return
    }

    switch args[0] {
    case "host":
        C.helpHost()
    case "initscreen":
        C.helpInitscreen()
    case "whoami":
        C.helpWhoami()
    case "lsa":
        C.helpLsa()
    case "lsd":
        C.helpLsd()
    case "lsf":
        C.helpLsf()
    case "cd":
        C.helpCd()
    case "clear":
        C.helpClear()
    default:
        fmt.Printf("Invalid Standard Command: %s\n\n", args[0])
        C.printHelp()
    }
}


func init() {
    Register("help", helpPrinter);
}

