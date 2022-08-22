package http

import (
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"

	"github.com/martel-innovate/osmops/osmops/util"
)

type jsonResReader struct {
	deserialized interface{}
}

func (r *jsonResReader) Handle(res *http.Response) error {
	if res == nil {
		return fmt.Errorf("nil response")
	}
	if r.deserialized == nil {
		return fmt.Errorf("nil deserialization target")
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary // (*)
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(r.deserialized)

	// (*) json-iterator lib.
	// We use it in the JsonBody builder to work around encoding/json's
	// inability to serialise map[interface {}]interface{} types. Here
	// we're parsing JSON into a data structure and AFAICT the built-in
	// json lib can parse pretty much any valid JSON you throw at it.
	// So the only reason to use json-iterator in place of encoding/json
	// is performance: json-iterator is way faster than encoding/json.
}

// ReadJsonResponse builds a ResHandler to deserialise a JSON response body,
// returning any error that stopped it from deserializing the response body.
//
// Example.
//
//     client := &http.Client{Timeout: time.Second * 10}
//     target := &MyData{}
//     Request(
//         GET, At(url),
//         Accept(MediaType.JSON),
//     ).
//     SetHandler(ExpectSuccess(), ReadJsonResponse(target)).
//     RunWith(client.Do)
//
func ReadJsonResponse(target interface{}) ResHandler {
	return &jsonResReader{deserialized: target}
}

type expectSuccessfulResponse struct{}

func (e expectSuccessfulResponse) Handle(res *http.Response) error {
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("expected successful response, got: %s", res.Status)
	}
	return nil
}

// ExpectSuccess builds a ResHandler to check for successful responses.
// If the response code is in the range 200-299 (both inclusive), the
// returned ResHandler does nothing. Otherwise it returns an error---that
// stops any following ResHandler to run.
func ExpectSuccess() ResHandler {
	return expectSuccessfulResponse{}
}

// TODO. Implement expect for other status code ranges too?
// Informational responses (100–199)
// Successful responses (200–299)     --> DONE
// Redirects (300–399)
// Client errors (400–499)
// Server errors (500–599)

type expectStatusCodeInSet struct {
	expectedStatusCodes util.IntSet
}

func (e *expectStatusCodeInSet) Handle(res *http.Response) error {
	if !e.expectedStatusCodes.Contains(res.StatusCode) {
		return fmt.Errorf("unexpected response status: %s", res.Status)
	}
	return nil
}

// ExpectStatusCodeOneOf builds a ResHandler to check the response status code
// is among the given ones.
// If the response code is in the given list, the returned ResHandler does
// nothing. Otherwise it returns an error---that stops any following ResHandler
// to run.
func ExpectStatusCodeOneOf(expectedStatusCode ...int) ResHandler {
	return &expectStatusCodeInSet{
		expectedStatusCodes: util.ToIntSet(expectedStatusCode...),
	}
}
