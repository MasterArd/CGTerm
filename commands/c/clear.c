#include "lib/sout.h"

void clearScreen() {
    sout("\033[H\033[2J\033[3J");
}
