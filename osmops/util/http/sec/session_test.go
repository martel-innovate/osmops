package sec

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

var newTokenManagerErrorIfNilArgsFixtures = []struct {
	provider TokenProvider
	store    TokenStore
}{
	{nil, nil}, {nil, &MemoryTokenStore{}},
	{func() (*Token, error) { return nil, nil }, nil},
}

func TestNewTokenManagerErrorIfNilArgs(t *testing.T) {
	for k, d := range newTokenManagerErrorIfNilArgsFixtures {
		if _, err := NewTokenManager(d.provider, d.store); err == nil {
			t.Errorf("[%d] want error; got: nil", k)
		}
	}
}

type fakeProvider struct {
	callCount int
	lastToken *Token
}

func (p *fakeProvider) generateToken(secondsValid time.Duration) (
	*Token, error) {
	p.callCount += 1
	data := fmt.Sprintf("secret-%d", p.callCount)
	p.lastToken = NewToken(data, secondsAfterNow(secondsValid))
	return p.lastToken, nil
}

func (p *fakeProvider) fetchNewValidToken() (*Token, error) {
	return p.generateToken(600)
}

func (p *fakeProvider) fetchNewExpiredToken() (*Token, error) {
	return p.generateToken(0)
}

func (p *fakeProvider) fetchError() (*Token, error) {
	return nil, errors.New("ouch!")
}

func TestFetchFreshToken(t *testing.T) {
	provider := &fakeProvider{}
	store := &MemoryTokenStore{}
	mngr, _ := NewTokenManager(provider.fetchNewValidToken, store)

	token, err := mngr.GetAccessToken()
	if err != nil {
		t.Fatalf("want: token; got: %v", err)
	}
	if token != provider.lastToken {
		t.Errorf("want: %v; got: %v", provider.lastToken, token)
	}
	if provider.callCount != 1 {
		t.Errorf("want: 1; got: %d", provider.callCount)
	}
	if token != store.token {
		t.Errorf("want: %v; got: %v", token, store.token)
	}
}

func TestUseTokenInStore(t *testing.T) {
	provider := &fakeProvider{}
	store := &MemoryTokenStore{
		token: NewToken("data", secondsAfterNow(600)),
	}
	mngr, _ := NewTokenManager(provider.fetchNewValidToken, store)

	token, err := mngr.GetAccessToken()
	if err != nil {
		t.Fatalf("want: token; got: %v", err)
	}
	if token != store.token {
		t.Errorf("want: %v; got: %v", token, store.token)
	}
	if provider.callCount != 0 {
		t.Errorf("want: 0; got: %d", provider.callCount)
	}
}

func TestRefreshTokenAboutToExpire(t *testing.T) {
	provider := &fakeProvider{}
	store := &MemoryTokenStore{
		token: NewToken("data", secondsAfterNow(10)),
	}
	mngr, _ := NewTokenManager(provider.fetchNewValidToken, store)

	token, err := mngr.GetAccessToken()
	if err != nil {
		t.Fatalf("want: token; got: %v", err)
	}
	if token != provider.lastToken {
		t.Errorf("want: %v; got: %v", provider.lastToken, token)
	}
	if token != store.token {
		t.Errorf("want: %v; got: %v", token, store.token)
	}
	if provider.callCount != 1 {
		t.Errorf("want: 1; got: %d", provider.callCount)
	}
}

func TestFetchNewExpiredToken(t *testing.T) {
	provider := &fakeProvider{}
	store := &MemoryTokenStore{}
	mngr, _ := NewTokenManager(provider.fetchNewExpiredToken, store)

	if token, err := mngr.GetAccessToken(); err == nil {
		t.Errorf("want: error; got: %v", token)
	}
}

func TestFetchError(t *testing.T) {
	provider := &fakeProvider{}
	store := &MemoryTokenStore{}
	mngr, _ := NewTokenManager(provider.fetchError, store)

	if token, err := mngr.GetAccessToken(); err == nil {
		t.Errorf("want: error; got: %v", token)
	}
}
