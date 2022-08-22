// sec provides basic means to manage the retrieval and refresh of auth tokens.
package sec

import (
	"errors"
	"fmt"
)

// TokenStore defines the how to store and retrieve token data between calls.
type TokenStore interface {
	// Get retrieves the the previously stored token if any. A nil return
	// value means there's no token in the store.
	Get() *Token
	// Set stores the current token, discarding any previous token.
	Set(t *Token)
	// Clear removes the token from the store, if present.
	Clear()
}

// MemoryTokenStore stores tokens in memory.
type MemoryTokenStore struct {
	token *Token
}

// NOTE. We can implement a file system store too if needed, but I don't
// think it's going to be any useful at this stage.

func (s *MemoryTokenStore) Get() *Token {
	if s.token == nil {
		return nil
	}
	return s.token
}

func (s *MemoryTokenStore) Set(t *Token) {
	s.token = t
}

func (s *MemoryTokenStore) Clear() {
	s.token = nil
}

// TokenProvider acquires a fresh token from an auth endpoint, returning
// an error if something goes wrong.
type TokenProvider func() (*Token, error)

// TokenManager manages the storage and lifecycle of tokens.
type TokenManager struct {
	acquireToken TokenProvider
	store        TokenStore
}

// NewTokenManager instantiates a TokenManager, returning an error if any of
// the inputs are nil.
func NewTokenManager(provider TokenProvider, store TokenStore) (
	*TokenManager, error) {
	if provider == nil {
		return nil, errors.New("nil provider")
	}
	if store == nil {
		return nil, errors.New("nil store")
	}
	return &TokenManager{
		acquireToken: provider,
		store:        store,
	}, nil
}

// GetAccessToken retrieves a valid access token if possible, otherwise it
// returns an error. A token is valid if it can still be used for at least
// 30 seconds before it expires.
// If a valid token is in the store, then GetAccessToken returns it. Otherwise
// it delegates the fetching of a fresh token to the TokenProvider. If the
// provider can acquire a valid token, then the token gets stored in the
// TokenStore before returning it. In all other cases, GetAccessToken returns
// an error.
func (m *TokenManager) GetAccessToken() (*Token, error) {
	currentToken := m.store.Get()
	if currentToken != nil && currentToken.SecondsLeftBeforeExpiry() > 30 {
		return currentToken, nil
	}

	m.store.Clear()
	if newToken, err := m.acquireToken(); err != nil {
		return nil, err
	} else {
		if newToken.HasExpired() {
			return nil, fmt.Errorf(
				"auth endpoint returned expired token: %+v", *newToken)
		}
		m.store.Set(newToken)
		return newToken, nil
	}
}
