package nbic

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/martel-innovate/osmops/osmops/pkgr"
	"github.com/martel-innovate/osmops/osmops/util/file"

	//lint:ignore ST1001 HTTP EDSL is more readable w/o qualified import
	. "github.com/martel-innovate/osmops/osmops/util/http"
)

func (s *Session) CreateOrUpdatePackage(source file.AbsPath) error {
	handler, err := newPkgHandler(s, source)
	if err != nil {
		return err
	}
	return handler.process()
}

// pkgReader wraps Package to consolidate in one place all the assumptions
// this module makes about OSM packages in an OsmOps-managed repo.
// Specifically:
//
// - pkg name = pkg ID
// - VNF pkg => pgk name ends w/ "_knf"
// - NS pkg => pkg name ends w/ "_ns"
//
// None of the above needs to be true in general, but OsmOps relies on that
// at the moment to simplify the implementation. Eventually, we'll redo this
// properly, i.e. use a semantic approach (parse, interpret OSM files) rather
// than naming conventions and guesswork.
type pkgReader struct {
	pkg  *pkgr.Package
	data []byte
}

func newPkgReader(pkgSource file.AbsPath) (*pkgReader, error) {
	pkg, err := pkgr.Pack(pkgSource)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(pkg.Data)
	return &pkgReader{
		pkg:  pkg,
		data: data,
	}, err
}

func (r *pkgReader) Source() file.AbsPath {
	return r.pkg.Source.Directory()
}

func (r *pkgReader) Name() string {
	return r.pkg.Name
}

func (r *pkgReader) Id() string {
	return r.Name()
}

func (r *pkgReader) Data() []byte {
	return r.data
}

func (r *pkgReader) Hash() string {
	return r.pkg.Hash
}

func (r *pkgReader) IsNs() bool {
	return strings.HasSuffix(r.Name(), "_ns")
}

func (r *pkgReader) IsKnf() bool {
	return strings.HasSuffix(r.Name(), "_knf")
}

type pkgHandler struct {
	session  *Session
	pkg      *pkgReader
	endpoint *url.URL
	isUpdate bool
}

func newPkgHandler(sesh *Session, pkgSrc file.AbsPath) (*pkgHandler, error) {
	reader, err := newPkgReader(pkgSrc)
	if err != nil {
		return nil, err
	}
	handler := &pkgHandler{
		session: sesh,
		pkg:     reader,
	}
	if reader.IsKnf() {
		return mkPkgHandler(
			handler, handler.session.lookupVnfDescriptorId,
			handler.session.conn.VnfPackagesContent,
			handler.session.conn.VnfPackageContent)
	}
	if reader.IsNs() {
		return mkPkgHandler(
			handler, handler.session.lookupNsDescriptorId,
			handler.session.conn.NsPackagesContent,
			handler.session.conn.NsPackageContent)
	}
	return nil, unsupportedPackageType(reader)
}

func unsupportedPackageType(pkg *pkgReader) error {
	return fmt.Errorf("unsupported package type: %v", pkg.Source())
}

type lookupDescId func(pkgId string) (string, error)
type createEndpoint func() *url.URL
type updateEndpoint func(osmPkgId string) *url.URL

func mkPkgHandler(h *pkgHandler, getOsmId lookupDescId,
	createUrl createEndpoint, updateUrl updateEndpoint) (*pkgHandler, error) {
	osmPkgId, err := getOsmId(h.pkg.Id())
	if _, ok := err.(*missingDescriptor); ok {
		h.isUpdate = false
		h.endpoint = createUrl()

		return h, nil
	}
	if err == nil {
		h.isUpdate = true
		h.endpoint = updateUrl(osmPkgId)
	}
	return h, err
}

func (h *pkgHandler) process() error {
	run := h.post
	if h.isUpdate {
		run = h.put
	}
	_, err := run()
	return err
}

func (h *pkgHandler) post() (*http.Response, error) {
	req := Request(
		POST, At(h.endpoint),
		h.session.NbiAccessToken(),
		Accept(MediaType.JSON),  // same as what OSM client does
		Content(MediaType.GZIP), // ditto
		ContentFilename(h.pkg),  // ditto
		ContentFileMd5(h.pkg),   // ditto
		Body(h.pkg.Data()),
	)
	req.SetHandler(ExpectSuccess())
	return req.RunWith(h.session.transport)
}

func (h *pkgHandler) put() (*http.Response, error) {
	descData, err := h.findPkgDescriptor()
	if err != nil {
		return nil, err
	}
	req := Request(
		PUT, At(h.endpoint),
		h.session.NbiAccessToken(),
		Accept(MediaType.JSON),
		Content(MediaType.YAML),
		Body(descData),
	)
	req.SetHandler(ExpectSuccess())
	return req.RunWith(h.session.transport)
}

// NOTE. Package update. It's kinda weird the way it works, but most likely
// I'm missing something. In fact, our initial implementation of put uploaded
// the tarball. As it turns out, OSM client does something different. It tries
// finding a YAML file in the package dir, blindly assumes it's a VNFD or NSD
// and PUTs it in OSM. What if there are other files in the package? Well,
// I've got no idea why OSM client does that, but I've changed put's impl
// to be in line with OSM client's.
// OSM client's update methods:
// - https://osm.etsi.org/gitlab/osm/osmclient/-/blob/master/osmclient/sol005/vnfd.py
// - https://osm.etsi.org/gitlab/osm/osmclient/-/blob/master/osmclient/sol005/nsd.py

func (h *pkgHandler) findPkgDescriptor() ([]byte, error) {
	candidates := []string{}
	for _, archivePath := range h.pkg.pkg.Source.SortedFilePaths() {
		p := strings.ToLower(archivePath)
		if strings.HasSuffix(p, ".yaml") || strings.HasSuffix(p, ".yml") {
			candidates = append(candidates, archivePath)
		}
	}
	if len(candidates) == 0 { // same as what OSM client does
		return []byte{}, noDescriptorFound(h.pkg)
	}
	if len(candidates) > 1 { // same as what OSM client does
		return []byte{}, moreThanOneDescriptorFound(h.pkg)
	}
	return h.pkg.pkg.Source.FileContent(candidates[0])
}

func moreThanOneDescriptorFound(pkg *pkgReader) error {
	msg := "found more than one potential descriptor in: %v"
	return fmt.Errorf(msg, pkg.Source())
}

func noDescriptorFound(pkg *pkgReader) error {
	msg := "no descriptor found in: %v"
	return fmt.Errorf(msg, pkg.Source())
}

func ContentFilename(pkg *pkgReader) ReqBuilder {
	name := fmt.Sprintf("%s.tar.gz", pkg.Name())
	return func(request *http.Request) error {
		request.Header.Set("Content-Filename", name)
		return nil
	}
}

func ContentFileMd5(pkg *pkgReader) ReqBuilder {
	return func(request *http.Request) error {
		request.Header.Set("Content-File-MD5", pkg.Hash())
		return nil
	}
}
