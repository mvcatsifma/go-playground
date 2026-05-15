package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

type a struct {
	Foo string `json:"foo"`
}

type b struct {
}

type httpResp interface {
	a | b
}

// unmarshaller is a type-parameterized function that decodes a response body into T.
type unmarshaller[T httpResp] func(closer io.ReadCloser) (T, error)

func main() {
	var ua = func(body io.ReadCloser) (a, error) {
		var a = a{}
		data, err := io.ReadAll(body)
		if err != nil {
			return a, err
		}
		err = json.Unmarshal(data, &a)
		if err != nil {
			return a, err
		}
		return a, nil
	}

	var ub = func(body io.ReadCloser) (b, error) {
		return b{}, errors.New("unmarshal error")
	}

	reader := io.NopCloser(strings.NewReader("{\"foo\": \"bar\"}"))
	respA := http.Response{
		Body:       reader,
		StatusCode: http.StatusOK,
	}
	resA, err := handleResponse[a](respA, ua)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", resA)

	respB := http.Response{
		Body:       reader,
		StatusCode: http.StatusBadRequest,
	}
	resB, err := handleResponse[b](respB, ub)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	log.Printf("%+v\n", resB)
}

// handleResponse dispatches on status code and delegates body decoding to the caller-supplied unmarshaller.
func handleResponse[T httpResp](resp http.Response, u unmarshaller[T]) (T, error) {
	var r T
	if resp.StatusCode == http.StatusOK {
		var err error
		r, err = u(resp.Body)
		if err != nil {
			return r, err
		}
		return r, nil
	}
	if resp.StatusCode == http.StatusBadRequest {
		return r, errors.New("bad request")

	}
	return r, nil
}
