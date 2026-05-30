#include "lib/sout.h"
#include <stdio.h>
#include <stdlib.h>


char* versionFetch() {
    FILE *file = fopen("version", "r");
    if (!file) 
        return "Unknown";

    char *buffer = malloc(256);
    char *read = fgets(buffer, 256, file);
    if (read == NULL) {
        buffer = NULL;
    }
    
    fclose(file);
    return buffer;
}

void printHelp() {        
    char *version = versionFetch();

    sout("CGTerm version %s\n", version);
    sout("CGT shell, License GPL-3.0.\n\n");
    
    sout("Standard Commands:\n");
    
    sout("  %s\t\t%s\n", "host",        "Manage host configuration");
    sout("  %s\t%s\n", "initscreen",  "Initialize terminal screen");
    sout("  %s\t\t%s\n", "clear",       "Clear the terminal screen");
    sout("  %s\t%s\n", "whoami",      "Print current user info");
    sout("  %s\t\t%s\n", "lsa",         "List all files");
    sout("  %s\t\t%s\n", "lsd",         "List directories only");
    sout("  %s\t\t%s\n", "lsf",         "List files only");
    sout("  %s\t\t%s\n", "cd",          "Change directory");
    sout("  %s\t%s\n", "version",     "Print CGT version");


    
    sout("\nRepository at <https://github.com/MasterArd/CGTerm> .\n");
}
