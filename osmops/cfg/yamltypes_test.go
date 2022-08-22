package cfg

import (
	"testing"
)

var opsConfigValidationFailFixtures = []OpsConfig{
	{TargetDir: "", ConnectionFile: ""},
	{TargetDir: " ", ConnectionFile: "\n"},
	{TargetDir: "valid", ConnectionFile: "\n"},
	{TargetDir: "\t", ConnectionFile: "./val/id"},
}

func TestOpsConfigValidationFail(t *testing.T) {
	for k, d := range opsConfigValidationFailFixtures {
		if got := d.Validate(); got == nil {
			t.Errorf("[%d] want: error; got: valid", k)
		}
	}
}

var opsConfigValidationOkFixtures = []OpsConfig{
	{ConnectionFile: "./"},
	{TargetDir: "", ConnectionFile: "./"},
	{TargetDir: ".", ConnectionFile: "./"},
	{TargetDir: " /a/", ConnectionFile: "/a/b "},
	{TargetDir: "valid", ConnectionFile: "./val/id"},
	{TargetDir: "\tval/id\n", ConnectionFile: "\n/val/id/\t"},
}

func TestOpsConfigValidationOk(t *testing.T) {
	for k, d := range opsConfigValidationOkFixtures {
		if got := d.Validate(); got != nil {
			t.Errorf("[%d] want: valid; got: %s", k, got)
		}
	}
}

var osmConnectionValidationFailFixtures = []OsmConnection{
	{Hostname: "", User: "u", Password: "p"},
	{}, {Hostname: "h", Password: "p"}, {Hostname: "h:80", Password: "p"},
	{Hostname: "h:20", User: "u", Password: "p"},
	{Hostname: "h:20", User: "u", Project: "p"},
}

func TestOsmConnectionValidationFail(t *testing.T) {
	for k, d := range osmConnectionValidationFailFixtures {
		if got := d.Validate(); got == nil {
			t.Errorf("[%d] want: error; got: valid", k)
		}
	}
}

var osmConnectionValidationOkFixtures = []OsmConnection{
	{Hostname: "h:0", Project: "p", User: "u", Password: "p"},
	{Hostname: "h:1", Project: "p", User: "u", Password: "*"},
}

func TestOsmConnectionValidationOk(t *testing.T) {
	for k, d := range osmConnectionValidationOkFixtures {
		if got := d.Validate(); got != nil {
			t.Errorf("[%d] want: valid; got: %s", k, got)
		}
	}
}

var kduNsActionValidationFailFixtures = []KduNsAction{
	{},
	{
		Kind:    "NsInstance",
		Name:    "x",
		NsdName: "x",
		VnfName: "x",
	},
	{
		Kind:           "NSinsTance",
		Name:           "x",
		VnfName:        "x",
		VimAccountName: "x",
		Kdu: Kdu{
			Params: "x",
		},
	},
	{
		Name:           "x",
		NsdName:        "x",
		VnfName:        "x",
		VimAccountName: "x",
		Kdu: Kdu{
			Name:   "x",
			Params: "x",
		},
	},
	{
		Kind:           "ain't right",
		Name:           "x",
		Description:    "wada wada",
		NsdName:        "x",
		VnfName:        "x",
		VimAccountName: "x",
		Kdu: Kdu{
			Name:   "x",
			Params: "x",
		},
	},
	{
		Kind:    "nsinstance",
		VnfName: "x",
		Kdu: Kdu{
			Name:   "x",
			Params: "x",
		},
	},
	{
		Kind:    "NsInstance",
		Name:    "x",
		VnfName: "x",
		Kdu: Kdu{
			Name:   "x",
			Params: "x",
		},
	},
	{
		Kind:           "NsInstance",
		Name:           "x",
		VnfName:        "x",
		VimAccountName: "x",
		Kdu: Kdu{
			Name:   "x",
			Params: "x",
		},
	},
	{
		Kind: "NsInstance",
		Name: "x",
		Kdu: Kdu{
			Name:   "x",
			Params: "x",
		},
	},
}

func TestKduNsActionValidationFail(t *testing.T) {
	for k, d := range kduNsActionValidationFailFixtures {
		if got := d.Validate(); got == nil {
			t.Errorf("[%d] want: error; got: valid", k)
		}
	}
}

var kduNsActionValidationOkFixtures = []KduNsAction{
	{
		Kind:           "nsiNStance",
		Name:           "x",
		NsdName:        "x",
		VnfName:        "x",
		VimAccountName: "x",
		Kdu: Kdu{
			Name:   "x",
			Params: "x",
		},
	},
	{
		Kind:           "NsInstance",
		Name:           "x",
		Description:    "wada wada",
		NsdName:        "x",
		VnfName:        "x",
		VimAccountName: "x",
		Kdu: Kdu{
			Name: "x",
		},
	},
	{
		Kind:           "nsinstance",
		Name:           "x",
		NsdName:        "x",
		VnfName:        "x",
		VimAccountName: "x",
		Kdu: Kdu{
			Name: "x",
		},
	},
}

func TestKduNsActionValidationOk(t *testing.T) {
	for k, d := range kduNsActionValidationOkFixtures {
		if got := d.Validate(); got != nil {
			t.Errorf("[%d] want: valid; got: %s", k, got)
		}
	}
}
