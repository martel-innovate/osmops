package cfg

import (
	"reflect"
	"testing"
)

func TestFromBytesErrorOnInvalidYaml(t *testing.T) {
	data := []byte(`x: { y`)
	if _, err := readOpsConfig(data); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestReadOpsConfig(t *testing.T) {
	data := `
targetDir: deploy/ment
fileExtensions:
  - .x
  - .ya.ml
connectionFile: /the/secret/stash.yaml
`
	want := &OpsConfig{
		TargetDir:      "deploy/ment",
		FileExtensions: []string{".x", ".ya.ml"},
		ConnectionFile: "/the/secret/stash.yaml",
	}

	got, err := readOpsConfig([]byte(data))
	if err != nil {
		t.Errorf("failed to read config object: %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestReadOpsConfigMissingTargetDir(t *testing.T) {
	data := `
fileExtensions:
  - .x
connectionFile: /the/secret/stash.yaml
`
	want := &OpsConfig{
		FileExtensions: []string{".x"},
		ConnectionFile: "/the/secret/stash.yaml",
	}

	got, err := readOpsConfig([]byte(data))
	if err != nil {
		t.Errorf("failed to read config object: %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestReadOpsConfigMissingFileExtensions(t *testing.T) {
	data := `
connectionFile: /the/secret/stash.yaml
`
	want := &OpsConfig{
		ConnectionFile: "/the/secret/stash.yaml",
	}

	got, err := readOpsConfig([]byte(data))
	if err != nil {
		t.Errorf("failed to read config object: %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestReadOpsConfigMissingConnectionFile(t *testing.T) {
	data := `
targetDir: deploy/ment
`
	got, err := readOpsConfig([]byte(data))
	if err == nil {
		t.Errorf("want: validation fail; got: %v", got)
	}
}

func TestReadOsmConnection(t *testing.T) {
	data := `
hostname: osm.dev:8008
project: pea
user: silly-billy
password: "yo! "
`
	want := &OsmConnection{
		Hostname: "osm.dev:8008",
		Project:  "pea",
		User:     "silly-billy",
		Password: "yo! ",
	}

	got, err := readOsmConnection([]byte(data))
	if err != nil {
		t.Errorf("failed to read config object: %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestReadInvalidOsmConnection(t *testing.T) {
	data := `
hostname: missing.port
user: silly-billy
password: "yo! "
`
	got, err := readOsmConnection([]byte(data))
	if err == nil {
		t.Errorf("want: validation fail; got: %v", got)
	}
}

func TestReadKduNsAction(t *testing.T) {
	data := `
kind: NsInstance
name: silly billy
description: look ma!
nsdName: nascar
vnfName: WTH
vimAccountName: emacs rocks
kdu:
  name: kudu buck
  params:
    p: 1
    q: 2
`
	want := &KduNsAction{
		Kind:           "NsInstance",
		Name:           "silly billy",
		Description:    "look ma!",
		NsdName:        "nascar",
		VnfName:        "WTH",
		VimAccountName: "emacs rocks",
		Kdu: Kdu{
			Name: "kudu buck",
			Params: map[interface{}]interface{}{
				"p": 1,
				"q": 2,
			},
		},
	}

	got, err := readKduNsAction([]byte(data))
	if err != nil {
		t.Errorf("failed to read config object: %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestReadInvalidKduNsAction(t *testing.T) {
	data := `
kind: invalid!
name: silly billy
description: look ma!
nsdName: nascar
vnfName: WTH
vimAccountName: emacs rocks
kdu:
  name: kudu buck
  params:
    p: 1
`
	got, err := readKduNsAction([]byte(data))
	if err == nil {
		t.Errorf("want: validation fail; got: %v", got)
	}
}

func TestReadKduNsActionWithNoParams(t *testing.T) {
	data := `
kind: NsInstance
name: silly billy
description: look ma!
nsdName: nascar
vnfName: WTH
vimAccountName: emacs rocks
kdu:
  name: kudu buck
`
	got, err := readKduNsAction([]byte(data))
	if err != nil {
		t.Fatalf("want: data; got: %v", err)
	}
	if got.Kdu.Params != nil {
		t.Errorf("want: nil; got: %+v", got.Kdu.Params)
	}
}

func TestReadKduNsActionWithSingleParam(t *testing.T) {
	data := `
kind: NsInstance
name: t3
nsdName: d3
vnfName: f3
vimAccountName: v3
kdu:
  name: k3
  params:
    replicaCount: "3"

`
	got, err := readKduNsAction([]byte(data))
	if err != nil {
		t.Fatalf("want: data; got: %v", err)
	}

	ps, ok := got.Kdu.Params.(map[interface{}]interface{})
	if !ok {
		t.Fatalf("want: params map; got: %+v", got.Kdu.Params)
	}
	v, ok := ps["replicaCount"]
	if !ok {
		t.Fatalf("want: replicaCount value; got: not there")
	}
	sv, ok := v.(string)
	if !ok {
		t.Fatalf("want: string value; got: %v", v)
	}
	if sv != "3" {
		t.Errorf(`want: "3"; got: "%s"`, sv)
	}
}
