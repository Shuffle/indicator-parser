package ioc

import (
	"strings"
	"unicode"
)

func ParseDomain(value string) string {
	if !IsUrl(value) {
		value = "http://" + value
	}

	url, err := ParseURL(value)
	if err != nil {
		return ""
	}
	return url.Host
}

func Parse[T string](value string, types []IndicatorType) []Indicator[T] {
	tokens := strings.FieldsFunc(value, func(c rune) bool {
		//return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		return unicode.IsSpace(c)
	})

	var results []Indicator[T]
	for _, token := range tokens {
		switch {
		case IsIp(token):
			results = append(results, Indicator[T]{
				Type: IPV4,
				Data: any(token).(T),
			})
		case IsUrl(token):
			results = append(results, Indicator[T]{
				Type: URL_LINK,
				Data: any(token).(T),
			})
			fallthrough
		case IsDomain(token):
			results = append(results, Indicator[T]{
				Type: Domain,
				Data: any(ParseDomain(token)).(T),
			})
		case IsEmail(token):
			results = append(results, Indicator[T]{
				Type: Email,
				Data: any(token).(T),
			})
		}
	}

	return results
}
