package nbic

import (
	"strings"
	"testing"
)

func TestLookupVnfDescIdUseCachedData(t *testing.T) {
	nbic := &Session{
		vnfdMap: map[string]string{"silly_ns": "324567"},
	}
	id, err := nbic.lookupVnfDescriptorId("silly_ns")
	if err != nil {
		t.Errorf("want: 324567; got: %v", err)
	}
	if id != "324567" {
		t.Errorf("want: 324567; got: %s", id)
	}
}

func TestLookupVnfDescIdErrorOnMiss(t *testing.T) {
	nbic := &Session{
		vnfdMap: map[string]string{"silly_ns": "324567"},
	}
	if _, err := nbic.lookupVnfDescriptorId("not there!"); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestLookupVnfDescIdFetchDataFromServer(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	wantId := "4ffdeb67-92e7-46fa-9fa2-331a4d674137"
	if gotId, err := nbic.lookupVnfDescriptorId("openldap_knf"); err != nil {
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
	if rr2.req.URL.Path != urls.VnfPackagesContent().Path {
		t.Errorf("want: %s; got: %s", urls.VnfPackagesContent().Path, rr2.req.URL.Path)
	}
}

func TestLookupVnfDescIdFetchDataFromServerTokenError(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, UserCredentials{}, nbi.exchange)

	if _, err := nbic.lookupVnfDescriptorId("openldap_knf"); err == nil {
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

func TestMissingDescErrFormat(t *testing.T) {
	err := &missingDescriptor{
		typ:  "x",
		name: "y",
	}
	got := err.Error()

	if !strings.Contains(got, "x") {
		t.Errorf("want: contains type; got: no type")
	}
	if !strings.Contains(got, "y") {
		t.Errorf("want: contains name; got: no name")
	}
}
