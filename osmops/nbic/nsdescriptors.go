package nbic

type nsDescView struct { // only the response fields we care about.
	Id   string `json:"_id"`
	Name string `json:"id"`
}

type nsDescMap map[string]string

func buildNsDescMap(ds []nsDescView) nsDescMap {
	descMap := map[string]string{}
	for _, d := range ds {
		descMap[d.Name] = d.Id
	}
	return descMap
}

// NOTE. NSD name to ID lookup.
// For our nsDescMap to work, there must be a bijection between NSD IDs and
// name IDs. Luckily, this is the case since OSM NBI enforces uniqueness of
// NSD name IDs. If you try uploading another package with a NSD having the
// same name ID of an existing one, OSM NBI will complain loudly, e.g.
//
// HTTP/1.1 409 Conflict
// ...
// {
//     "code": "CONFLICT",
//     "status": 409,
//     "detail": "nsd with id 'openldap_ns' already exists for this project"
// }

func (c *Session) getNsDescriptors() ([]nsDescView, error) {
	data := []nsDescView{}
	if _, err := c.getJson(c.conn.NsDescriptors(), &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Session) lookupNsDescriptorId(name string) (string, error) {
	if c.nsdMap == nil {
		if ds, err := c.getNsDescriptors(); err != nil {
			return "", err
		} else {
			c.nsdMap = buildNsDescMap(ds)
		}
	}
	if id, ok := c.nsdMap[name]; !ok {
		return "", &missingDescriptor{typ: "NSD", name: name}
	} else {
		return id, nil
	}
}
