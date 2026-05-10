#include <stdio.h>
#include "lib/printf.h"


#if defined(_WIN32)
    #include <windows.h>
#else
    #include <unistd.h>
#endif



void _putchar(char character) {
#if defined(_WIN32)
    DWORD written;
    WriteFile(GetStdHandle(STD_OUTPUT_HANDLE), &character, 1, &written, NULL);
#else
    int w = write(1, &character, 1);
#endif
}


void clearScreen() {
    

    printf_("\033[H\033[2J\033[3J");
    fflush(stdout);


}

