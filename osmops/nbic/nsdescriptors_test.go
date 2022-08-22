package nbic

import (
	"testing"
)

func TestLookupNsDescIdUseCachedData(t *testing.T) {
	nbic := &Session{
		nsdMap: map[string]string{"silly_ns": "324567"},
	}
	id, err := nbic.lookupNsDescriptorId("silly_ns")
	if err != nil {
		t.Errorf("want: 324567; got: %v", err)
	}
	if id != "324567" {
		t.Errorf("want: 324567; got: %s", id)
	}
}

func TestLookupNsDescIdErrorOnMiss(t *testing.T) {
	nbic := &Session{
		nsdMap: map[string]string{"silly_ns": "324567"},
	}
	if _, err := nbic.lookupNsDescriptorId("not there!"); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestLookupNsDescIdFetchDataFromServer(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	wantId := "aba58e40-d65f-4f4e-be0a-e248c14d3e03"
	if gotId, err := nbic.lookupNsDescriptorId("openldap_ns"); err != nil {
		t.Errorf("want: %s; got: %v", wantId, err)
	} else {
		if gotId != wantId {
			t.Errorf("want: %s; got: %v", wantId, gotId)
		}
	}

	if len(nbi.exchanges) != 2 {
		t.Fatalf("want: 2; got: %d", len(nbi.exchanges))
	}
	rr1, rr2 := nbi.exchanges[0], nbi.exchanges[1]
	if rr1.req.URL.Path != urls.Tokens().Path {
		t.Errorf("want: %s; got: %s", urls.Tokens().Path, rr1.req.URL.Path)
	}
	if rr2.req.URL.Path != urls.NsDescriptors().Path {
		t.Errorf("want: %s; got: %s", urls.NsDescriptors().Path, rr2.req.URL.Path)
	}
}

func TestLookupNsDescIdFetchDataFromServerTokenError(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, UserCredentials{}, nbi.exchange)

	if _, err := nbic.lookupNsDescriptorId("openldap_ns"); err == nil {
		t.Errorf("want: error; got: nil")
	}

	if len(nbi.exchanges) != 1 {
		t.Fatalf("want: 1; got: %d", len(nbi.exchanges))
	}
	rr1 := nbi.exchanges[0]
	if rr1.req.URL.Path != urls.Tokens().Path {
		t.Errorf("want: %s; got: %s", urls.Tokens().Path, rr1.req.URL.Path)
	}
}
