package bitbucket

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	// Returned if the specified resource does not exist.
	ErrNotFound = errors.New("Not Found")

	// Returned if the caller attempts to make a call or modify a resource
	// for which the caller is not authorized.
	//
	// The request was a valid request, the caller's authentication credentials
	// succeeded but those credentials do not grant the caller permission to
	// access the resource.
	ErrForbidden = errors.New("Forbidden")

	// Returned if the call requires authentication and either the credentials
	// provided failed or no credentials were provided.
	ErrNotAuthorized = errors.New("Unauthorized")

	// Returned if the caller submits a badly formed request. For example,
	// the caller can receive this return if you forget a required parameter.
	ErrBadRequest = errors.New("Bad Request")
)

// DefaultClient uses DefaultTransport, and is used internall to execute
// all http.Requests. This may be overriden for unit testing purposes.
//
// IMPORTANT: this is not thread safe and should not be touched with
// the exception overriding for mock unit testing.
var DefaultClient = http.DefaultClient

func (c *Client) do(method string, path string, params url.Values, values interface{}, v interface{}) error {

	// create the URI
	uri, err := url.Parse("https://api.bitbucket.org/1.0" + path)
	if err != nil {
		return err
	}

	if params != nil && len(params) > 0 {
		uri.RawQuery = params.Encode()
	}

	// create the request
	req := &http.Request{
		URL:        uri,
		Method:     method,
		ProtoMajor: 1,
		ProtoMinor: 1,
		Close:      true,
		Header:     make(http.Header),
	}

	// construct the body of the request
	if values != nil {
		var body []byte
		if v, ok := values.(url.Values); ok {
			body = []byte(v.Encode())

			// (we'll need this in order to sign the request)
			req.Form = v
		} else {
			var err error
			body, err = json.Marshal(values)
			if err != nil {
				return nil
			}
		}
		req.Body = nopCloser(string(body))
		fmt.Printf("request body:%s", string(body))
	}

	if v, ok := values.(url.Values); ok {
		req.Form = v
	}

	// add authentication to the request
	c.auth.authenticate(req)

	// make the request using the default http client
	resp, err := DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Read the bytes from the body (make sure we defer close the body)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// fmt.Printf("%s %s\n", method, uri)
	// if uri.String() == "https://api.bitbucket.org/1.0/groups/webdevwilson/testgroupsremovemember" {
	// 	fmt.Printf("body:%s\n", body)
	// 	fmt.Printf("status: %d\n", resp.StatusCode)
	// }

	// Check for an http error status (ie not 200 StatusOK)
	switch resp.StatusCode {
	case 404:
		return ErrNotFound
	case 403:
		return ErrForbidden
	case 401:
		return ErrNotAuthorized
	case 400:
		return ErrBadRequest
	}

	// Unmarshall the JSON response
	if v != nil {
		return json.Unmarshal(body, v)
	}

	return nil
}

func nopCloser(str string) io.ReadCloser {
	body := []byte(str)
	buf := bytes.NewBuffer(body)
	return ioutil.NopCloser(buf)
}
