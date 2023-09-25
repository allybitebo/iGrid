package postgres

import (
	"context"
	"database/sql"
	"github.com/piusalfred/registry"
	"github.com/piusalfred/registry/logger"
	"github.com/piusalfred/registry/pkg/errors"
	sql2 "github.com/piusalfred/registry/sql"
	"log"
	"os"
)

var (
	ErrRegionNotFound = errors.New("region not found")
)

type regionsRepo struct {
	db       *sql.DB
	dbLogger logger.Logger
}

func NewRegionRepository(db *sql.DB) registry.RegionRepository {

	dlog, err := logger.New(os.Stdout, "debug")

	if err != nil {
		log.Fatal("could not create regions repository database logger")
	}
	return &regionsRepo{
		db:       db,
		dbLogger: dlog,
	}
}

func (r regionsRepo) Get(ctx context.Context, id string) (registry.Region, error) {
	/*row := r.db.QueryRow(sql2.RegionGetById, id)
	region := registry.Region{}

	switch err := row.Scan(
		&region.ID, &region.Name, &region.Desc); err {

	case sql.ErrNoRows:
		return registry.Region{}, ErrRegionNotFound

	case nil:
		return region, nil

	default:
		return registry.Region{}, err
	}*/

	panic("do something")
}

func (r regionsRepo) Add(ctx context.Context, region registry.Region) (err error) {

	_, err = r.db.Exec(sql2.RegionAddNew,
		region.ID, region.Name, region.Desc)

	if err != nil {
		return err
	}

	return nil
}

func (r regionsRepo) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func (r regionsRepo) List(ctx context.Context) ([]registry.Region, error) {
	rows, err := r.db.Query(sql2.RegionsSelectAll)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var regions []registry.Region
	for rows.Next() {
		r := registry.Region{}
		err := rows.Scan(&r.ID, &r.Name, &r.Desc)
		if err != nil {
			return nil, err
		}

		regions = append(regions, r)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return regions, nil
}

func (r regionsRepo) Update(ctx context.Context, id string, user registry.Region) (registry.Region, error) {
	panic("implement me")
}
