package nbic

import (
	"fmt"
)

type vimAccountView struct { // only the response fields we care about.
	Id   string `json:"_id"`
	Name string `json:"name"`
}

type vimAccountMap map[string]string

func buildVimAccountMap(vs []vimAccountView) vimAccountMap {
	accountMap := map[string]string{}
	for _, v := range vs {
		accountMap[v.Name] = v.Id
	}
	return accountMap
}

// NOTE. VIM account name to ID lookup.
// For our vimAccountMap to work, there must be a bijection between VIM account
// IDs and names. Lucklily, this is the case since OSM NBI enforces uniqueness
// of VIM account names. If you try creating a VIM account with the same name
// as an existing one, you get an error, e.g.
//
// HTTP/1.1 409 Conflict
// ...
// ---
// code: CONFLICT
// detail: name 'openvim-site' already exists for vim_accounts
// status: 409

func (c *Session) getVimAccounts() ([]vimAccountView, error) {
	data := []vimAccountView{}
	if _, err := c.getJson(c.conn.VimAccounts(), &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Session) lookupVimAccountId(name string) (string, error) {
	if c.vimAccMap == nil {
		if vs, err := c.getVimAccounts(); err != nil {
			return "", err
		} else {
			c.vimAccMap = buildVimAccountMap(vs)
		}
	}
	if id, ok := c.vimAccMap[name]; !ok {
		return "", fmt.Errorf("no VIM account found for name ID: %s", name)
	} else {
		return id, nil
	}
}
