package registry

import (
	"context"
	"time"
)

type EventName int

const (
	CREATE_USER EventName = iota
	DELETE_USER
	UPDATE_USER
	LIST_USERS
	GET_USER
	PUBLISH
	SUBSCRIBE
	CREATE_NODE
	LIST_NODES
)

type Event struct {
	UUID      string        `json:"uuid"`
	Name      string        `json:"name"`
	Region    string        `json:"region"`
	Actor     string        `json:"actor"`
	Action    string        `json:"action"`
	Result    string        `json:"result"`
	Err       string        `json:"err"`
	Timestamp time.Duration `json:"timestamp"`
	ExecTime  time.Duration `json:"exec_time"`
}

type EventStore interface {
	Save(ctx context.Context, event Event) (err error)
	Pull(ctx context.Context) (events []Event, err error)
	Between(ctx context.Context, start time.Duration, end time.Duration) (events []Event, err error)
	Before(ctx context.Context, end time.Duration) (events []Event, err error)
	After(ctx context.Context, start time.Duration) (events []Event, err error)
	ByID(ctx context.Context, id string) (events []Event, err error)
	ByEventName(ctx context.Context, name EventName) (events []Event, err error)
}

type eventStore struct {
}

var _ EventStore = (*eventStore)(nil)

func NewEventStore() EventStore {
	return eventStore{}
}

func (e eventStore) Save(ctx context.Context, event Event) (err error) {
	panic("implement me")
}

func (e eventStore) Pull(ctx context.Context) (events []Event, err error) {
	panic("implement me")
}

func (e eventStore) Between(ctx context.Context, start time.Duration, end time.Duration) (events []Event, err error) {
	panic("implement me")
}

func (e eventStore) Before(ctx context.Context, end time.Duration) (events []Event, err error) {
	panic("implement me")
}

func (e eventStore) After(ctx context.Context, start time.Duration) (events []Event, err error) {
	panic("implement me")
}

func (e eventStore) ByID(ctx context.Context, id string) (events []Event, err error) {
	panic("implement me")
}

func (e eventStore) ByEventName(ctx context.Context, name EventName) (events []Event, err error) {
	panic("implement me")
}
