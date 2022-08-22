package nbic

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/martel-innovate/osmops/osmops/util"
)

var usrCreds = UserCredentials{
	Username: "admin", Password: "admin", Project: "admin",
}

func sameCreds(expected UserCredentials, req *http.Request) bool {
	got := UserCredentials{}
	json.NewDecoder(req.Body).Decode(&got)
	return reflect.DeepEqual(expected, got)
}

func newConn() Connection {
	address, _ := util.ParseHostAndPort("localhost:8080")
	return Connection{Address: *address}
}

type mockTransport struct {
	received    *http.Request
	replyWith   *http.Response
	timesCalled int
}

func (m *mockTransport) send(req *http.Request) (*http.Response, error) {
	m.received = req
	m.timesCalled += 1
	return m.replyWith, nil
}

func TestGetExpiredToken(t *testing.T) {
	mock := &mockTransport{
		replyWith: &http.Response{
			StatusCode: http.StatusOK,
			Body:       stringReader(expiredNbiTokenPayload),
		},
	}
	mngr, _ := NewAuthz(newConn(), usrCreds, mock.send)

	if _, err := mngr.GetAccessToken(); err == nil {
		t.Errorf("want: error; got: nil")
	}
	if mock.timesCalled != 1 {
		t.Errorf("want: 1; got: %d", mock.timesCalled)
	}
	if !sameCreds(usrCreds, mock.received) {
		t.Errorf("want: same usrCreds; got: different")
	}
}

func TestGetValidToken(t *testing.T) {
	mock := &mockTransport{
		replyWith: &http.Response{
			StatusCode: http.StatusOK,
			Body:       stringReader(validNbiTokenPayload),
		},
	}
	mngr, _ := NewAuthz(newConn(), usrCreds, mock.send)
	token, err := mngr.GetAccessToken()

	if err != nil {
		t.Errorf("want: token; got: %v", err)
	}

	wantData := "TuD41hLjDvjlR2cPcAFvWcr6FGvRhIk2"
	if token.String() != wantData {
		t.Errorf("want: %s; got: %s", wantData, token.String())
	}
	if token.HasExpired() {
		t.Errorf("want: still valid; got: expired")
	}

	if mock.timesCalled != 1 {
		t.Errorf("want: 1; got: %d", mock.timesCalled)
	}
	if !sameCreds(usrCreds, mock.received) {
		t.Errorf("want: same usrCreds; got: different")
	}
}

func TestGetTokenStopIfResponseNotOkay(t *testing.T) {
	mock := &mockTransport{
		replyWith: &http.Response{
			StatusCode: 500,
			Body:       stringReader(validNbiTokenPayload),
		},
	}
	mngr, _ := NewAuthz(newConn(), usrCreds, mock.send)
	if _, err := mngr.GetAccessToken(); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestGetTokenPayloadWithNoTokenFields(t *testing.T) {
	mock := &mockTransport{
		replyWith: &http.Response{
			StatusCode: http.StatusOK,
			Body:       stringReader(`{"x": 1}`),
		},
	}
	mngr, _ := NewAuthz(newConn(), usrCreds, mock.send)
	if _, err := mngr.GetAccessToken(); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestGetTokenPayloadDeserializationError(t *testing.T) {
	mock := &mockTransport{
		replyWith: &http.Response{
			StatusCode: http.StatusOK,
			Body:       stringReader(`["expecting", "an object", "not an array!"]`),
		},
	}
	mngr, _ := NewAuthz(newConn(), usrCreds, mock.send)
	if _, err := mngr.GetAccessToken(); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestNewAuthzErrorOnNilTransport(t *testing.T) {
	if _, err := NewAuthz(Connection{}, UserCredentials{}, nil); err == nil {
		t.Errorf("want: error; got: nil")
	}
}
