package main

import (
	"C"
	"encoding/json"
	ioc "github.com/yashsinghcodes/indicator-parser/go/ioc"
)

func ParseIOC(value *C.char, typesJSON *C.char) *C.char {
	goValue := C.GoString(value)
	goTypesJSON := C.GoString(typesJSON)

	var types []ioc.IndicatorType
	err := json.Unmarshal([]byte(goTypesJSON), &types)
	if err != nil {
		return C.CString("[]")
	}

	results := ioc.Parse(goValue, types)

	jsonData, err := json.Marshal(results)
	if err != nil {
		return C.CString("[]")
	}
	return C.CString(string(jsonData))
}

func main() {}
