// Common ReqBuilder functions.

package http

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	jsoniter "github.com/json-iterator/go"

	u "github.com/martel-innovate/osmops/osmops/util"
)

var GET ReqBuilder = func(request *http.Request) error {
	request.Method = "GET"
	return nil
}

var POST = func(request *http.Request) error {
	request.Method = "POST"
	return nil
}

var PUT = func(request *http.Request) error {
	request.Method = "PUT"
	return nil
}

func At(url *url.URL) ReqBuilder {
	return func(request *http.Request) error {
		if url == nil {
			return errors.New("nil URL")
		}
		request.URL = url
		request.Host = url.Host
		return nil
	}
}

var MediaType = struct {
	u.StrEnum
	JSON, YAML, GZIP u.EnumIx
}{
	StrEnum: u.NewStrEnum("application/json", "application/yaml",
		"application/gzip"),
	JSON: 0,
	YAML: 1,
	GZIP: 2,
}

func Content(mediaType u.EnumIx) ReqBuilder {
	return func(request *http.Request) error {
		request.Header.Set("Content-Type", MediaType.LabelOf(mediaType))
		return nil
	}
}

func Accept(mediaType ...u.EnumIx) ReqBuilder {
	return func(request *http.Request) error {
		ts := []string{}
		for _, mt := range mediaType {
			ts = append(ts, MediaType.LabelOf(mt))
		}
		if len(ts) > 0 {
			request.Header.Set("Accept", strings.Join(ts, ", "))
		}

		return nil
	}
	// TODO implement weights too? Not needed for OSM client.
}

func Authorization(value string) ReqBuilder {
	return func(request *http.Request) error {
		request.Header.Set("Authorization", value)
		return nil
	}
}

type BearerTokenProvider func() (string, error)

func BearerToken(acquireToken BearerTokenProvider) ReqBuilder {
	return func(request *http.Request) error {
		if token, err := acquireToken(); err != nil {
			return err
		} else {
			authValue := fmt.Sprintf("Bearer %s", token)
			return Authorization(authValue)(request)
		}
	}
}

func Body(content []byte) ReqBuilder {
	return func(request *http.Request) error {
		request.ContentLength = int64(len(content))

		if len(content) == 0 {
			// see code comments in Request.NewRequestWithContext about an
			// empty body and backward compat.
			request.Body = http.NoBody
			request.GetBody = func() (io.ReadCloser, error) {
				return http.NoBody, nil
			}
		} else {
			request.Body = io.NopCloser(bytes.NewBuffer(content))

			// the following code does the same as Request.NewRequestWithContext
			// so 307 and 308 redirects can replay the body.
			request.GetBody = func() (io.ReadCloser, error) {
				r := bytes.NewReader(content)
				return io.NopCloser(r), nil
			}
		}

		return nil
	}
}

func JsonBody(content interface{}) ReqBuilder {
	return func(request *http.Request) error {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary // (*)
		if data, err := json.Marshal(content); err != nil {
			return err
		} else {
			return Body(data)(request)
		}
	}
	// (*) json-iterator lib.
	// We use it in place of json from Go's standard lib b/c it can handle
	// the serialisation of fields of type map[interface {}]interface{}
	// where the built-in json module will blow up w/
	//    json: unsupported type: map[interface {}]interface{}
	// If you're reading in YAML and then writing it out as JSON you could get
	// bitten by this. For example, say you use "gopkg.in/yaml.v2" to read
	// some YAML that has a field containing arbitrary JSON into a struct
	// with a field X of type interface{}---you don't know what the JSON looks
	// like, but later on you still want to be able to write it out.
	// The YAML lib will read the JSON into X with a type of
	//   map[interface {}]interface{}
	// but when you call the built-in json.Marshal, it'll blow up in your face
	// b/c it doesn't know how to handle that type.
	// See:
	// - https://stackoverflow.com/questions/35377477
}

// TODO also implement streaming body? most of the standard libs aren't built
// w/ streaming in mind, so in practice you'll likely have the whole body in
// memory most of the time for common cases---e.g. JSON, YAML.

// TODO nil pointer checks. Mostly not implemented!! Catch all occurrences
// of slices, i/f, function args and return an error if nil gets passed in.
// Then write test cases for each. What a schlep!
