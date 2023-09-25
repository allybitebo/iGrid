package api

import (
	"bytes"
	"context"
	"encoding/json"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	"github.com/piusalfred/registry"
	"io/ioutil"
	http1 "net/http"
	"net/url"
	"strings"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options map[string][]http.ClientOption) (registry.Service, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	var getUserEndpoint endpoint.Endpoint
	{
		getUserEndpoint = http.NewClient("POST", copyURL(u, "/get-user"), encodeRequest, decodeGetUserResponse, options["GetUser"]...).Endpoint()
	}

	var addUserEndpoint endpoint.Endpoint
	{
		addUserEndpoint = http.NewClient("POST", copyURL(u, "/add-user"), encodeRequest, decodeAddUserResponse, options["AddUser"]...).Endpoint()
	}

	var listUserEndpoint endpoint.Endpoint
	{
		listUserEndpoint = http.NewClient("POST", copyURL(u, "/list-user"), encodeRequest, decodeListUserResponse, options["ListUser"]...).Endpoint()
	}

	var deleteUserEndpoint endpoint.Endpoint
	{
		deleteUserEndpoint = http.NewClient("POST", copyURL(u, "/delete-user"), encodeRequest, decodeDeleteUserResponse, options["DeleteUser"]...).Endpoint()
	}

	var updateUserEndpoint endpoint.Endpoint
	{
		updateUserEndpoint = http.NewClient("POST", copyURL(u, "/update-user"), encodeRequest, decodeUpdateUserResponse, options["UpdateUser"]...).Endpoint()
	}

	var addNodeEndpoint endpoint.Endpoint
	{
		addNodeEndpoint = http.NewClient("POST", copyURL(u, "/add-node"), encodeRequest, decodeAddNodeResponse, options["AddNode"]...).Endpoint()
	}

	var getNodeEndpoint endpoint.Endpoint
	{
		getNodeEndpoint = http.NewClient("POST", copyURL(u, "/get-node"), encodeRequest, decodeGetNodeResponse, options["GetNode"]...).Endpoint()
	}

	var listNodesEndpoint endpoint.Endpoint
	{
		listNodesEndpoint = http.NewClient("POST", copyURL(u, "/list-nodes"), encodeRequest, decodeListNodesResponse, options["ListNodes"]...).Endpoint()
	}

	var deleteNodeEndpoint endpoint.Endpoint
	{
		deleteNodeEndpoint = http.NewClient("POST", copyURL(u, "/delete-node"), encodeRequest, decodeDeleteNodeResponse, options["DeleteNode"]...).Endpoint()
	}

	var updateNodeEndpoint endpoint.Endpoint
	{
		updateNodeEndpoint = http.NewClient("POST", copyURL(u, "/update-node"), encodeRequest, decodeUpdateNodeResponse, options["UpdateNode"]...).Endpoint()
	}

	var addRegionEndpoint endpoint.Endpoint
	{
		addRegionEndpoint = http.NewClient("POST", copyURL(u, "/add-region"), encodeRequest, decodeAddRegionResponse, options["AddRegion"]...).Endpoint()
	}

	var listRegionsEndpoint endpoint.Endpoint
	{
		listRegionsEndpoint = http.NewClient("POST", copyURL(u, "/list-regions"), encodeRequest, decodeListRegionsResponse, options["ListRegions"]...).Endpoint()
	}

	return Endpoints{
		AddNodeEndpoint:     addNodeEndpoint,
		AddRegionEndpoint:   addRegionEndpoint,
		AddUserEndpoint:     addUserEndpoint,
		DeleteNodeEndpoint:  deleteNodeEndpoint,
		DeleteUserEndpoint:  deleteUserEndpoint,
		GetNodeEndpoint:     getNodeEndpoint,
		GetUserEndpoint:     getUserEndpoint,
		ListNodesEndpoint:   listNodesEndpoint,
		ListRegionsEndpoint: listRegionsEndpoint,
		ListUserEndpoint:    listUserEndpoint,
		UpdateNodeEndpoint:  updateNodeEndpoint,
		UpdateUserEndpoint:  updateUserEndpoint,
	}, nil
}

// EncodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// SON-encodes any request to the request body. Primarily useful in a client.
func encodeRequest(_ context.Context, r *http1.Request, request interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodeAuthUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAuthUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp AuthUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp GetUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeAddUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAddUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp AddUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp ListUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp DeleteUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp UpdateUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeAddNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAddNodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp AddNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetNodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp GetNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListNodesResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListNodesResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp ListNodesResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteNodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp DeleteNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateNodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp UpdateNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeAddRegionResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAddRegionResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp AddRegionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListRegionsResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListRegionsResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, ErrorDecoder(r)
	}
	var resp ListRegionsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
