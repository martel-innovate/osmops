// Client to interact with OSM north-bound interface (NBI).
package nbic

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"github.com/martel-innovate/osmops/osmops/util/file"

	//lint:ignore ST1001 HTTP EDSL is more readable w/o qualified import
	. "github.com/martel-innovate/osmops/osmops/util/http"
	"github.com/martel-innovate/osmops/osmops/util/http/sec"
)

// Workflow defines functions to carry out high-level tasks, usually involving
// several NBI calls.
type Workflow interface {
	// CreateOrUpdateNsInstance creates or updates an NS instance in OSM
	// through NBI.
	//
	// If there's no instance with the specified name, then a new one gets
	// created. Otherwise, it's an update. Notice OSM allows duplicate instance
	// names (bug?), hence it's not safe to update an instance given it's
	// name---which instance to update if there's more than one with the
	// same name? So CreateOrUpdateNsInstance errors out if the given name
	// is tied to more than one instance.
	//
	// For now we only support creating or updating KNFs. For a create or
	// update operation to work, the target KNF must've been "on-boarded"
	// in OSM already. So there must be, in OSM, a NSD and VNFD for it.
	CreateOrUpdateNsInstance(data *NsInstanceContent) error

	// CreateOrUpdatePackage uploads the given package to OSM through NBI.
	//
	// CreateOrUpdatePackage blindly assumes that the given directory in
	// the OSMOps repo contains either a KNF or NS package. If the directory
	// name ends with "_knf", CreateOrUpdatePackage treats the whole directory
	// as a KNF package. Likewise, if the directory name ends with "_ns",
	// CreateOrUpdatePackage treats it as an NS package. (CreateOrUpdatePackage
	// will report an error if the directory name doesn't have an "_ns" or
	// "_knf" suffix.)
	//
	// CreateOrUpdatePackage also relies on another naming convention to
	// figure out the package ID. In fact, it assumes the directory name
	// is also the package ID declared in the KNF or NS YAML stanza.
	//
	// CreateOrUpdatePackage expects to find the source files of the OSM
	// package in the given source directory or subdirectories. It reads,
	// recursively, the files in source, creates a gzipped tar archive in
	// the OSM format (including creating the "checksums.txt" file) and
	// then streams it to OSM NBI to create or update the package in OSM.
	CreateOrUpdatePackage(source file.AbsPath) error
}

const REQUEST_TIMEOUT_SECONDS = 600

func newHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // (1)
			},
		},
		Timeout: time.Second * REQUEST_TIMEOUT_SECONDS, // (2)
	}
	// NOTE.
	// 1. Man-in-the-middle attacks. OSM client doesn't validate the server
	// cert, so we do the same. But this is a huge security loophole since it
	// opens the door to man-in-the-middle attacks.
	// 2. Request timeout. Always specify it, see
	// - https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
}

type Session struct {
	conn      Connection
	creds     UserCredentials
	transport ReqSender
	authz     *sec.TokenManager
	nsdMap    nsDescMap
	vnfdMap   vnfDescMap
	vimAccMap vimAccountMap
	nsInstMap nsInstanceMap
}

func New(conn Connection, creds UserCredentials, transport ...ReqSender) (
	*Session, error) {
	httpc := newHttpClient()

	agent := httpc.Do
	if len(transport) > 0 {
		agent = transport[0]
	}

	authz, err := NewAuthz(conn, creds, agent)
	if err != nil {
		return nil, err
	}

	return &Session{
		conn:      conn,
		creds:     creds,
		transport: agent,
		authz:     authz,
	}, nil
}

func (c *Session) NbiAccessToken() ReqBuilder {
	provider := func() (string, error) {
		if token, err := c.authz.GetAccessToken(); err != nil {
			return "", err
		} else {
			return token.String(), nil
		}
	}
	return BearerToken(provider)
}

func (c *Session) getJson(endpoint *url.URL, data interface{}) (
	*http.Response, error) {
	return Request(
		GET, At(endpoint),
		c.NbiAccessToken(),
		Accept(MediaType.JSON),
	).
		SetHandler(ExpectSuccess(), ReadJsonResponse(data)).
		RunWith(c.transport)
}

func (c *Session) postJson(endpoint *url.URL, inData interface{},
	outData ...interface{}) (*http.Response, error) {
	req := Request(
		POST, At(endpoint),
		c.NbiAccessToken(),
		Accept(MediaType.JSON),
		Content(MediaType.YAML), // same as what OSM client does
		JsonBody(inData),
	)
	if len(outData) > 0 {
		req.SetHandler(ExpectSuccess(), ReadJsonResponse(outData[0]))
	}
	return req.RunWith(c.transport)
}
