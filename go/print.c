#include <stdio.h>
#include <stdlib.h>

void cgo_print(char* s) {
    printf("%s\n", s);
    fflush(stdout);
}
