package parser

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

// ParseIndicator parses an IOC string and returns the results
func ParseIndicator(value string) []ioc.Indicator {
	return ioc.Parse(value, []ioc.IndicatorType{})
}

// ParseIndicatorToJSON parses an IOC string and returns the results as JSON
func ParseIndicatorToJSON(value string) string {
	results := ParseIndicator(value)
	jsonData, err := json.Marshal(results)
	if err != nil {
		logPrint(fmt.Sprintf("Error marshaling results: %v", err))
		return "[]"
	}
	return string(jsonData)
}
