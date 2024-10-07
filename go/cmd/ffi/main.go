package main

/*
#include <stdio.h>
#include <stdlib.h>
static inline void c_print(const char* s) {
    printf("%s\n", s);
    fflush(stdout);
}
*/
import "C"
import (
	"github.com/Shuffle/indicator-parser/go/parser"
)

//export ParseIOC
func ParseIOC(value *C.char, typesJSON *C.char) *C.char {
	goValue := C.GoString(value)
	result := parser.ParseIndicatorToJSON(goValue)
	return C.CString(result)
}

func main() {}
