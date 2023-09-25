package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/piusalfred/registry"
	"net/http"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func MakeHTTPHandler(service registry.Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(service)

	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(ErrorEncoder),
	}

	//GET /users/{id}
	//POST /users
	//GET /users
	//DELETE /users/{id}
	//PATCH /users/{id}

	r.Methods(http.MethodGet).Path("/auth/{id}").Handler(kithttp.NewServer(
		e.AuthUserEndpoint,
		decodeAuthUserRequest,
		encodeAuthUserResponse,
		options...,
	))

	r.Methods(http.MethodGet).Path("/users/{id}").Handler(kithttp.NewServer(
		e.GetUserEndpoint,
		decodeGetUserRequest,
		encodeGetUserResponse,
		options...,
	))

	r.Methods(http.MethodPost).Path("/users").Handler(kithttp.NewServer(
		e.AddUserEndpoint,
		decodeAddUserRequest,
		encodeAddUserResponse,
		options...,
	))

	r.Methods(http.MethodGet).Path("/users").Handler(kithttp.NewServer(
		e.ListUserEndpoint,
		decodeListUserRequest,
		encodeListUserResponse,
		options...,
	))

	r.Methods(http.MethodDelete).Path("/users/{id}").Handler(kithttp.NewServer(
		e.DeleteUserEndpoint,
		decodeDeleteUserRequest,
		encodeDeleteUserResponse,
		options...,
	))

	r.Methods(http.MethodPatch).Path("/users/{id}").Handler(kithttp.NewServer(
		e.UpdateUserEndpoint,
		decodeUpdateUserRequest,
		encodeUpdateUserResponse,
		options...,
	))

	//regions
	r.Methods(http.MethodPost).Path("/regions").Handler(kithttp.NewServer(
		e.AddRegionEndpoint,
		decodeAddRegionRequest,
		encodeAddRegionResponse,
		options...,
	))

	r.Methods(http.MethodGet).Path("/regions").Handler(kithttp.NewServer(
		e.ListRegionsEndpoint,
		decodeListRegionsRequest,
		encodeListRegionsResponse,
		options...,
	))

	//nodes
	r.Methods(http.MethodGet).Path("/nodes/{id}").Handler(kithttp.NewServer(
		e.GetNodeEndpoint,
		decodeGetNodeRequest,
		encodeGetNodeResponse,
		options...,
	))

	r.Methods(http.MethodPost).Path("/nodes").Handler(kithttp.NewServer(
		e.AddNodeEndpoint,
		decodeAddNodeRequest,
		encodeAddNodeResponse,
		options...,
	))

	r.Methods(http.MethodGet).Path("/nodes").Handler(kithttp.NewServer(
		e.ListNodesEndpoint,
		decodeListNodesRequest,
		encodeListNodesResponse,
		options...,
	))

	r.Methods(http.MethodDelete).Path("/nodes/{id}").Handler(kithttp.NewServer(
		e.DeleteNodeEndpoint,
		decodeDeleteNodeRequest,
		encodeDeleteNodeResponse,
		options...,
	))

	r.Methods(http.MethodPatch).Path("/nodes/{id}").Handler(kithttp.NewServer(
		e.UpdateNodeEndpoint,
		decodeUpdateNodeRequest,
		encodeUpdateNodeResponse,
		options...,
	))

	return r
}

func encodeAuthUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

func decodeAuthUserRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	password, ok := vars["password"]
	if !ok {
		return nil, ErrBadRouting
	}
	req := AuthUserRequest{
		Id:       id,
		Password: password,
	}
	return req, nil
}

// decodeGetUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	req := GetUserRequest{Id: id}
	return req, nil
}

// encodeGetUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeAddUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := AddUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAddUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeListUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := ListUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeListUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeDeleteUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	req := DeleteUserRequest{Id: id}
	return req, nil
}

// encodeDeleteUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeUpdateUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)
	_, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	req := UpdateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeAddNodeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddNodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := AddNodeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAddNodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddNodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeGetNodeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetNodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	req := GetNodeRequest{Id: id}
	return req, nil
}

// encodeGetNodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetNodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeListNodesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListNodesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := ListNodesRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeListNodesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListNodesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeDeleteNodeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteNodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	req := DeleteNodeRequest{Id: id}
	return req, nil
}

// encodeDeleteNodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteNodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeUpdateNodeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateNodeRequest(_ context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)
	_, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	req := UpdateNodeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateNodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateNodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeAddRegionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddRegionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := AddRegionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAddRegionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddRegionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return

}

// decodeListRegionsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListRegionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := ListRegionsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeListRegionsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListRegionsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
