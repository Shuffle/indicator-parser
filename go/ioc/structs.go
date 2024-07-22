package ioc

type IndicatorType int

// Future scope to be generic for IP and URL
type Indicator[T string] struct {
	Type IndicatorType `json:"type"`
	Data T             `json:"data"`
}

const (
	IPV4 IndicatorType = iota
	URL_LINK
	Domain
	Email
)
