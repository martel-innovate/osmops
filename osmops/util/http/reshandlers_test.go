package http

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

type TestData struct {
	X int         `json:"x"`
	Y interface{} `json:"y"`
}

func stringReader(data string) io.ReadCloser {
	return io.NopCloser(strings.NewReader(data))
}

func send(response *http.Response) ReqSender {
	return func(req *http.Request) (*http.Response, error) {
		return response, nil
	}
}

func TestJsonReaderErrorOnNilResponse(t *testing.T) {
	target := TestData{}
	reader := ReadJsonResponse(&target)
	if err := reader.Handle(nil); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestJsonReaderErrorOnNilTarget(t *testing.T) {
	reader := ReadJsonResponse(nil)
	if err := reader.Handle(&http.Response{}); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestJsonReaderErrorOnUnexpectedResponseCode(t *testing.T) {
	target := TestData{}
	response := &http.Response{
		StatusCode: 400,
		Body:       stringReader(`{"x": 1, "y": {"z": 2}}`),
	}
	_, err := Request(GET).
		SetHandler(ExpectSuccess(), ReadJsonResponse(&target)).
		RunWith(send(response))
	if err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestJsonReaderGetData(t *testing.T) {
	target := TestData{}
	response := &http.Response{
		Body: stringReader(`{"x": 1, "y": {"z": 2}}`),
	}
	res, err := Request(GET).
		SetHandler(ReadJsonResponse(&target)).
		RunWith(send(response))

	if err != nil {
		t.Errorf("want: deserialized JSON; got: %v", err)
	}
	if res != response {
		t.Errorf("want: %v; got: %v", response, res)
	}
	if target.X != 1.0 {
		t.Errorf("want: deserialized JSON; got: %+v", target)
	}
	if y, ok := target.Y.(map[string]interface{}); !ok {
		t.Errorf("want: deserialized JSON; got: %+v", target)
	} else {
		if y["z"] != 2.0 {
			t.Errorf("want: deserialized JSON; got: %+v", target)
		}
	}
}

func TestExpectSuccess(t *testing.T) {
	response := &http.Response{}
	for code := 200; code < 300; code++ {
		response.StatusCode = code
		_, err := Request(GET).
			SetHandler(ExpectSuccess()).
			RunWith(send(response))
		if err != nil {
			t.Errorf("want: success; got: %v", err)
		}
	}
	for _, code := range []int{100, 199, 300, 400, 500} {
		response.StatusCode = code
		_, err := Request(GET).
			SetHandler(ExpectSuccess()).
			RunWith(send(response))
		if err == nil {
			t.Errorf("[%d] want: error; got: nil", code)
		}
	}
}

func TestExpectStatusCodeOneOf(t *testing.T) {
	response := &http.Response{}
	want := []int{200, 201, 404}
	for _, code := range want {
		response.StatusCode = code
		_, err := Request(GET).
			SetHandler(ExpectStatusCodeOneOf(want...)).
			RunWith(send(response))
		if err != nil {
			t.Errorf("want: success; got: %v", err)
		}
	}
	for _, code := range []int{100, 199, 300, 400, 500} {
		response.StatusCode = code
		_, err := Request(GET).
			SetHandler(ExpectStatusCodeOneOf(want...)).
			RunWith(send(response))
		if err == nil {
			t.Errorf("[%d] want: error; got: nil", code)
		}
	}
}

func TestExpectStatusCodeNone(t *testing.T) {
	response := &http.Response{StatusCode: 200}
	_, err := Request(GET).
		SetHandler(ExpectStatusCodeOneOf()).
		RunWith(send(response))
	if err == nil {
		t.Errorf("want: error; got: nil")
	}
}
