package util

import (
	"fmt"
	"testing"
)

func TestEmptyString(t *testing.T) {
	if _, err := NewNonEmptyStr(""); err == nil {
		t.Errorf("instantiated a non-empty string with an empty string!")
	}
}

var nonEmptyStringFixtures = []string{" ", "\n", " wada wada "}

func TestNonEmptyString(t *testing.T) {
	for k, d := range nonEmptyStringFixtures {
		if s, err := NewNonEmptyStr(d); err != nil {
			t.Errorf("[%d] want: valid; got: %v", k, err)
		} else {
			if d != s.Value() {
				t.Errorf("[%d] want: %s; got: %s", k, d, s.Value())
			}
		}
	}
}

var invalidHostnameFixtures = []string{
	"", "\n", ":", ":80", "some.host:", "some host", "some host.com",
	"what?is.this", "em@il", "what.the.h*ll",
	"x1234567890123456789012345678901234567890123456789012345678901234.com",
}

func TestInvalidHostname(t *testing.T) {
	for k, d := range invalidHostnameFixtures {
		if err := IsHostname(d); err == nil {
			t.Errorf("[%d] want: error; got: valid", k)
		}
	}
}

var validHostnameFixtures = []string{
	"::123", "1.2.3.4", "_h.com", "a-b.some_where", "some.host",
	"x12345678901234567890123456789012345678901234567890123456789012.com",
}

func TestValidHostname(t *testing.T) {
	for k, d := range validHostnameFixtures {
		if err := IsHostname(d); err != nil {
			t.Errorf("[%d] want: valid; got: %v", k, err)
		}
	}
}

var invalidHostnameAndPortFixtures = []string{
	"", "\n", ":", ":80", "some.host:", "some host:80", "some.host:123456789",
}

func TestInvalidHostnameAndPort(t *testing.T) {
	for k, d := range invalidHostnameAndPortFixtures {
		if err := IsHostAndPort(d); err == nil {
			t.Errorf("[%d] want: error; got: valid", k)
		}
	}
}

var parseHostAndPortFixtures = []struct {
	in       string
	wantHost string
	wantPort int
}{
	{"h:0", "h", 0}, {"h:1", "h", 1}, {"h:65535", "h", 65535},
	{"[::123]:0", "::123", 0}, {"[::123]:1", "::123", 1},
	{"[::123]:65535", "::123", 65535},
	{"1.2.3.4:0", "1.2.3.4", 0}, {"1.2.3.4:1", "1.2.3.4", 1},
	{"1.2.3.4:65535", "1.2.3.4", 65535},
}

func TestParseHostAndPort(t *testing.T) {
	for k, d := range parseHostAndPortFixtures {
		if hp, err := ParseHostAndPort(d.in); err != nil {
			t.Errorf("[%d] want: valid parse; got: %v", k, err)
		} else {
			if d.wantHost != hp.Host() || d.wantPort != hp.Port() {
				t.Errorf("[%d] want: %s:%d; got: %v",
					k, d.wantHost, d.wantPort, hp)
			}

			repr := fmt.Sprintf("%s:%d", d.wantHost, d.wantPort)
			if repr != hp.String() {
				t.Errorf("[%d] want string repr: %s; got: %v", k, repr, hp)
			}
		}
	}
}

var httpUrlErrorFixtures = []string{"", "a", "a/b"}

func TestHttpUrlError(t *testing.T) {
	hp, _ := ParseHostAndPort("x:80")
	for k, d := range httpUrlErrorFixtures {
		if got, err := hp.Http(d); err == nil {
			t.Errorf("[%d] want error; got: %v", k, got)
		}
	}
}

var httpUrlFixtures = []struct {
	inPath string
	want   string
}{
	{"/", "http://x:80/"},
	{"/a", "http://x:80/a"}, {"/a/", "http://x:80/a/"},
	{"/a/b", "http://x:80/a/b"}, {"/a/b/", "http://x:80/a/b/"},
}

func TestHttpUrl(t *testing.T) {
	hp, _ := ParseHostAndPort("x:80")
	for k, d := range httpUrlFixtures {
		got, err := hp.Http(d.inPath)
		if err != nil {
			t.Fatalf("[%d] want string repr: %s; got: %v", k, d.want, err)
		}
		if got.String() != d.want {
			t.Errorf("[%d] want string repr: %s; got: %v", k, d.want, got)
		}
	}
}

var httpsUrlFixtures = []struct {
	inPath string
	want   string
}{
	{"/", "https://x:80/"},
	{"/a", "https://x:80/a"}, {"/a/", "https://x:80/a/"},
	{"/a/b", "https://x:80/a/b"}, {"/a/b/", "https://x:80/a/b/"},
}

func TestHttspUrl(t *testing.T) {
	hp, _ := ParseHostAndPort("x:80")
	for k, d := range httpsUrlFixtures {
		got, err := hp.Https(d.inPath)
		if err != nil {
			t.Fatalf("[%d] want string repr: %s; got: %v", k, d.want, err)
		}
		if got.String() != d.want {
			t.Errorf("[%d] want string repr: %s; got: %v", k, d.want, got)
		}
	}
}

func TestEmptyStrEnum(t *testing.T) {
	e := NewStrEnum()
	if e.IndexOf("") != NotALabel || e.IndexOf("x") != NotALabel {
		t.Errorf("empty enum should have no label indexes")
	}
	if e.LabelOf(0) != "" || e.LabelOf(1) != "" {
		t.Errorf("empty enum should have no labels")
	}
	if e.Validate("") == nil || e.Validate("x") == nil {
		t.Errorf("empty enum should always fail validation")
	}
}

type enumTest = struct {
	StrEnum
	A, B, C EnumIx
}

func NewEnumTest() enumTest {
	return enumTest{
		StrEnum: NewStrEnum("A", "b", "C"),
		A:       0,
		B:       1,
		C:       2,
	}
}

func TestStrEnumLookup(t *testing.T) {
	e := NewEnumTest()
	ixs := []EnumIx{e.A, e.B, e.C}
	for _, ix := range ixs {
		lbl := e.LabelOf(ix)
		if ix != e.IndexOf(lbl) {
			t.Errorf("want: %d == IndexOf(LabelOf(%d)); "+
				"got: %d != IndexOf(%s = LabelOf(%d)) == %d",
				ix, ix, ix, lbl, ix, e.IndexOf(lbl))
		}
	}
}

func TestStrEnumValidation(t *testing.T) {
	e := NewEnumTest()
	if err := e.Validate(e.LabelOf(e.A)); err != nil {
		t.Errorf("[1] want: valid; got: %v", err)
	}
	if err := e.Validate("wada wada"); err == nil {
		t.Errorf("[2] want: error; got: valid")
	}
}

func TestStrEnumCaseInsensitive(t *testing.T) {
	e := NewEnumTest()
	if err := e.Validate("B"); err != nil {
		t.Errorf("want: uppercase B is valid; got: %v", err)
	}
	if e.IndexOf("B") == NotALabel {
		t.Errorf("want: uppercase B is index of b; got: not a label")
	}
}

func TestEmptyIntSet(t *testing.T) {
	s := ToIntSet()
	if s.Contains(0) {
		t.Errorf("want empty; got: %v", s)
	}
}

func TestNonEmptyIntSet(t *testing.T) {
	s := ToIntSet(1, 2)
	if s.Contains(0) {
		t.Errorf("want: 0 not in s; got: %v", s)
	}
	if !s.Contains(1) {
		t.Errorf("want: 1 in s; got: %v", s)
	}
	if !s.Contains(2) {
		t.Errorf("want: 2 in s; got: %v", s)
	}
}
