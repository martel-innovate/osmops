package nbic

import (
	"testing"
)

func TestLookupVimAccIdUseCachedData(t *testing.T) {
	nbic := &Session{
		vimAccMap: map[string]string{"silly_vim": "324567"},
	}
	id, err := nbic.lookupVimAccountId("silly_vim")
	if err != nil {
		t.Errorf("want: 324567; got: %v", err)
	}
	if id != "324567" {
		t.Errorf("want: 324567; got: %s", id)
	}
}

func TestLookupVimAccIdErrorOnMiss(t *testing.T) {
	nbic := &Session{
		vimAccMap: map[string]string{"silly_vim": "324567"},
	}
	if _, err := nbic.lookupVimAccountId("not there!"); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestLookupVimAccIdFetchDataFromServer(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	wantId := "4a4425f7-3e72-4d45-a4ec-4241186f3547"
	if gotId, err := nbic.lookupVimAccountId("mylocation1"); err != nil {
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
	if rr2.req.URL.Path != urls.VimAccounts().Path {
		t.Errorf("want: %s; got: %s", urls.VimAccounts().Path, rr2.req.URL.Path)
	}
}

func TestLookupVimAccIdFetchDataFromServerTokenError(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, UserCredentials{}, nbi.exchange)

	if _, err := nbic.lookupVimAccountId("mylocation1"); err == nil {
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
