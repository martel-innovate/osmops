package nbic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	u "github.com/martel-innovate/osmops/osmops/util/http"
)

func stringReader(data string) io.ReadCloser {
	return io.NopCloser(strings.NewReader(data))
}

type requestReply struct {
	req *http.Request
	res *http.Response
}

type mockNbi struct {
	handlers  map[string]u.ReqSender
	exchanges []requestReply
	packages  map[string][]byte
}

func newMockNbi() *mockNbi {
	mock := &mockNbi{
		handlers:  map[string]u.ReqSender{},
		exchanges: []requestReply{},
		packages:  map[string][]byte{},
	}

	mock.handlers[handlerKey("POST", "/osm/admin/v1/tokens")] = tokenHandler
	mock.handlers[handlerKey("GET", "/osm/nsd/v1/ns_descriptors")] = nsDescHandler
	mock.handlers[handlerKey("GET", "/osm/admin/v1/vim_accounts")] = vimAccHandler
	mock.handlers[handlerKey("GET", "/osm/nslcm/v1/ns_instances_content")] = nsInstContentHandler
	mock.handlers[handlerKey("POST", "/osm/nslcm/v1/ns_instances_content")] = nsInstContentHandler
	mock.handlers[handlerKey("POST",
		"/osm/nslcm/v1/ns_instances/0335c32c-d28c-4d79-9b94-0ffa36326932/action")] = nsInstActionHandler
	mock.handlers[handlerKey("GET",
		"/osm/vnfpkgm/v1/vnf_packages_content")] = vnfDescHandler
	mock.handlers[handlerKey("POST",
		"/osm/vnfpkgm/v1/vnf_packages_content")] = mock.createPkgHandler
	mock.handlers[handlerKey("PUT",
		"/osm/vnfpkgm/v1/vnf_packages_content/")] = mock.updatePkgHandler
	mock.handlers[handlerKey("POST",
		"/osm/nsd/v1/ns_descriptors_content")] = mock.createPkgHandler
	mock.handlers[handlerKey("PUT",
		"/osm/nsd/v1/ns_descriptors_content/")] = mock.updatePkgHandler

	return mock
}

func handlerKey(method string, path string) string {
	return fmt.Sprintf("%s %s", method, path)
}

func (s *mockNbi) lookupHandler(req *http.Request) (u.ReqSender, error) {
	key := handlerKey(req.Method, req.URL.Path)
	if handle, ok := s.handlers[key]; ok {
		return handle, nil
	}
	for k, handle := range s.handlers {
		if strings.HasPrefix(key, k) {
			return handle, nil
		}
	}
	return nil, fmt.Errorf("no handler for request: %s", key)
}

func (s *mockNbi) exchange(req *http.Request) (*http.Response, error) {
	handle, err := s.lookupHandler(req)
	if err != nil {
		return &http.Response{StatusCode: http.StatusInternalServerError}, err
	}

	res, err := handle(req)
	rr := requestReply{req: req, res: res}
	s.exchanges = append(s.exchanges, rr)

	return res, err
}

func tokenHandler(req *http.Request) (*http.Response, error) {
	reqCreds := UserCredentials{}
	json.NewDecoder(req.Body).Decode(&reqCreds)
	if reqCreds.Password != usrCreds.Password {
		return &http.Response{StatusCode: http.StatusUnauthorized}, nil
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       stringReader(validNbiTokenPayload),
	}, nil
}

func vnfDescHandler(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       stringReader(vnfDescriptors),
	}, nil
}

func nsDescHandler(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       stringReader(nsDescriptors),
	}, nil
}

func vimAccHandler(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       stringReader(vimAccounts),
	}, nil
}

func nsInstContentHandler(req *http.Request) (*http.Response, error) {
	if req.Method == "GET" {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       stringReader(nsInstancesContent),
		}, nil
	}

	// POST
	return &http.Response{StatusCode: http.StatusCreated}, nil
}

func nsInstActionHandler(req *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusAccepted}, nil
}

func (m *mockNbi) createPkgHandler(req *http.Request) (*http.Response, error) {
	name := strings.TrimSuffix(req.Header.Get("Content-Filename"), ".tar.gz")
	if name == "" {
		return &http.Response{StatusCode: http.StatusBadRequest}, nil
	}

	if _, ok := m.packages[name]; ok {
		return &http.Response{StatusCode: http.StatusConflict}, nil
	}

	pkgTgzData, _ := io.ReadAll(req.Body)
	m.packages[name] = pkgTgzData
	return &http.Response{StatusCode: http.StatusCreated}, nil
}

func (m *mockNbi) updatePkgHandler(req *http.Request) (*http.Response, error) {
	osmPkgId := path.Base(req.URL.Path)
	pkgTgzData, _ := io.ReadAll(req.Body)
	m.packages[osmPkgId] = pkgTgzData
	return &http.Response{StatusCode: http.StatusOK}, nil
}
