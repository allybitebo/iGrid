package api

import "github.com/piusalfred/registry"

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetUserResponse collects the response parameters for the GetUser method.
type AuthUserResponse struct {
	Err error `json:"err"`
}

// Failed implements Failer.
func (r AuthUserResponse) Failed() error {
	return r.Err
}

// GetUserResponse collects the response parameters for the GetUser method.
type GetUserResponse struct {
	User registry.User `json:"user"`
	Err  error         `json:"err"`
}

// Failed implements Failer.
func (r GetUserResponse) Failed() error {
	return r.Err
}

// AddUserResponse collects the response parameters for the AddUser method.
type AddUserResponse struct {
	Err error `json:"err"`
}

// Failed implements Failer.
func (r AddUserResponse) Failed() error {
	return r.Err
}

// ListUserResponse collects the response parameters for the ListUser method.
type ListUserResponse struct {
	Users []registry.User `json:"users"`
	Err   error           `json:"err"`
}

// UpdateUserResponse collects the response parameters for the UpdateUser method.
type UpdateUserResponse struct {
	User registry.User `json:"r0"`
	Err  error         `json:"e1"`
}

// Failed implements Failer.
func (r UpdateUserResponse) Failed() error {
	return r.Err
}

// DeleteUserResponse collects the response parameters for the DeleteUser method.
type DeleteUserResponse struct {
	Err error `json:"e0"`
}

// Failed implements Failer.
func (r DeleteUserResponse) Failed() error {
	return r.Err
}

// Failed implements Failer.
func (r ListUserResponse) Failed() error {
	return r.Err
}

// AddRegionResponse collects the response parameters for the AddRegion method.
type AddRegionResponse struct {
	Err error `json:"err"`
}

// Failed implements Failer.
func (r AddRegionResponse) Failed() error {
	return r.Err
}

// ListRegionsResponse collects the response parameters for the ListRegions method.
type ListRegionsResponse struct {
	Regions []registry.Region `json:"regions"`
	Err     error             `json:"err"`
}

// Failed implements Failer.
func (r ListRegionsResponse) Failed() error {
	return r.Err
}

// Failed implements Failer.
func (r AddNodeResponse) Failed() error {
	return r.Err
}

// GetNodeResponse collects the response parameters for the GetNode method.
type GetNodeResponse struct {
	Node registry.Node `json:"node"`
	Err  error         `json:"err"`
}
