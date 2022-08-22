// Utils to make your life slightly easier when working with the HTTP client
// from `net/http`.
//
// Request building
//
// You can put together a request assembly line by mixing and matching little,
// discrete, reusable pieces of functionality encapsulated by ReqBuilder
// functions. Request building reads like an HTTP request on the wire and is
// more type-safe than doing it the Go way. Example:
//
//     req, err := BuildRequest(
//         POST, At(url),
//         Content(MediaType.JSON),
//         Body(content),
//     )
//
// Response handling
//
// Similarly, ResHandler lets you build a response processing pipeline. You
// can use either or both request assembly and response pipeline facilities
// with Go's built-in HTTP client from `net/http`. Example:
//
//     client := &http.Client{Timeout: time.Second * 10}
//     if res, err := client.Do(req); err != nil {
//         err = HandleResponse(req, jsonReader, responseLogger)
//         // HandleResponse calls res.Body.Close for you if needed.
//     }
//
// Message exchange
//
// Out of convenience, there's also an Exchange type to string builders and
// handlers together in an HTTP request-reply message flow where execution
// stops at the first error---so you don't have to litter your code with
// `if err ...` statements. Here's an example that also showcases some of
// the built-in ResHandlers.
//
//     url, _ := url.Parse("http://yapi")
//     responseData := YouData{}
//     client := &http.Client{Timeout: time.Second * 10}
//     Request(
//         POST, At(url),
//         Content(MediaType.JSON),
//         Body(content),
//     ).
//     SetHandler(ExpectSuccess(), ReadJsonResponse(&responseData)).
//     RunWith(client.Do)
//
// Notice RunWith takes a ReqSender so you can easily unit-test your code by
// swapping out an actual HTTP call with a stub. Example:
//
//     send := func(req *http.Request) (*http.Response, error) {
//         return &http.Response{StatusCode: 200}, nil
//     }
//     Request(...).RunWith(send)
//
package http

import (
	"errors"
	"fmt"
	"net/http"
)

// ReqBuilder sets some fields of an HTTP request, possibly returning an error
// if something goes wrong.
// You chain ReqBuilders to build a full HTTP request, each builder contributes
// its bit and the request building process stops at the first error. Basically
// a poor man's monomorphic either+IO monad stack---ask Google.
type ReqBuilder func(request *http.Request) error

func emptyRequest() *http.Request {
	bare := &http.Request{}
	withCtx := bare.WithContext(bare.Context())
	withCtx.Header = make(http.Header)
	return withCtx
}

// BuildRequest runs the given builders to assemble an HTTP request.
// If all the builders run successfully, then the returned request is okay.
// Otherwise BuildRequest stops as soon as a builder errors out, returning
// that error.
func BuildRequest(builders ...ReqBuilder) (*http.Request, error) {
	request := emptyRequest()
	for _, build := range builders {
		if err := build(request); err != nil {
			return request, err
		}
	}
	return request, nil
}

// ResHandler processes an HTTP response.
// You can chain ResHandlers to do more than one thing with the response so
// the code stays modular---single responsibility principle, anyone? The
// response processing chain stops at the first error. Basically a poor man's
// monomorphic either+IO monad stack---ask Google.
type ResHandler interface {
	// Handle processes the given response, possibly returning an error if
	// something goes wrong.
	Handle(response *http.Response) error
}

// HandleResponse feeds the given response to each ResHandler.
// It runs the handlers in the same order as the input arguments, stopping
// at the first one that errors out and returning that error. If all the
// handlers are successful, the returned error will be nil.
// If the response contains a body, HandleResponse automatically closes the
// associated reader just before returning, so handlers don't have to do that.
func HandleResponse(response *http.Response, handlers ...ResHandler) error {
	if response == nil {
		return errors.New("nil response")
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	for k, h := range handlers {
		if h == nil {
			return fmt.Errorf("nil response handler [%d]", k)
		}
		if err := h.Handle(response); err != nil {
			return err
		}
	}
	return nil
}

// Exchange represents an HTTP request-reply message flow initiated by the
// client.
type Exchange struct {
	builders []ReqBuilder
	handlers []ResHandler
}

// Request instantiates a new Exchange with the given ReqBuilder functions.
func Request(builders ...ReqBuilder) *Exchange {
	return &Exchange{builders: builders}
}

// SetHandler specifies the handlers that the Exchange will use to process
// the response.
func (e *Exchange) SetHandler(handlers ...ResHandler) *Exchange {
	e.handlers = handlers
	return e
}

// ReqSender is a function to send an HTTP request and receive a response
// from the server. An error gets returned if something goes wrong---e.g.
// a network failure.
type ReqSender func(*http.Request) (*http.Response, error)

// RunWith performs the HTTP message Exchange by building the request,
// invoking the given send function with it, and finally processing the
// response.
//
// The request gets built by calling the ReqBuilder functions passed to
// the Request factory function. If there's a request build error, then
// RunWith stops there, returning the error along with a nil response.
//
// Otherwise, if the request gets built properly, RunWith calls the given
// send function with it to carry out the client-server HTTP exchange. If
// the send function returns an error, RunWith stops there, returning the
// error and whatever response the send function returned---that would be
// nil in most cases.
//
// Otherwise, if the response was received successfully, RunWith passes it
// on to the handlers configured by SetHandler, in turn and in the same order
// as the handlers were passed to SetHandler. If there's no handlers, then
// the response gets returned without further processing along with a nil
// error. Similarly, if all handlers are successful, the response gets
// returned with a nil error. Otherwise the response gets returned with
// the error output by the first failed handler---RunWith won't call any
// handlers following the failed one.
func (e *Exchange) RunWith(send ReqSender) (*http.Response, error) {
	if send == nil {
		return nil, errors.New("nil ReqSender")
	}

	req, err := BuildRequest(e.builders...)
	if err != nil {
		return nil, err
	}

	res, err := send(req)
	if err != nil {
		return res, err
	}

	return res, HandleResponse(res, e.handlers...)
}

// TODO nil pointer checks. Mostly not implemented!! Catch all occurrences
// of slices, i/f, function args and return an error if nil gets passed in.
// Then write test cases for each. What a schlep!
