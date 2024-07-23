package ioc

type IndicatorType string

// Future scope to be generic for IP and URL
type Indicator[T string] struct {
	Type IndicatorType `json:"type"`
	Data T             `json:"data"`
}

const (
	IPV4     IndicatorType = "IPV4"
	URL_LINK IndicatorType = "URL_LINK"
	Domain   IndicatorType = "DOMAIN"
	Email    IndicatorType = "EMAIL"
)
