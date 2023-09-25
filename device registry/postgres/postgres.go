package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/piusalfred/registry"
)

var (
	_ registry.UserRepository   = (*userRepo)(nil)
	_ registry.NodeRepository   = (*nodesRepo)(nil)
	_ registry.RegionRepository = (*regionsRepo)(nil)
)

const (
	hostname = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
	sslmode  = "disable"
)

// Connect creates a connection to the PostgreSQL instance and applies any
func Connect() (*sql.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", hostname, port, user, dbname, password, sslmode)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

/*func (p postgresRepo) rowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)

	p.dbLogger.Info(fmt.Sprintf("executing %s inside rowExists method", query))
	err := p.db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		p.dbLogger.Error(fmt.Sprintf("row does not exists:: the error encountered is %v", err))
		return false
	}
	return exists
}

func (p postgresRepo) Auth(ctx context.Context, addr string, token string) (bool, error) {

	p.dbLogger.Debug(fmt.Sprintf("inside Auth() implementation of postgresql"))
	sqlStmt := "SELECT addr FROM nodes WHERE addr=$1 AND token=$2"

	if p.rowExists(sqlStmt, addr, token) {
		return true, nil
	}

	return false, errors.New("invalid credentials")
}

func (p postgresRepo) Revoke(ctx context.Context, addr string, token string) (bool, error) {
	panic("implement me")
}

func (p postgresRepo) Save(ctx context.Context, node Node) error {
	sqlStmt := ` INSERT INTO nodes (addr,name,lat,long) VALUES($1,$2,$3,$4);`
	_, err := p.db.Exec(sqlStmt, node.Addr, node.Name, node.Latitude, node.Longitude)

	if err != nil {
		return err
	}

	return nil
}

func (p postgresRepo) Get(ctx context.Context, addr string) (Node, error) {
	sqlStmt := `SELECT * FROM nodes WHERE addr=$1;`
	var n Node
	row := p.db.QueryRow(sqlStmt, addr)
	err := row.Scan(&n.Addr, &n.Name, &n.Latitude, &n.Longitude, &n.Token, &n.Status, &n.Type, &n.CreatedAt)
	if err != nil {
		return Node{}, err
	}

	return n, nil

}

func (p postgresRepo) Delete(ctx context.Context, addr string) error {
	sqlStatement := `DELETE FROM nodes WHERE addr = $1;`
	_, err := p.db.Exec(sqlStatement, addr)
	if err != nil {
		return err
	}

	return nil
}

func (p postgresRepo) Update(ctx context.Context, node Node, addr string) error {
	sqlStatement := `UPDATE nodes SET name = $2, lat = $3, long = $4 WHERE addr = $1;`
	_, err := p.db.Exec(sqlStatement, addr, node.Name, node.Latitude, node.Longitude)
	if err != nil {
		return err
	}

	return nil
}

func (p postgresRepo) List(ctx context.Context) ([]Node, error) {
	sqlStmt := `SELECT * FROM nodes;`
	rows, err := p.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var nodes []Node
	for rows.Next() {
		n := Node{}
		err := rows.Scan(&n.Addr, &n.Name, &n.Latitude, &n.Longitude, &n.Token, &n.Status, &n.Type, &n.CreatedAt)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, n)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return nodes, nil

}
*/
