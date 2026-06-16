package commands

/*
#cgo CFLAGS: -I${SRCDIR}/c/lib
#include "c/batt.c"
*/
import "C"

func BattCMD(args []string) {
	C.runBatt()
}

func init() {
	Register("batt", BattCMD)
}
