package registry

import (
	"context"
	"github.com/piusalfred/registry/logger"
	"github.com/piusalfred/registry/pkg/errors"
)

var (
	ErrBadBodyRequest = errors.New("bad request body, make sure all details are there")
)

// Service describes the service.
type Service interface {
	AuthUser(ctx context.Context, id, password string) error
	//GetUser fetches all users details by specifying the id
	//id is the user uuid/email
	//token is a generated token/password if a user is admin
	GetUser(ctx context.Context, id string) (User, error)

	AddUser(ctx context.Context, user User) error

	//ListUser returns all the list of all available users
	ListUser(ctx context.Context) ([]User, error)

	DeleteUser(ctx context.Context, id string) error

	UpdateUser(ctx context.Context, id string, user User) (User, error)

	AddNode(ctx context.Context, node Node) error

	//GetUser fetches all users details by specifying the id
	//id is the user uuid/email
	//token is a generated token/password if a user is admin
	GetNode(ctx context.Context, id string) (Node, error)

	//ListUser returns all the list of all available users
	ListNodes(ctx context.Context) ([]Node, error)

	DeleteNode(ctx context.Context, id string) error

	UpdateNode(ctx context.Context, id string, user Node) (Node, error)

	AddRegion(ctx context.Context, region Region) error

	ListRegions(ctx context.Context) ([]Region, error)
}

type service struct {
	Users        UserRepository
	Nodes        NodeRepository
	Regions      RegionRepository
	Hasher       Hasher
	Logger       logger.Logger
	UUIDProvider UUIDProvider
}

func (svc service) AuthUser(ctx context.Context, id, password string) error {
	user, err := svc.GetUser(ctx, id)
	if err != nil {
		return err
	}

	hashedPassword := user.Password
	return svc.Hasher.Compare(hashedPassword, password)
}

func (svc service) GetUser(ctx context.Context, id string) (user User, err error) {
	user, err = svc.Users.Get(ctx, id)
	return user, err
}
func (svc service) AddUser(ctx context.Context, user User) (err error) {
	//ONLY NAME EMAIL REGION AND PASSWORD
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return ErrBadBodyRequest
	}

	name := user.Name
	mail := user.Email
	pass := user.Password
	region := user.Region

	u, err := CreateUser(svc.Hasher, svc.UUIDProvider, name, mail, pass, region)

	if err != nil {
		return err
	}

	err = svc.Users.Add(ctx, u)

	return
}
func (svc service) ListUser(ctx context.Context) (users []User, err error) {
	users, err = svc.Users.List(ctx)
	return users, err
}
func (svc service) DeleteUser(ctx context.Context, id string) (err error) {
	err = svc.Users.Delete(ctx, id)
	return
}
func (svc service) UpdateUser(ctx context.Context, id string, user User) (u User, err error) {
	u, err = svc.Users.Update(ctx, id, user)
	return
}

func (svc *service) AddNode(ctx context.Context, node Node) (err error) {

	addr := node.Addr
	name := node.Name
	regi := node.Region
	latd := node.Latd
	long := node.Long
	master := node.Master
	typ := node.Type

	nodeN, err := CreateNode(svc.UUIDProvider, addr, name, regi, latd, long, master, typ)
	err = svc.Nodes.Add(ctx, nodeN)
	return err
}
func (svc *service) GetNode(ctx context.Context, id string) (node Node, err error) {
	node, err = svc.Nodes.Get(ctx, id)
	return node, err
}
func (svc *service) ListNodes(ctx context.Context) (nodes []Node, err error) {
	nodes, err = svc.Nodes.List(ctx)
	return nodes, err
}
func (svc *service) DeleteNode(ctx context.Context, id string) (err error) {
	err = svc.Nodes.Delete(ctx, id)
	return err
}
func (svc *service) UpdateNode(ctx context.Context, id string, node Node) (n Node, err error) {

	n, err = svc.Nodes.Update(ctx, id, node)
	return n, err
}
func (svc *service) AddRegion(ctx context.Context, region Region) (err error) {
	return svc.Regions.Add(ctx, region)
}
func (svc *service) ListRegions(ctx context.Context) (regions []Region, err error) {
	regions, err = svc.Regions.List(ctx)
	return
}

// NewService returns a naive, stateless implementation of Service.
func NewService(users UserRepository, nodes NodeRepository,
	regions RegionRepository, hasher Hasher,
	logger logger.Logger, provider UUIDProvider) Service {
	return &service{
		Users:        users,
		Nodes:        nodes,
		Regions:      regions,
		Hasher:       hasher,
		Logger:       logger,
		UUIDProvider: provider,
	}
}

/*// New returns a Service with all of the expected middleware wired in.
func New(middleware []api.Middleware) Service {
	var svc Service = NewService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}*/
