package nbic

import (
	"fmt"

	u "github.com/martel-innovate/osmops/osmops/util"
)

type nsInstanceView struct { // only the response fields we care about.
	Id   string `json:"_id"`
	Name string `json:"name"`
}

type nsInstanceMap map[string][]string

func (m nsInstanceMap) addMapping(name string, id string) {
	if entry, ok := m[name]; ok {
		m[name] = append(entry, id)
	} else {
		m[name] = []string{id}
	}
}

func buildNsInstanceMap(vs []nsInstanceView) nsInstanceMap {
	nsMap := nsInstanceMap{}
	for _, v := range vs {
		nsMap.addMapping(v.Name, v.Id)
	}
	return nsMap
}

// NOTE. NS instance name to ID lookup.
// OSM NBI doesn't enforce uniqueness of NS names. In fact, it lets you happily
// create a new instance even if an existing one has the same name, e.g.
//
// $ curl localhost/osm/nslcm/v1/ns_instances_content \
// -v -X POST \
// -H 'Authorization: Bearer 0WhgBufy1Wt82NbF9OsmftwpRfcsV4sU' \
// -H 'Content-Type: application/yaml' \
// -d'{"nsdId": "aba58e40-d65f-4f4e-be0a-e248c14d3e03", "nsName": "ldap", "nsDescription": "default description", "vimAccountId": "4a4425f7-3e72-4d45-a4ec-4241186f3547"}'
// ...
// HTTP/1.1 201 Created
// ...
// ---
// id: 794ef9a2-8bbb-42c1-869a-bab6422982ec
// nslcmop_id: 0fdfaa6a-b742-480c-9701-122b3f732e4
//
// This is why we map an NS instance name to a list of IDs.

func (c *Session) getNsInstancesContent() ([]nsInstanceView, error) {
	data := []nsInstanceView{}
	if _, err := c.getJson(c.conn.NsInstancesContent(), &data); err != nil {
		return nil, err
	}
	return data, nil
}

type maybeNsInstId *string

func (c *Session) lookupNsInstanceId(name string) (maybeNsInstId, error) {
	if c.nsInstMap == nil {
		if vs, err := c.getNsInstancesContent(); err != nil {
			return nil, err
		} else {
			c.nsInstMap = buildNsInstanceMap(vs)
		}
	}
	if ids, ok := c.nsInstMap[name]; !ok {
		return nil, nil
	} else {
		if len(ids) != 1 {
			return nil,
				fmt.Errorf("NS instance name not bound to a single ID: %v", ids)
		}
		return &ids[0], nil
	}
}

// NsInstanceContent holds the data to create or update an NS instance.
// For now we only support creating or updating KNFs. For a create or
// update operation to work, the target KNF must've been "on-boarded"
// in OSM already. So there must be, in OSM, a NSD and VNFD for it.
type NsInstanceContent struct {
	// The name of the target NS instance to create or update.
	Name string
	// Short description of the NS instance.
	Description string
	// The name of the NSD that defines the NS instance.
	NsdName string
	// The name of the VIM account to use for creating/updating the NS instance.
	VimAccountName string
	// The name of the VNF to use when updating the NS instance.
	VnfName string
	// The name of the KDU for the NS as specified in the VNFD.
	KduName string
	// Any KNF-specific parameters to create or update the NS instance.
	KduParams interface{}
}

type nsInstContentDto struct {
	NsName                 string                      `json:"nsName"`
	NsdId                  string                      `json:"nsdId"`
	NsDescription          string                      `json:"nsDescription"`
	VimAccountId           string                      `json:"vimAccountId"`
	AdditionalParamsForVnf []additionalParamsForVnfDto `json:"additionalParamsForVnf,omitempty"`
}

type additionalParamsForVnfDto struct {
	MemberVnfIndex         string                      `json:"member-vnf-index"`
	AdditionalParamsForKdu []additionalParamsForKduDto `json:"additionalParamsForKdu"`
}

type additionalParamsForKduDto struct {
	KduName          string      `json:"kdu_name"`
	AdditionalParams interface{} `json:"additionalParams"`
}

type nsInstanceContentActionDto struct {
	MemberVnfIndex  string      `json:"member_vnf_index"`
	KduName         string      `json:"kdu_name"`
	Primitive       string      `json:"primitive"`
	PrimitiveParams interface{} `json:"primitive_params"`
}

func (c *Session) CreateOrUpdateNsInstance(data *NsInstanceContent) error {
	if data == nil {
		return fmt.Errorf("nil data")
	}

	nsId, err := c.lookupNsInstanceId(data.Name)
	if err != nil {
		return err
	}
	if nsId == nil {
		return c.createNsInstance(data)
	}
	return c.updateNsInstance(*nsId, data)
}

func toNsInstContentDto(nsdId string, vimAccId string,
	data *NsInstanceContent) *nsInstContentDto {
	dto := nsInstContentDto{
		NsName:        data.Name,
		NsdId:         nsdId,
		NsDescription: data.Description,
		VimAccountId:  vimAccId,
	}
	if data.KduParams != nil {
		dto.AdditionalParamsForVnf = []additionalParamsForVnfDto{
			{
				MemberVnfIndex: data.VnfName,
				AdditionalParamsForKdu: []additionalParamsForKduDto{
					{
						KduName:          data.KduName,
						AdditionalParams: data.KduParams,
					},
				},
			},
		}
	}
	return &dto
}

func (c *Session) createNsInstance(data *NsInstanceContent) error {
	nsdId, err := c.lookupNsDescriptorId(data.NsdName)
	if err != nil {
		return err
	}
	vimAccId, err := c.lookupVimAccountId(data.VimAccountName)
	if err != nil {
		return err
	}
	dto := toNsInstContentDto(nsdId, vimAccId, data)

	_, err = c.postJson(c.conn.NsInstancesContent(), dto)
	return err
}

var nsAction = struct {
	u.StrEnum
	CREATE, UPGRADE, DELETE u.EnumIx
}{
	StrEnum: u.NewStrEnum("create", "upgrade", "delete"),
	CREATE:  0,
	UPGRADE: 1,
	DELETE:  2,
}

func toNsInstanceContentActionDto(nsId string, data *NsInstanceContent) *nsInstanceContentActionDto {
	return &nsInstanceContentActionDto{
		MemberVnfIndex:  data.VnfName,
		KduName:         data.KduName,
		Primitive:       nsAction.LabelOf(nsAction.UPGRADE),
		PrimitiveParams: data.KduParams,
	}
}

func (c *Session) updateNsInstance(nsId string, data *NsInstanceContent) error {
	dto := toNsInstanceContentActionDto(nsId, data)
	_, err := c.postJson(c.conn.NsInstancesAction(nsId), dto)
	return err
}
