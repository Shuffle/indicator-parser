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
	"encoding/json"
	"fmt"
	"log"
	"os"

	ioc "github.com/Shuffle/indicator-parser/go/ioc"
)

var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("ioc_parser.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	logger = log.New(logFile, "", log.LstdFlags)
}

func logPrint(s string) {
	logger.Println(s)
	fmt.Println(s)
}

//export ParseIOC
func ParseIOC(value *C.char, typesJSON *C.char) *C.char {
	goValue := C.GoString(value)
	// types are useless for now
	results := ioc.Parse(goValue, []ioc.IndicatorType{})

	jsonData, err := json.Marshal(results)
	if err != nil {
		logPrint(fmt.Sprintf("Error marshaling results: %v", err))
		return C.CString("[]")
	}

	return C.CString(string(jsonData))
}

func main() {}
