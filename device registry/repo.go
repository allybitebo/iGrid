package registry

import "context"

type Repository interface {
}

type UserRepository interface {
	Get(ctx context.Context, id string) (User, error)
	Add(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]User, error)
	Update(ctx context.Context, id string, user User) (User, error)
}

type NodeRepository interface {
	Get(ctx context.Context, id string) (Node, error)
	Add(ctx context.Context, user Node) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]Node, error)
	Update(ctx context.Context, id string, user Node) (Node, error)
}

type RegionRepository interface {
	Get(ctx context.Context, id string) (Region, error)
	Add(ctx context.Context, user Region) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]Region, error)
	Update(ctx context.Context, id string, user Region) (Region, error)
}
