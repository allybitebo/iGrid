package api

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	registry "github.com/piusalfred/registry"
	http1 "net/http"
	"net/url"
	"strings"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	AuthUserEndpoint    endpoint.Endpoint
	GetUserEndpoint     endpoint.Endpoint
	AddUserEndpoint     endpoint.Endpoint
	ListUserEndpoint    endpoint.Endpoint
	DeleteUserEndpoint  endpoint.Endpoint
	UpdateUserEndpoint  endpoint.Endpoint
	AddNodeEndpoint     endpoint.Endpoint
	GetNodeEndpoint     endpoint.Endpoint
	ListNodesEndpoint   endpoint.Endpoint
	DeleteNodeEndpoint  endpoint.Endpoint
	UpdateNodeEndpoint  endpoint.Endpoint
	AddRegionEndpoint   endpoint.Endpoint
	ListRegionsEndpoint endpoint.Endpoint
}

// NewServerEndpoints returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func MakeServerEndpoints(s registry.Service) Endpoints {
	return Endpoints{
		AuthUserEndpoint:    MakeAuthUserEndpoint(s),
		AddNodeEndpoint:     MakeAddNodeEndpoint(s),
		AddRegionEndpoint:   MakeAddRegionEndpoint(s),
		AddUserEndpoint:     MakeAddUserEndpoint(s),
		DeleteNodeEndpoint:  MakeDeleteNodeEndpoint(s),
		DeleteUserEndpoint:  MakeDeleteUserEndpoint(s),
		GetNodeEndpoint:     MakeGetNodeEndpoint(s),
		GetUserEndpoint:     MakeGetUserEndpoint(s),
		ListNodesEndpoint:   MakeListNodesEndpoint(s),
		ListRegionsEndpoint: MakeListRegionsEndpoint(s),
		ListUserEndpoint:    MakeListUserEndpoint(s),
		UpdateNodeEndpoint:  MakeUpdateNodeEndpoint(s),
		UpdateUserEndpoint:  MakeUpdateUserEndpoint(s),
	}

}

// MakeClientEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the remote instance, via a transport/http.Client.
// Useful in a registry client.
func MakeClientEndpoints(instance string) (Endpoints, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}
	tgt.Path = ""

	options := []kithttp.ClientOption{}

	var authUserEndpoint endpoint.Endpoint
	{
		authUserEndpoint = kithttp.NewClient(
			http1.MethodGet,
			tgt,
			encodeAuthUserRequest,
			decodeAuthUserResponse,
			options...).Endpoint()
	}

	var getUserEndpoint endpoint.Endpoint
	{
		getUserEndpoint = kithttp.NewClient(
			http1.MethodGet,
			tgt,
			encodeGetUserRequest,
			decodeGetUserResponse,
			options...).Endpoint()
	}

	var addUserEndpoint endpoint.Endpoint
	{
		addUserEndpoint = kithttp.NewClient(
			http1.MethodPost,
			tgt,
			encodeAddUserRequest,
			decodeAddUserResponse,
		).Endpoint()
	}

	var listUserEndpoint endpoint.Endpoint
	{
		listUserEndpoint = kithttp.NewClient(
			http1.MethodGet,
			tgt,
			encodeListUserRequest,
			decodeListUserResponse,
		).Endpoint()
	}

	var deleteUserEndpoint endpoint.Endpoint
	{
		deleteUserEndpoint = kithttp.NewClient(
			http1.MethodDelete,
			tgt,
			encodeDeleteUserRequest,
			decodeDeleteUserResponse,
		).Endpoint()
	}

	var updateUserEndpoint endpoint.Endpoint
	{
		updateUserEndpoint = kithttp.NewClient(
			http1.MethodPatch,
			tgt,
			encodeUpdateUserRequest,
			decodeUpdateUserResponse,
		).Endpoint()
	}

	var addRegionEndpoint endpoint.Endpoint
	{
		addRegionEndpoint = kithttp.NewClient(
			http1.MethodPost,
			tgt,
			encodeAddRegionRequest,
			decodeAddRegionResponse,
		).Endpoint()
	}

	var listRegionsEndpoint endpoint.Endpoint
	{
		listRegionsEndpoint = kithttp.NewClient(
			http1.MethodGet,
			tgt,
			encodeListRegionsRequest,
			decodeListRegionsResponse,
		).Endpoint()
	}

	var getNodeEndpoint endpoint.Endpoint
	{
		getNodeEndpoint = kithttp.NewClient(
			http1.MethodGet,
			tgt,
			encodeGetNodeRequest,
			decodeGetNodeResponse,
			options...).Endpoint()
	}

	var addNodeEndpoint endpoint.Endpoint
	{
		addNodeEndpoint = kithttp.NewClient(
			http1.MethodPost,
			tgt,
			encodeAddNodeRequest,
			decodeAddNodeResponse,
		).Endpoint()
	}

	var listNodeEndpoint endpoint.Endpoint
	{
		listNodeEndpoint = kithttp.NewClient(
			http1.MethodGet,
			tgt,
			encodeListNodeRequest,
			decodeListNodesResponse,
		).Endpoint()
	}

	var deleteNodeEndpoint endpoint.Endpoint
	{
		deleteNodeEndpoint = kithttp.NewClient(
			http1.MethodDelete,
			tgt,
			encodeDeleteNodeRequest,
			decodeDeleteNodeResponse,
		).Endpoint()
	}

	var updateNodeEndpoint endpoint.Endpoint
	{
		updateNodeEndpoint = kithttp.NewClient(
			http1.MethodPatch,
			tgt,
			encodeUpdateNodeRequest,
			decodeUpdateNodeResponse,
		).Endpoint()
	}

	// Note that the request encoders need to modify the request URL, changing
	// the path. That's fine: we simply need to provide specific encoders for
	// each endpoint.

	return Endpoints{
		AuthUserEndpoint:    authUserEndpoint,
		GetUserEndpoint:     getUserEndpoint,
		AddUserEndpoint:     addUserEndpoint,
		ListUserEndpoint:    listUserEndpoint,
		DeleteUserEndpoint:  deleteUserEndpoint,
		UpdateUserEndpoint:  updateUserEndpoint,
		AddNodeEndpoint:     addNodeEndpoint,
		GetNodeEndpoint:     getNodeEndpoint,
		ListNodesEndpoint:   listNodeEndpoint,
		DeleteNodeEndpoint:  deleteNodeEndpoint,
		UpdateNodeEndpoint:  updateNodeEndpoint,
		AddRegionEndpoint:   addRegionEndpoint,
		ListRegionsEndpoint: listRegionsEndpoint,
	}, nil

}

func encodeUpdateNodeRequest(ctx context.Context, req *http1.Request, request interface{}) error {

	r := request.(UpdateNodeRequest)
	nodeID := url.QueryEscape(r.Id)
	req.URL.Path = "/users/" + nodeID
	return encodeRequest(ctx, req, request)
}

func encodeDeleteNodeRequest(ctx context.Context, req *http1.Request, request interface{}) error {
	r := request.(DeleteNodeRequest)
	nodeID := url.QueryEscape(r.Id)
	req.URL.Path = "/users/" + nodeID
	return encodeRequest(ctx, req, request)
}

func encodeListNodeRequest(ctx context.Context, req *http1.Request, request interface{}) error {
	req.URL.Path = "/nodes"
	return encodeRequest(ctx, req, request)
}

func encodeAddNodeRequest(ctx context.Context, req *http1.Request, request interface{}) error {
	req.URL.Path = "/nodes"
	return encodeRequest(ctx, req, request)
}

func encodeGetNodeRequest(ctx context.Context, req *http1.Request, request interface{}) error {
	r := request.(GetNodeRequest)
	nodeID := url.QueryEscape(r.Id)
	req.URL.Path = "/nodes/" + nodeID
	return encodeRequest(ctx, req, request)
}

func encodeListRegionsRequest(ctx context.Context, req *http1.Request, request interface{}) error {
	req.URL.Path = "/regions"
	return encodeRequest(ctx, req, request)
}

func encodeAddRegionRequest(ctx context.Context, req *http1.Request, request interface{}) error {
	req.URL.Path = "/regions"
	return encodeRequest(ctx, req, request)
}

func encodeUpdateUserRequest(ctx context.Context, req *http1.Request, request interface{}) error {
	// r.Methods("PATCH").Path("/users/{id}")
	r := request.(UpdateUserRequest)
	userID := url.QueryEscape(r.Id)
	req.URL.Path = "/users/" + userID
	return encodeRequest(ctx, req, request)
}

func encodeDeleteUserRequest(ctx context.Context, req *http1.Request, request interface{}) error {
	// r.Methods("DELETE").Path("/users/{id}")
	r := request.(DeleteUserRequest)
	userID := url.QueryEscape(r.Id)
	req.URL.Path = "/users/" + userID
	return encodeRequest(ctx, req, request)
}

func encodeListUserRequest(ctx context.Context, req *http1.Request, request interface{}) error {
	req.URL.Path = "/users"
	return encodeRequest(ctx, req, request)
}

func encodeAddUserRequest(ctx context.Context, req *http1.Request, request interface{}) error {
	// r.Methods("POST").Path("/users")
	req.URL.Path = "/users"
	return encodeRequest(ctx, req, request)
}

func encodeAuthUserRequest(ctx context.Context, req *http1.Request, request interface{}) error {
	// r.Methods("GET").Path("/users/{id}")
	r := request.(AuthUserRequest)
	userID := url.QueryEscape(r.Id)
	req.URL.Path = "/auth/" + userID
	return encodeRequest(ctx, req, request)
}

func encodeGetUserRequest(ctx context.Context, req *http1.Request, request interface{}) error {
	// r.Methods("GET").Path("/users/{id}")
	r := request.(GetUserRequest)
	userID := url.QueryEscape(r.Id)
	req.URL.Path = "/users/" + userID
	return encodeRequest(ctx, req, request)
}

// MakeGetUserEndpoint returns an endpoint that invokes GetUser on the service.
func MakeAuthUserEndpoint(s registry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AuthUserRequest)
		e1 := s.AuthUser(ctx, req.Id, req.Password)

		return AuthUserResponse{Err: e1}, nil

	}
}

// MakeGetUserEndpoint returns an endpoint that invokes GetUser on the service.
func MakeGetUserEndpoint(s registry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		r0, e1 := s.GetUser(ctx, req.Id)
		return GetUserResponse{
			User: r0,
			Err:  e1,
		}, nil
	}
}

// MakeAddUserEndpoint returns an endpoint that invokes AddUser on the service.
func MakeAddUserEndpoint(s registry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddUserRequest)
		e0 := s.AddUser(ctx, req.User)
		return AddUserResponse{Err: e0}, nil
	}
}

// MakeListUserEndpoint returns an endpoint that invokes ListUser on the service.
func MakeListUserEndpoint(s registry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r0, e1 := s.ListUser(ctx)
		return ListUserResponse{
			Err:   e1,
			Users: r0,
		}, nil
	}
}

// MakeDeleteUserEndpoint returns an endpoint that invokes DeleteUser on the service.
func MakeDeleteUserEndpoint(s registry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		e0 := s.DeleteUser(ctx, req.Id)
		return DeleteUserResponse{Err: e0}, nil
	}
}

// MakeUpdateUserEndpoint returns an endpoint that invokes UpdateUser on the service.
func MakeUpdateUserEndpoint(s registry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		r0, e1 := s.UpdateUser(ctx, req.Id, req.User)
		return UpdateUserResponse{
			Err:  e1,
			User: r0,
		}, nil
	}
}

// AddNodeRequest collects the request parameters for the AddNode method.
type AddNodeRequest struct {
	Node registry.Node `json:"node"`
}

// AddNodeResponse collects the response parameters for the AddNode method.
type AddNodeResponse struct {
	Err error `json:"e0"`
}

// MakeAddNodeEndpoint returns an endpoint that invokes AddNode on the service.
func MakeAddNodeEndpoint(s registry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddNodeRequest)
		e0 := s.AddNode(ctx, req.Node)
		return AddNodeResponse{Err: e0}, nil
	}
}

// MakeGetNodeEndpoint returns an endpoint that invokes GetNode on the service.
func MakeGetNodeEndpoint(s registry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetNodeRequest)
		r0, e1 := s.GetNode(ctx, req.Id)
		return GetNodeResponse{
			Err:  e1,
			Node: r0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetNodeResponse) Failed() error {
	return r.Err
}

// ListNodesRequest collects the request parameters for the ListNodes method.
type ListNodesRequest struct{}

// ListNodesResponse collects the response parameters for the ListNodes method.
type ListNodesResponse struct {
	Nodes []registry.Node `json:"nodes"`
	Err   error           `json:"err"`
}

// MakeListNodesEndpoint returns an endpoint that invokes ListNodes on the service.
func MakeListNodesEndpoint(s registry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r0, e1 := s.ListNodes(ctx)
		return ListNodesResponse{
			Err:   e1,
			Nodes: r0,
		}, nil
	}
}

// Failed implements Failer.
func (r ListNodesResponse) Failed() error {
	return r.Err
}

// DeleteNodeRequest collects the request parameters for the DeleteNode method.
type DeleteNodeRequest struct {
	Id string `json:"id"`
}

// DeleteNodeResponse collects the response parameters for the DeleteNode method.
type DeleteNodeResponse struct {
	Err error `json:"e0"`
}

// MakeDeleteNodeEndpoint returns an endpoint that invokes DeleteNode on the service.
func MakeDeleteNodeEndpoint(s registry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteNodeRequest)
		e0 := s.DeleteNode(ctx, req.Id)
		return DeleteNodeResponse{Err: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteNodeResponse) Failed() error {
	return r.Err
}

// UpdateNodeRequest collects the request parameters for the UpdateNode method.
type UpdateNodeRequest struct {
	Id   string        `json:"id"`
	User registry.Node `json:"user"`
}

// UpdateNodeResponse collects the response parameters for the UpdateNode method.
type UpdateNodeResponse struct {
	Node registry.Node `json:"r0"`
	Err  error         `json:"e1"`
}

// MakeUpdateNodeEndpoint returns an endpoint that invokes UpdateNode on the service.
func MakeUpdateNodeEndpoint(s registry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateNodeRequest)
		r0, e1 := s.UpdateNode(ctx, req.Id, req.User)
		return UpdateNodeResponse{
			Err:  e1,
			Node: r0,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateNodeResponse) Failed() error {
	return r.Err
}

// MakeAddRegionEndpoint returns an endpoint that invokes AddRegion on the service.
func MakeAddRegionEndpoint(s registry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRegionRequest)
		e0 := s.AddRegion(ctx, req.Region)
		return AddRegionResponse{Err: e0}, nil
	}
}

// MakeListRegionsEndpoint returns an endpoint that invokes ListRegions on the service.
func MakeListRegionsEndpoint(s registry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r0, e1 := s.ListRegions(ctx)
		return ListRegionsResponse{
			Err:     e1,
			Regions: r0,
		}, nil
	}
}

func (e Endpoints) AuthUser(ctx context.Context, id, password string) (err error) {
	request := AuthUserRequest{
		Id:       id,
		Password: password,
	}
	response, err := e.AuthUserEndpoint(ctx, request)
	if err != nil {
		return
	}

	err = response.(AuthUserResponse).Err
	return
}

// GetUser implements Service. Primarily useful in a client.
func (e Endpoints) GetUser(ctx context.Context, id string) (user registry.User, err error) {
	request := GetUserRequest{Id: id}
	response, err := e.GetUserEndpoint(ctx, request)
	if err != nil {
		return
	}

	user = response.(GetUserResponse).User
	return
}

// AddUser implements Service. Primarily useful in a client.
func (e Endpoints) AddUser(ctx context.Context, user registry.User) (e0 error) {
	request := AddUserRequest{User: user}
	response, err := e.AddUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddUserResponse).Err
}

// ListUser implements Service. Primarily useful in a client.
func (e Endpoints) ListUser(ctx context.Context) (users []registry.User, err error) {
	request := ListUserRequest{}
	response, err := e.ListUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	users = response.(ListUserResponse).Users
	return
}

// DeleteUser implements Service. Primarily useful in a client.
func (e Endpoints) DeleteUser(ctx context.Context, id string) (e0 error) {
	request := DeleteUserRequest{Id: id}
	response, err := e.DeleteUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteUserResponse).Err
}

// UpdateUser implements Service. Primarily useful in a client.
func (e Endpoints) UpdateUser(ctx context.Context, id string, user registry.User) (r0 registry.User, e1 error) {
	request := UpdateUserRequest{
		Id:   id,
		User: user,
	}
	response, err := e.UpdateUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateUserResponse).User, response.(UpdateUserResponse).Err
}

// AddNode implements Service. Primarily useful in a client.
func (e Endpoints) AddNode(ctx context.Context, node registry.Node) (e0 error) {
	request := AddNodeRequest{Node: node}
	response, err := e.AddNodeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddNodeResponse).Err
}

// GetNode implements Service. Primarily useful in a client.
func (e Endpoints) GetNode(ctx context.Context, id string) (r0 registry.Node, e1 error) {
	request := GetNodeRequest{Id: id}
	response, err := e.GetNodeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetNodeResponse).Node, response.(GetNodeResponse).Err
}

// ListNodes implements Service. Primarily useful in a client.
func (e Endpoints) ListNodes(ctx context.Context) (r0 []registry.Node, e1 error) {
	request := ListNodesRequest{}
	response, err := e.ListNodesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListNodesResponse).Nodes, response.(ListNodesResponse).Err
}

// DeleteNode implements Service. Primarily useful in a client.
func (e Endpoints) DeleteNode(ctx context.Context, id string) (e0 error) {
	request := DeleteNodeRequest{Id: id}
	response, err := e.DeleteNodeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteNodeResponse).Err
}

// UpdateNode implements Service. Primarily useful in a client.
func (e Endpoints) UpdateNode(ctx context.Context, id string, user registry.Node) (r0 registry.Node, e1 error) {
	request := UpdateNodeRequest{
		Id:   id,
		User: user,
	}
	response, err := e.UpdateNodeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateNodeResponse).Node, response.(UpdateNodeResponse).Err
}

// AddRegion implements Service. Primarily useful in a client.
func (e Endpoints) AddRegion(ctx context.Context, region registry.Region) (e0 error) {
	request := AddRegionRequest{Region: region}
	response, err := e.AddRegionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddRegionResponse).Err
}

// ListRegions implements Service. Primarily useful in a client.
func (e Endpoints) ListRegions(ctx context.Context) (r0 []registry.Region, e1 error) {
	request := ListRegionsRequest{}
	response, err := e.ListRegionsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListRegionsResponse).Regions, response.(ListRegionsResponse).Err
}
