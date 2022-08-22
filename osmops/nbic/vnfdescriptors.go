package nbic

import (
	"fmt"
)

type vnfDescView struct { // only the response fields we care about.
	Id   string `json:"_id"`
	Name string `json:"id"`
}

type vnfDescMap map[string]string

func buildVnfDescMap(ds []vnfDescView) vnfDescMap {
	descMap := map[string]string{}
	for _, d := range ds {
		descMap[d.Name] = d.Id
	}
	return descMap
}

// NOTE. VNFD name to ID lookup.
// For our vnfDescMap to work, there must be a bijection between VNFD IDs and
// name IDs. Luckily, this is the case since OSM NBI enforces uniqueness of
// VNFD name IDs. If you try uploading another package with a VNFD having the
// same name ID of an existing one, OSM NBI will complain loudly, e.g.
//
// HTTP/1.1 409 Conflict
// ...
// {
//     "code": "CONFLICT",
//     "status": 409,
//     "detail": "vnfd with id 'openldap_knf' already exists for this project"
// }

func (c *Session) getVnfDescriptors() ([]vnfDescView, error) {
	data := []vnfDescView{}
	_, err := c.getJson(c.conn.VnfPackagesContent(), &data)
	return data, err
}

func (c *Session) lookupVnfDescriptorId(name string) (string, error) {
	if c.vnfdMap == nil {
		if ds, err := c.getVnfDescriptors(); err != nil {
			return "", err
		} else {
			c.vnfdMap = buildVnfDescMap(ds)
		}
	}
	if id, ok := c.vnfdMap[name]; !ok {
		return "", &missingDescriptor{typ: "VNFD", name: name}
	} else {
		return id, nil
	}
}

type missingDescriptor struct {
	typ  string
	name string
}

func (e *missingDescriptor) Error() string {
	return fmt.Sprintf("no %s found for name ID: %s", e.typ, e.name)
}
