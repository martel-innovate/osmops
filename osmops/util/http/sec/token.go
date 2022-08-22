package sec

import (
	"time"
)

// Token represents an opaque token string credential with an expiry date.
type Token struct {
	expiresAt time.Time
	data      string
}

// SecondsLeftBeforeExpiry returns the number of seconds until the token
// expires or 0 if the token has expired already.
func (t *Token) SecondsLeftBeforeExpiry() uint64 {
	now := time.Now()
	delta := t.expiresAt.Sub(now)
	if delta.Seconds() < 0 {
		return 0
	}
	return uint64(delta.Seconds())
}

// HasExpired tells if the token has gone past its expiry date.
func (t *Token) HasExpired() bool {
	return t.SecondsLeftBeforeExpiry() == 0
}

// String returns the wire representation of the token.
func (t *Token) String() string {
	return t.data
}

// NewToken creates a Token from its wire representation and expiry date as
// number of fractional seconds since the Epoch.
func NewToken(data string, expiry float64) *Token {
	secondsSinceTheEpoch := int64(expiry) // (1, 2)
	if secondsSinceTheEpoch < 0 {
		secondsSinceTheEpoch = 0
	}
	return &Token{
		expiresAt: time.Unix(secondsSinceTheEpoch, 0),
		data:      data,
	}
	// NOTE
	// 1. Truncation. The Go spec says the cast from float to int drops
	// the fractional part, which is what we want since we're working with
	// seconds here. But in general, you gotta be careful with that kind of
	// cast, since e.g. 1.99999999 would get converted to 1. Again, a second
	// difference isn't really a big deal here. If we ever need that kind of
	// accuracy, here's the solution:
	// - https://stackoverflow.com/questions/8022389
	// 2. Overflow. If expiry is NaN, secondsSinceTheEpoch will be negative.
}
