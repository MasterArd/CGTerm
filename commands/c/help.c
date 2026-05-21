#include "lib/sout.h"
#include "include/version.h"

void helpHost();
void helpInitscreen();
void helpClear();
void helpWhoami();
void helpLsa();
void helpLsd();
void helpLsf();
void helpCd();

void printHelp() {        
    // version & release (typeof #define) forwarded from include/version.h
    sout("CGTerm cgt, version %s(%s)-release and is GPL-3.0 LICENSE\n", version, release);
    sout("Shell Standard Commands. Type `help` to see this list\nType `help command` to get detail description about the command\n\n");
    sout("  host, initscreen, clear,\n");
    sout("  whoami, lsa, lsd, lsf, cd\n");
}

void helpHost() {
    sout("Print System's host name.\n");
}

void helpInitscreen() {
    sout("Display screen size\n");
}

void helpClear() {
    sout("Clear the terminal buffer and reset the cursor position.\n");
}

void helpWhoami() {
    sout("Print System's username\n");
}

void helpLsa() {
    sout("List all files and directories\n");
}

void helpLsd() {
    sout("List all directories\n");
}

void helpLsf() {
    sout("List all files\n");
}

void helpCd() {
    sout("Change directory\n");
    sout("Usage: cd <directory>\n");
    sout("Directory should be an absolute path or a relative path\n");
}
