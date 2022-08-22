package engine

import (
	"gopkg.in/yaml.v2"

	"github.com/martel-innovate/osmops/osmops/nbic"
	"github.com/martel-innovate/osmops/osmops/util"
)

type kduP struct {
	Params interface{} `yaml:"params"`
}

func kduParams() interface{} {
	yamlData := []byte(`---
params:
  replicaCount: "2"
`)

	kdu := kduP{}
	if err := yaml.Unmarshal(yamlData, &kdu); err != nil {
		panic(err)
	}
	return kdu.Params
}

func T() {
	hp, _ := util.ParseHostAndPort("192.168.64.19:80")
	conn := nbic.Connection{Address: *hp, Secure: false}
	usrCreds := nbic.UserCredentials{
		Username: "admin", Password: "admin", Project: "admin",
	}
	client, _ := nbic.New(conn, usrCreds)

	data := nbic.NsInstanceContent{
		Name:           "ldap3",
		Description:    "wada wada",
		NsdName:        "openldap_ns",
		VimAccountName: "mylocation1",
		VnfName:        "openldap",
		KduName:        "ldap",
		KduParams:      kduParams(),
	}
	err := client.CreateOrUpdateNsInstance(&data)
	if err != nil {
		panic(err)
	}
}
