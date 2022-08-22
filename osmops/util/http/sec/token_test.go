package sec

import (
	"math"
	"testing"
	"time"
)

var tokenDataFixtures = []string{
	"", " ", "xxx", "x\ny", "fbb-34v-99==",
}

func TestTokenData(t *testing.T) {
	for k, d := range tokenDataFixtures {
		token := NewToken(d, 0)
		got := token.String()
		if got != d {
			t.Errorf("[%d] want: %s; got: %s", k, d, got)
		}
	}
}

func secondsBeforeNow(howMany time.Duration) float64 {
	timepoint := time.Now().Add(-howMany * time.Second).Unix()
	return float64(timepoint)
}

func secondsAfterNow(howMany time.Duration) float64 {
	timepoint := time.Now().Add(howMany * time.Second).Unix()
	return float64(timepoint)
}

var tokenExpiryFixtures = []float64{
	math.NaN(), 0, -1, secondsBeforeNow(0), secondsBeforeNow(1),
	1631127131.1251214, // Wed Sep 08 2021 18:52:11 GMT+0000
}

func TestTokenExpiry(t *testing.T) {
	for k, expiry := range tokenExpiryFixtures {
		token := NewToken("secret", expiry)
		if !token.HasExpired() {
			t.Errorf("[%d] want: expired; got: still valid", k)
		}
	}
}

var tokenNotExpiredFixtures = []float64{
	secondsAfterNow(5), secondsAfterNow(600),
	2631127131.1251214, // Sat May 17 2053 20:38:51 GMT+0000
}

func TestTokenNotExpired(t *testing.T) {
	for k, expiry := range tokenNotExpiredFixtures {
		token := NewToken("secret", expiry)
		if token.HasExpired() {
			t.Errorf("[%d] want: still valid; got: expired", k)
		}
	}
}

var tokenSecondsB4ExpiryFixtures = []struct {
	expiry float64
	want   uint64
}{
	{secondsAfterNow(10), 10}, {secondsAfterNow(20), 20},
	{secondsAfterNow(600), 600}, {secondsAfterNow(3600), 3600},
}

func TestTokenSecondsB4Expiry(t *testing.T) {
	for k, d := range tokenSecondsB4ExpiryFixtures {
		token := NewToken("secret", d.expiry)
		got := token.SecondsLeftBeforeExpiry()
		if math.Abs(float64(got)-float64(d.want)) > 2 {
			t.Errorf("[%d] want: %d; got: %d", k, d.want, got)
		}
	}
}
