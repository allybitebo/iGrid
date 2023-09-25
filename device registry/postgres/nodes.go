package postgres

import (
	"context"
	"database/sql"
	"github.com/piusalfred/registry"
	"github.com/piusalfred/registry/logger"
	sql2 "github.com/piusalfred/registry/sql"
	"log"
	"os"
)

type nodesRepo struct {
	db       *sql.DB
	dbLogger logger.Logger
}

func NewNodeRepository(db *sql.DB) registry.NodeRepository {

	dlog, err := logger.New(os.Stdout, "debug")

	if err != nil {
		log.Fatal("could not create nodes repository database logger")
	}
	return &nodesRepo{
		db:       db,
		dbLogger: dlog,
	}
}

func (nodes nodesRepo) Get(ctx context.Context, id string) (registry.Node, error) {

	row := nodes.db.QueryRow(sql2.NodeGetById, id)

	node := registry.Node{}

	switch err := row.Scan(
		&node.UUID,
		&node.Addr,
		&node.Name,
		&node.Type,
		&node.Region,
		&node.Latd,
		&node.Long,
		&node.Created,
		&node.Master); err {

	case sql.ErrNoRows:
		return registry.Node{}, ErrUserNotFound

	case nil:
		return node, nil

	default:
		return registry.Node{}, err
	}
}

func (nodes nodesRepo) Add(ctx context.Context, node registry.Node) error {

	_, err := nodes.db.Exec(sql2.NodeAddNew,
		node.UUID,
		node.Addr,
		node.Name,
		node.Type,
		node.Region,
		node.Latd,
		node.Long,
		node.Created,
		node.Master,
	)
	if err != nil {
		return err
	}

	return nil
}

func (nodes nodesRepo) Delete(ctx context.Context, id string) error {

	_, err := nodes.db.Exec(sql2.NodeDelete, id)
	if err != nil {
		return err
	}

	return nil
}

func (nodes nodesRepo) List(ctx context.Context) ([]registry.Node, error) {

	rows, err := nodes.db.Query(sql2.NodeGetAll)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ns []registry.Node

	for rows.Next() {
		node := registry.Node{}
		err := rows.Scan(
			&node.UUID,
			&node.Addr,
			&node.Name,
			&node.Type,
			&node.Region,
			&node.Latd,
			&node.Long,
			&node.Created,
			&node.Master)
		if err != nil {
			return nil, err
		}

		ns = append(ns, node)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return ns, nil
}

func (nodes nodesRepo) Update(ctx context.Context, id string, user registry.Node) (registry.Node, error) {
	panic("implement me")
}
