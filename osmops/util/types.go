package util

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type NonEmptyStr struct{ data string }

func (d NonEmptyStr) Value() string {
	return d.data
}

func NewNonEmptyStr(s string) (NonEmptyStr, error) {
	if len(s) == 0 {
		return NonEmptyStr{}, errors.New("nil or empty string")
	}
	return NonEmptyStr{data: s}, nil
}

type HostAndPort struct {
	h string
	p int
}

func ParsePort(p string) (int, error) {
	p = strings.TrimSpace(p)
	if port, err := strconv.Atoi(p); err == nil {
		if 0 <= port && port <= 65535 {
			return port, nil
		}
	}
	return 0, fmt.Errorf("invalid port: %s", p)
}

var hostnameRx = regexp.MustCompile(
	`^(([a-zA-Z0-9_-]){1,63}\.)*([a-zA-Z0-9_-]){1,63}$`)

// This article explains quite well what makes up a valid hostname:
// - https://en.wikipedia.org/wiki/Hostname

func IsHostname(host string) error {
	if 0 < len(host) && len(host) < 254 {
		if net.ParseIP(host) != nil || hostnameRx.MatchString(host) {
			return nil
		}
	}
	return fmt.Errorf("invalid hostname: %s", host)
}

func ParseHostAndPort(hp string) (*HostAndPort, error) {
	hp = strings.TrimSpace(hp) // (1)
	if host, portString, err := net.SplitHostPort(hp); err != nil {
		return nil, err
	} else {
		if err := IsHostname(host); err != nil { // (2)
			return nil, err
		}
		if port, err := ParsePort(portString); err != nil { // (3)
			return nil, err
		} else {
			return &HostAndPort{host, port}, nil
		}
	}

	// (1) SplitHostPort doesn't trim space, e.g.
	//       SplitHostPort(" h:1 ") == (" h", "1 ", nil)
	// (2) SplitHostPort doesn't check the host part is a valid IP4 or IP6 or
	//     a valid hostname e.g.
	//       SplitHostPort(":123") == ("", "123", nil)
	//       SplitHostPort("??:123") == ("??", "123", nil)
	// (3) SplitHostPort doesn't check the port range, e.g.
	//       SplitHostPort("h:123456789") == ("h", "123456789", nil)
}

func IsHostAndPort(value interface{}) error {
	s, _ := value.(string)
	_, err := ParseHostAndPort(s)
	return err
}

func (d *HostAndPort) Host() string {
	return d.h
}

func (d *HostAndPort) Port() int {
	return d.p
}

func (d *HostAndPort) String() string {
	return fmt.Sprintf("%s:%d", d.h, d.p)
}

func (d *HostAndPort) BuildHttpUrl(secure bool, path string) (*url.URL, error) {
	if u, err := url.ParseRequestURI(path); err != nil {
		return nil, err
	} else {
		if secure {
			u.Scheme = "https"
		} else {
			u.Scheme = "http"
		}
		u.Host = d.String()
		return u, nil
	}
}

func (d *HostAndPort) Http(path string) (*url.URL, error) {
	return d.BuildHttpUrl(false, path)
}

func (d *HostAndPort) Https(path string) (*url.URL, error) {
	return d.BuildHttpUrl(true, path)
}

type StrEnum struct {
	values []string
}

func NewStrEnum(labels ...string) StrEnum {
	e := StrEnum{values: make([]string, len(labels))}
	for k, v := range labels {
		e.values[k] = strings.ToLower(v)
	}
	return e
}

type EnumIx int

const NotALabel EnumIx = -1

func (d StrEnum) IndexOf(label string) EnumIx {
	lbl := strings.ToLower(label)
	for k, v := range d.values {
		if v == lbl {
			return EnumIx(k)
		}
	}
	return NotALabel
}

func (d StrEnum) LabelOf(index EnumIx) string {
	if 0 <= index && int(index) < len(d.values) {
		return d.values[index]
	}
	return "" // better return err? what if one of the labels is ""?!
}

func (d StrEnum) Validate(label interface{}) error {
	if v, ok := label.(string); ok {
		if d.IndexOf(v) != NotALabel {
			return nil
		}
	}
	return fmt.Errorf("not an enum label: %v", label)
}

type IntSet map[int]bool

// the joys of boilerplate, see: https://stackoverflow.com/questions/34018908

func ToIntSet(values ...int) IntSet {
	set := map[int]bool{}
	for _, v := range values {
		set[v] = true
	}
	return set
}

func (s IntSet) Contains(value int) bool {
	_, ok := s[value]
	return ok
}
