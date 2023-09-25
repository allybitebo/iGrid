package api

import "github.com/piusalfred/registry"

/*var (
	_ ReqValidator = (*GetNodeRequest)(nil)
)

type ReqValidator interface {
	validate()error
}*/

// GetNodeRequest collects the request parameters for the GetNode method.
type GetNodeRequest struct {
	Id string `json:"id"`
}

// GetUserRequest collects the request parameters for the GetUser method.
type GetUserRequest struct {
	Id string `json:"id"`
}

// GetUserRequest collects the request parameters for the GetUser method.
type AuthUserRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

// AddUserRequest collects the request parameters for the AddUser method.
type AddUserRequest struct {
	User registry.User `json:"user"`
}

// ListUserRequest collects the request parameters for the ListUser method.
type ListUserRequest struct{}

// DeleteUserRequest collects the request parameters for the DeleteUser method.
type DeleteUserRequest struct {
	Id string `json:"id"`
}

// UpdateUserRequest collects the request parameters for the UpdateUser method.
type UpdateUserRequest struct {
	Id   string        `json:"id"`
	User registry.User `json:"user"`
}

// AddRegionRequest collects the request parameters for the AddRegion method.
type AddRegionRequest struct {
	Region registry.Region `json:"region"`
}

// ListRegionsRequest collects the request parameters for the ListRegions method.
type ListRegionsRequest struct{}
