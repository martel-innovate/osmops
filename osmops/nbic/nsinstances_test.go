package nbic

import (
	"io/ioutil"
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestBuildNsInstanceMap(t *testing.T) {
	vs := []nsInstanceView{
		{Id: "1", Name: "a"}, {Id: "2", Name: "a"}, {Id: "3", Name: "b"},
	}
	nsMap := buildNsInstanceMap(vs)

	if got, ok := nsMap["a"]; !ok {
		t.Errorf("want: a; got: nil")
	} else {
		want := []string{"1", "2"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want: %v; got: %v", want, got)
		}
	}

	if got, ok := nsMap["b"]; !ok {
		t.Errorf("want: b; got: nil")
	} else {
		want := []string{"3"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want: %v; got: %v", want, got)
		}
	}

	if got, ok := nsMap["c"]; ok {
		t.Errorf("want: nil; got: %v", got)
	}
}

func TestLookupNsInstIdUseCachedData(t *testing.T) {
	nbic := &Session{
		nsInstMap: map[string][]string{"silly_ns": {"324567"}},
	}
	id, err := nbic.lookupNsInstanceId("silly_ns")
	if err != nil {
		t.Errorf("want: 324567; got: %v", err)
	}
	if *id != "324567" {
		t.Errorf("want: 324567; got: %s", *id)
	}
}

func TestLookupNsInstIdNilOnMiss(t *testing.T) {
	nbic := &Session{
		nsInstMap: map[string][]string{"silly_ns": {"324567"}},
	}
	if got, err := nbic.lookupNsInstanceId("not there!"); err != nil {
		t.Errorf("want: nil; got: %v", err)
	} else {
		if got != nil {
			t.Errorf("want: nil; got: %v", *got)
		}
	}
}

func TestLookupNsInstIdFetchDataFromServer(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	wantId := "0335c32c-d28c-4d79-9b94-0ffa36326932"
	if gotId, err := nbic.lookupNsInstanceId("ldap"); err != nil {
		t.Errorf("want: %s; got: %v", wantId, err)
	} else {
		if *gotId != wantId {
			t.Errorf("want: %s; got: %v", wantId, *gotId)
		}
	}

	if len(nbi.exchanges) != 2 {
		t.Fatalf("want: 2; got: %d", len(nbi.exchanges))
	}
	rr1, rr2 := nbi.exchanges[0], nbi.exchanges[1]
	if rr1.req.URL.Path != urls.Tokens().Path {
		t.Errorf("want: %s; got: %s", urls.Tokens().Path, rr1.req.URL.Path)
	}
	if rr2.req.URL.Path != urls.NsInstancesContent().Path {
		t.Errorf("want: %s; got: %s", urls.NsInstancesContent().Path, rr2.req.URL.Path)
	}
}

func TestLookupNsInstIdFetchDataFromServerTokenError(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, UserCredentials{}, nbi.exchange)

	if _, err := nbic.lookupNsInstanceId("ldap"); err == nil {
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

func TestLookupNsInstIdFetchDataFromServerDupNameError(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	if _, err := nbic.lookupNsInstanceId("dup-name"); err == nil {
		t.Errorf("want: error; got: nil")
	}

	if len(nbi.exchanges) != 2 {
		t.Fatalf("want: 2; got: %d", len(nbi.exchanges))
	}
	rr1, rr2 := nbi.exchanges[0], nbi.exchanges[1]
	if rr1.req.URL.Path != urls.Tokens().Path {
		t.Errorf("want: %s; got: %s", urls.Tokens().Path, rr1.req.URL.Path)
	}
	if rr2.req.URL.Path != urls.NsInstancesContent().Path {
		t.Errorf("want: %s; got: %s", urls.NsInstancesContent().Path, rr2.req.URL.Path)
	}
}

func TestCreateOrUpdateNsInstanceErrorOnNilData(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	if got := nbic.CreateOrUpdateNsInstance(nil); got == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestCreateNsInstanceErrorOnMissingNsd(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	data := NsInstanceContent{
		Name:           "not-there",
		Description:    "wada wada",
		NsdName:        "not there!",
		VimAccountName: "mylocation1",
	}
	if err := nbic.CreateOrUpdateNsInstance(&data); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestCreateNsInstanceErrorOnMissingVimAccount(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	data := NsInstanceContent{
		Name:           "not-there",
		Description:    "wada wada",
		NsdName:        "openldap_ns",
		VimAccountName: "not there!",
	}
	if err := nbic.CreateOrUpdateNsInstance(&data); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestCreateNsInstanceErrorOnDupNsInstanceName(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	data := NsInstanceContent{
		Name:           "dup-name",
		Description:    "wada wada",
		NsdName:        "openldap_ns",
		VimAccountName: "mylocation1",
	}
	if err := nbic.CreateOrUpdateNsInstance(&data); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func assertCreateNsInstanceHttpFlow(t *testing.T, urls Connection,
	flow []requestReply) string {
	if len(flow) != 5 {
		t.Fatalf("want: 5; got: %d", len(flow))
	}
	rr1, rr2, rr3, rr4, rr5 := flow[0], flow[1], flow[2], flow[3], flow[4]
	if rr1.req.URL.Path != urls.Tokens().Path {
		t.Errorf("want: %s; got: %s", urls.Tokens().Path, rr1.req.URL.Path)
	}
	if rr2.req.URL.Path != urls.NsInstancesContent().Path {
		t.Errorf("want: %s; got: %s", urls.NsInstancesContent().Path, rr2.req.URL.Path)
	}
	if rr3.req.URL.Path != urls.NsDescriptors().Path {
		t.Errorf("want: %s; got: %s", urls.NsDescriptors().Path, rr3.req.URL.Path)
	}
	if rr4.req.URL.Path != urls.VimAccounts().Path {
		t.Errorf("want: %s; got: %s", urls.VimAccounts().Path, rr4.req.URL.Path)
	}
	if rr5.req.URL.Path != urls.NsInstancesContent().Path {
		t.Errorf("want: %s; got: %s", urls.NsInstancesContent().Path, rr5.req.URL.Path)
	}
	if rr5.req.Method != "POST" {
		t.Errorf("want: POST; got: %s", rr5.req.Method)
	}
	got, err := ioutil.ReadAll(rr5.req.Body)
	if err != nil {
		t.Errorf("want: body; got: %v", err)
		return ""
	}
	return string(got)
}

func TestCreateNsInstanceWithNoAdditionalParams(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	data := NsInstanceContent{
		Name:           "not-there",
		Description:    "wada wada",
		NsdName:        "openldap_ns",
		VimAccountName: "mylocation1",
	}
	if err := nbic.CreateOrUpdateNsInstance(&data); err != nil {
		t.Errorf("want: create; got: %v", err)
	}

	want := `{"nsName":"not-there","nsdId":"aba58e40-d65f-4f4e-be0a-e248c14d3e03","nsDescription":"wada wada","vimAccountId":"4a4425f7-3e72-4d45-a4ec-4241186f3547"}`
	got := assertCreateNsInstanceHttpFlow(t, urls, nbi.exchanges)
	if got != want {
		t.Errorf("want: %s; got: %s", want, got)
	}
}

type kduYamlParams struct {
	Params interface{} `yaml:"params"`
}

var yamlParamsData = []byte(`---
params:
  replicaCount: "2"
`)

func TestCreateNsInstanceWithAdditionalParams(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	kdu := kduYamlParams{}
	if err := yaml.Unmarshal(yamlParamsData, &kdu); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	data := NsInstanceContent{
		Name:           "not-there",
		Description:    "wada wada",
		NsdName:        "openldap_ns",
		VimAccountName: "mylocation1",
		VnfName:        "openldap",
		KduName:        "ldap",
		KduParams:      kdu.Params,
	}
	if err := nbic.CreateOrUpdateNsInstance(&data); err != nil {
		t.Errorf("want: create; got: %v", err)
	}

	want := `{"nsName":"not-there","nsdId":"aba58e40-d65f-4f4e-be0a-e248c14d3e03","nsDescription":"wada wada","vimAccountId":"4a4425f7-3e72-4d45-a4ec-4241186f3547"`
	want += `,"additionalParamsForVnf":[{"member-vnf-index":"openldap","additionalParamsForKdu":[{"kdu_name":"ldap","additionalParams":{"replicaCount":"2"}}]}]}`
	got := assertCreateNsInstanceHttpFlow(t, urls, nbi.exchanges)
	if got != want {
		t.Errorf("want: %s; got: %s", want, got)
	}
}

func assertUpdateNsInstanceHttpFlow(t *testing.T, urls Connection,
	nsInstanceId string, flow []requestReply) string {
	if len(flow) != 3 {
		t.Fatalf("want: 3; got: %d", len(flow))
	}
	rr1, rr2, rr3 := flow[0], flow[1], flow[2]
	if rr1.req.URL.Path != urls.Tokens().Path {
		t.Errorf("want: %s; got: %s", urls.Tokens().Path, rr1.req.URL.Path)
	}
	if rr2.req.URL.Path != urls.NsInstancesContent().Path {
		t.Errorf("want: %s; got: %s", urls.NsInstancesContent().Path, rr2.req.URL.Path)
	}
	if rr3.req.URL.Path != urls.NsInstancesAction(nsInstanceId).Path {
		t.Errorf("want: %s; got: %s", urls.NsInstancesAction(nsInstanceId).Path, rr3.req.URL.Path)
	}
	if rr3.req.Method != "POST" {
		t.Errorf("want: POST; got: %s", rr3.req.Method)
	}
	got, err := ioutil.ReadAll(rr3.req.Body)
	if err != nil {
		t.Errorf("want: body; got: %v", err)
		return ""
	}
	return string(got)
}

func TestUpdateNsInstance(t *testing.T) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)

	nsInstanceId := "0335c32c-d28c-4d79-9b94-0ffa36326932"

	kdu := kduYamlParams{}
	if err := yaml.Unmarshal(yamlParamsData, &kdu); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	data := NsInstanceContent{
		Name:           "ldap",
		Description:    "wada wada",
		NsdName:        "openldap_ns",
		VimAccountName: "mylocation1",
		VnfName:        "openldap",
		KduName:        "ldap",
		KduParams:      kdu.Params,
	}
	if err := nbic.CreateOrUpdateNsInstance(&data); err != nil {
		t.Errorf("want: update; got: %v", err)
	}

	want := `{"member_vnf_index":"openldap","kdu_name":"ldap","primitive":"upgrade","primitive_params":{"replicaCount":"2"}}`
	got := assertUpdateNsInstanceHttpFlow(t, urls, nsInstanceId, nbi.exchanges)
	if got != want {
		t.Errorf("want: %s; got: %s", want, got)
	}
}
