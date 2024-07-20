package ioc

import (
	"reflect"
	"testing"
)

type tValues[T any] struct {
	input string
	want  T
}

func TestIsIP(t *testing.T) {
	values := []tValues[bool]{
		{"178.45.76.3", true},
		{"192.186.54.55", true},
		{"10.10.43.5", true},
		{"256.1.2.3", false},
		{"example.com", false},
	}

	for _, v := range values {
		if got := IsIp(v.input); got != v.want {
			t.Errorf("IsIp(%q) = %v, want %v\n", v.input, got, v.want)
		}
	}
}

func TestIsUrl(t *testing.T) {
	values := []tValues[bool]{
		{"https://example.com", true},
		{"http://www.google.com", true},
		{"smtp://mymail@outlook.com", false},
		{"example.com", false},
	}

	for _, v := range values {
		if got := IsUrl(v.input); got != v.want {
			t.Errorf("IsUrl(%q) = %v, want %v\n", v.input, got, v.want)
		}
	}
}

func TestIsDomain(t *testing.T) {
	values := []tValues[bool]{
		{"example.com", true},
		{"subdomain.example.com", true},
		{"http://example.com", true},
		{"192.168.1.1", false},
		{"user@mail.com", false},
	}

	for _, v := range values {
		if got := IsDomain(v.input); got != v.want {
			t.Errorf("IsDomain(%q) = %v, want %v\n", v.input, got, v.want)
		}
	}
}

func TestIsEmail(t *testing.T) {
	values := []tValues[bool]{
		{"user@example.com", true},
		{"user.name+tag@example.com", true},
		{"user@subdomain.example.com", true},
		{"example.com", false},
		{"user@", false},
		{"@example.com", false},
	}

	for _, v := range values {
		if got := IsEmail(v.input); got != v.want {
			t.Errorf("IsEmail(%q) = %v, want %v\n", v.input, got, v.want)
		}
	}
}

func TestParse(t *testing.T) {
	input := "https://example.com user@example.com 1.2.3.4 https://google.com/url/booo"
	types := []IndicatorType{IPV4, URL_LINK, Domain, Email}

	want := []Indicator[string]{
		{Type: URL_LINK, Data: "https://example.com"},
		{Type: Domain, Data: "example.com"},
		{Type: Email, Data: "user@example.com"},
		{Type: IPV4, Data: "1.2.3.4"},
		{Type: URL_LINK, Data: "https://google.com/url/booo"},
		{Type: Domain, Data: "google.com"},
	}

	got := Parse[string](input, types)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Parse() returned unexpected result")
		t.Errorf("Got:")
		for _, indicator := range got {
			t.Errorf("  {Type: %v, Data: %q}", indicator.Type, indicator.Data)
		}
		t.Errorf("Want:")
		for _, indicator := range want {
			t.Errorf("  {Type: %v, Data: %q}", indicator.Type, indicator.Data)
		}

		// Compare each element
		for i := 0; i < len(got); i++ {
			if i >= len(want) {
				t.Errorf("Got has more elements than Want")
				break
			}
			if got[i].Type != want[i].Type || got[i].Data != want[i].Data {
				t.Errorf("Mismatch at index %d:", i)
				t.Errorf("  Got:  {Type: %v, Data: %q}", got[i].Type, got[i].Data)
				t.Errorf("  Want: {Type: %v, Data: %q}", want[i].Type, want[i].Data)
			}
		}
		if len(want) > len(got) {
			t.Errorf("Want has more elements than Got")
		}
	}
}

func TestParseDomain(t *testing.T) {
	values := []tValues[string]{
		{"http://example.com", "example.com"},
		{"https://www.google.com/search?q=test", "www.google.com"},
		{"example.com", "example.com"},
		{"subdomain.example.com", "subdomain.example.com"},
	}

	for _, v := range values {
		if got := ParseDomain(v.input); got != v.want {
			t.Errorf("ParseDomain(%q) = %v, want %v\n", v.input, got, v.want)
		}
	}
}
