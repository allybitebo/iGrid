package api

import (
	"context"
	"fmt"
	"github.com/piusalfred/registry"
	"github.com/piusalfred/registry/logger"
	"time"
)

// Middleware describes a service middleware.
type Middleware func(registry.Service) registry.Service

func LoggingMiddleware(logger logger.Logger) Middleware {
	return func(next registry.Service) registry.Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   registry.Service
	logger logger.Logger
}

func (l loggingMiddleware) AuthUser(ctx context.Context, id, password string) (err error) {
	defer func(begin time.Time) {
		l.logger.Info(fmt.Sprintf(
			"method: auth took %v to autheticate user with id %s and returned err %v",
			time.Since(begin), id, err))
	}(time.Now())

	err = l.next.AuthUser(ctx, id, password)
	return
}

func (l loggingMiddleware) GetUser(ctx context.Context, id string) (user registry.User, err error) {
	defer func(begin time.Time) {
		l.logger.Info(fmt.Sprintf(
			"method: GetUser with an input %s took %v to return %v with an err %v",
			id, time.Since(begin), user, err))
	}(time.Now())

	user, err = l.next.GetUser(ctx, id)
	return
}

func (l loggingMiddleware) AddUser(ctx context.Context, user registry.User) (err error) {
	defer func(begin time.Time) {
		l.logger.Info(fmt.Sprintf(
			"method: AddUser with an input %v took %v to add user with an err %v",
			user, time.Since(begin), err))
	}(time.Now())

	err = l.next.AddUser(ctx, user)
	return
}

func (l loggingMiddleware) ListUser(ctx context.Context) (users []registry.User, err error) {
	defer func(begin time.Time) {
		l.logger.Info(fmt.Sprintf(
			"method: ListUser with took %v to list all users %v with an err %v",
			time.Since(begin), users, err))
	}(time.Now())

	users, err = l.next.ListUser(ctx)
	return
}

func (l loggingMiddleware) DeleteUser(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		l.logger.Info(fmt.Sprintf(
			"method: DeleteUser with took %v to delete user with id %s with an err %v",
			time.Since(begin), id, err))
	}(time.Now())

	err = l.next.DeleteUser(ctx, id)
	return
}

func (l loggingMiddleware) UpdateUser(ctx context.Context, id string, user registry.User) (u registry.User, err error) {
	defer func(begin time.Time) {
		l.logger.Info(fmt.Sprintf(
			"method: UpdateUser with took %v to update user with id %s from %v to %v with an err %v",
			time.Since(begin), id, user, u, err))
	}(time.Now())

	u, err = l.next.UpdateUser(ctx, id, user)
	return
}

func (l loggingMiddleware) AddNode(ctx context.Context, node registry.Node) (err error) {
	defer func(begin time.Time) {
		l.logger.Info(fmt.Sprintf(
			"method: AddNode with an input %v took %v to add user with an err %v",
			node, time.Since(begin), err))
	}(time.Now())

	err = l.next.AddNode(ctx, node)
	return
}

func (l loggingMiddleware) GetNode(ctx context.Context, id string) (node registry.Node, err error) {
	defer func(begin time.Time) {
		l.logger.Info(fmt.Sprintf(
			"method: GetNode with an input %s took %v to return a user with an err %v",
			id, time.Since(begin), err))
	}(time.Now())

	node, err = l.next.GetNode(ctx, id)
	return
}

func (l loggingMiddleware) ListNodes(ctx context.Context) (nodes []registry.Node, err error) {
	defer func(begin time.Time) {
		l.logger.Info(fmt.Sprintf(
			"method: ListNodes with took %v to list all users returned with an err %v",
			time.Since(begin), err))
	}(time.Now())

	nodes, err = l.next.ListNodes(ctx)
	return
}

func (l loggingMiddleware) DeleteNode(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		l.logger.Info(fmt.Sprintf(
			"method: DeleteNode with took %v to delete user with id %s with an err %v",
			time.Since(begin), id, err))
	}(time.Now())

	err = l.next.DeleteNode(ctx, id)
	return
}

func (l loggingMiddleware) UpdateNode(ctx context.Context, id string, node registry.Node) (n registry.Node, err error) {
	defer func(begin time.Time) {
		l.logger.Info(fmt.Sprintf(
			"method: UpdateNode with took %v to update user with id %s with an err %v",
			time.Since(begin), id, err))
	}(time.Now())

	n, err = l.next.UpdateNode(ctx, id, node)
	return
}

func (l loggingMiddleware) AddRegion(ctx context.Context, region registry.Region) (err error) {
	defer func(begin time.Time) {
		l.logger.Info(fmt.Sprintf(
			"method: AddRegion with an input %v took %v to add region with an err %v",
			region, time.Since(begin), err))
	}(time.Now())

	err = l.next.AddRegion(ctx, region)
	return
}

func (l loggingMiddleware) ListRegions(ctx context.Context) (regions []registry.Region, err error) {
	defer func(begin time.Time) {
		l.logger.Info(fmt.Sprintf(
			"method: ListRegions took %v to list regions %v with an err %v",
			time.Since(begin), regions, err))
	}(time.Now())

	regions, err = l.next.ListRegions(ctx)
	return
}
