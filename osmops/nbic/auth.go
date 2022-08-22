package nbic

import (
	"errors"
	"net/url"

	//lint:ignore ST1001 HTTP EDSL is more readable w/o qualified import
	. "github.com/martel-innovate/osmops/osmops/util/http"
	"github.com/martel-innovate/osmops/osmops/util/http/sec"
)

// UserCredentials holds the data needed to request an OSM NBI token.
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Project  string `json:"project_id"`
}

type tokenPayloadView struct { // only the response fields we care about.
	Id      string  `json:"id"`
	Expires float64 `json:"expires"`
}

type authMan struct {
	creds    UserCredentials
	endpoint *url.URL
	agent    ReqSender
}

// NewAuthz builds a TokenManager to acquire and refresh OSM NBI access tokens.
func NewAuthz(conn Connection, creds UserCredentials, transport ReqSender) (
	*sec.TokenManager, error) {
	if transport == nil {
		return nil, errors.New("nil transport")
	}

	theMan := &authMan{
		creds:    creds,
		endpoint: conn.Tokens(),
		agent:    transport,
	}

	return sec.NewTokenManager(theMan.acquireToken, &sec.MemoryTokenStore{})
}

func (m *authMan) acquireToken() (*sec.Token, error) {
	payload := tokenPayloadView{}
	_, err := Request(
		POST, At(m.endpoint),
		Content(MediaType.YAML), // same as what OSM client does
		Accept(MediaType.JSON),
		JsonBody(m.creds),
	).
		SetHandler(ExpectSuccess(), ReadJsonResponse(&payload)).
		RunWith(m.agent)

	if err != nil {
		return nil, err
	}
	return sec.NewToken(payload.Id, payload.Expires), nil
}
