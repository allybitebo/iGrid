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
	"time"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrUserNotUpdated = errors.New("user not updated")
)

type userRepo struct {
	db       *sql.DB
	dbLogger logger.Logger
}

type dbUser struct {
	ID       string    `json:"id,omitempty"`                  //id or user token | uuid
	Name     string    `json:"name"`                          //fullname
	Email    string    `json:"email"`                         //email
	Password string    `json:"password,omitempty"`            //password of user
	Group    int       `json:"group,omitempty"`               //user group
	Region   string    `json:"region_of_operation,omitempty"` //operating region in case of multi cloud
	Created  time.Time `json:"created,omitempty"`
}

func (u dbUser) toUser() registry.User {
	return registry.User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Group:    u.Group,
		Region:   u.Region,
		Created:  u.Created.Format(time.RFC3339),
	}
}

func fromUser(user registry.User) (dbUser, error) {

	now, err := time.Parse(time.RFC3339, user.Created)

	if err != nil {
		return dbUser{}, err
	}
	return dbUser{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Group:    user.Group,
		Region:   user.Region,
		Created:  now,
	}, nil
}

func NewUserRepository(db *sql.DB) registry.UserRepository {

	dlog, err := logger.New(os.Stdout, "debug")

	if err != nil {
		log.Fatal("could not create user repository database logger")
	}
	return &userRepo{
		db:       db,
		dbLogger: dlog,
	}
}

func (u userRepo) Get(ctx context.Context, id string) (registry.User, error) {
	row := u.db.QueryRow(sql2.UserSelectById, id)
	dUser := dbUser{}

	switch err := row.Scan(
		&dUser.ID, &dUser.Name, &dUser.Email,
		&dUser.Password, &dUser.Group,
		&dUser.Region, &dUser.Created); err {

	case sql.ErrNoRows:
		return registry.User{}, ErrUserNotFound

	case nil:
		return dUser.toUser(), nil

	default:
		return registry.User{}, err
	}
}

func (u userRepo) Add(ctx context.Context, user registry.User) error {

	dUser, err := fromUser(user)

	if err != nil {
		return err
	}

	_, err = u.db.Exec(sql2.UserInsertNew,
		dUser.ID, dUser.Name, dUser.Email, dUser.Password,
		dUser.Group, dUser.Region, dUser.Created)

	if err != nil {
		return err
	}

	return nil
}

func (u userRepo) Delete(ctx context.Context, id string) error {

	_, err := u.db.Exec(sql2.UserDelete, id)
	if err != nil {
		return err
	}

	return nil
}

func (u userRepo) List(ctx context.Context) ([]registry.User, error) {

	rows, err := u.db.Query(sql2.UsersSelectAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []registry.User

	for rows.Next() {
		u := dbUser{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Group, &u.Region, &u.Created)
		if err != nil {
			return nil, err
		}

		users = append(users, u.toUser())
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u userRepo) Update(ctx context.Context, id string, user registry.User) (registry.User, error) {
	//can only update group and region
	region := user.Region
	group := user.Group

	if region != "" {

		//update all
		if group <= 3 && group >= 1 {
			_, err := u.db.Exec(sql2.UserUpdateRandG, id, group, region)
			if err != nil {
				return registry.User{}, err
			}
		} else {
			_, err := u.db.Exec(sql2.UserUpdateRegion, id, region)
			if err != nil {
				return registry.User{}, err
			}
		}
	}

	_, err := u.db.Exec(sql2.UserUpdateGroup, id, group)
	if err != nil {
		return registry.User{}, err
	}

	updatedUser, err := u.Get(ctx, id)
	if err != nil {
		return registry.User{}, ErrUserNotUpdated
	}

	return updatedUser, nil
}
