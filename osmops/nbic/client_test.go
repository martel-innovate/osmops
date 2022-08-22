package nbic

import (
	"testing"
)

func TestNewNbicErrorOnNilTransport(t *testing.T) {
	if client, err := New(newConn(), usrCreds, nil); err == nil {
		t.Errorf("want: error; got: %+v", client)
	}
}

func TestNewNbicWithDefaultTransport(t *testing.T) {
	client, err := New(newConn(), usrCreds)
	if err != nil {
		t.Errorf("want: client; got: %v", err)
	}
	if client.transport == nil {
		t.Errorf("want: transport; got: nil")
	}
}

func TestGetJsonStopIfResponseNotOkay(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	if _, err := nbic.getJson(urls.buildUrl("/wrong"), nil); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestPostJsonStopIfResponseNotOkay(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	if _, err := nbic.postJson(urls.buildUrl("/wrong"), "42", nil); err == nil {
		t.Errorf("want: error; got: nil")
	}
}
