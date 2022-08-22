package nbic

import (
	"fmt"
	"net/url"

	"github.com/martel-innovate/osmops/osmops/util"
)

// Connection holds the data needed to establish a network connection with
// the OSM NBI.
type Connection struct {
	Address util.HostAndPort
	Secure  bool
}

func (b Connection) buildUrl(path string) *url.URL {
	if url, err := b.Address.BuildHttpUrl(b.Secure, path); err != nil {
		panic(err) // see note below
	} else {
		return url
	}
}

// NOTE. Panic on URL building.
// Ideally buildUrl should return (*url.URL, error) instead of panicing. But
// then it becomes a royal pain in the backside to write code that uses the
// URL functions below and testing for the URL build error case needs to
// happen at every calling site---e.g. if you call Tokens then ideally there
// should be a unit test to check what happens when Tokens returns an error.
// So we take a shortcut with the panic call. As long as we call all the
// URL functions below in our unit tests, we can be sure the panic won't
// happen at runtime.

// Tokens returns the URL to the NBI tokens endpoint.
func (b Connection) Tokens() *url.URL {
	return b.buildUrl("/osm/admin/v1/tokens")
}

// NsDescriptors returns the URL to the NBI NS descriptors endpoint.
func (b Connection) NsDescriptors() *url.URL {
	return b.buildUrl("/osm/nsd/v1/ns_descriptors")
}

// VimAccounts returns the URL to the VIM accounts endpoint.
func (b Connection) VimAccounts() *url.URL {
	return b.buildUrl("/osm/admin/v1/vim_accounts")
}

// NsInstances returns the URL to the NS instances content endpoint.
func (b Connection) NsInstancesContent() *url.URL {
	return b.buildUrl("/osm/nslcm/v1/ns_instances_content")
}

// NsInstancesAction returns the URL to the NS instances action endpoint
// for the NS instance identified by the given ID.
func (b Connection) NsInstancesAction(nsInstanceId string) *url.URL {
	path := fmt.Sprintf("/osm/nslcm/v1/ns_instances/%s/action", nsInstanceId)
	return b.buildUrl(path)
}

// VnfPackagesContent returns the URL to the VNF packages content endpoint.
func (b Connection) VnfPackagesContent() *url.URL {
	return b.buildUrl("/osm/vnfpkgm/v1/vnf_packages_content")
}

// VnfPackageContent returns the URL to the endpoint of the VNF package
// content identified by the given ID.
func (b Connection) VnfPackageContent(pkgId string) *url.URL {
	path := fmt.Sprintf("/osm/vnfpkgm/v1/vnf_packages_content/%s", pkgId)
	return b.buildUrl(path)
}

// NsPackagesContent returns the URL to the NS packages content endpoint.
func (b Connection) NsPackagesContent() *url.URL {
	return b.buildUrl("/osm/nsd/v1/ns_descriptors_content")
}

// NsPackageContent returns the URL to the endpoint of the NS package
// content identified by the given ID.
func (b Connection) NsPackageContent(pkgId string) *url.URL {
	path := fmt.Sprintf("/osm/nsd/v1/ns_descriptors_content/%s", pkgId)
	return b.buildUrl(path)
}
