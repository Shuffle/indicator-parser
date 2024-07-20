package ioc

type IndicatorType string

type Indicator[T any] struct {
	Type IndicatorType
	Data T
}

const (
	IPV4     IndicatorType = "ipv4"
	URL_LINK IndicatorType = "url"
	Domain   IndicatorType = "domain"
	Email    IndicatorType = "email"
)
