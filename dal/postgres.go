package dal

import (
	"context"
	"fmt"

	"crud-apis-db-app/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBInterface interface {
	Create(ctx context.Context, query string, args ...interface{}) error
	Read(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Update(ctx context.Context, query string, args ...interface{}) error
	Delete(ctx context.Context, query string, args ...interface{}) error
}

type Postgres struct {
	Connection *pgxpool.Pool
}

var DBClient = Postgres{}

func NewPostgres(config config.IConfig) (DBInterface, error) {
	cfgPg := config.Get().Postgres
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfgPg.User, cfgPg.Password, cfgPg.Host, cfgPg.Port, cfgPg.Dbname)

	pgsConfig, err := pgxpool.ParseConfig(connString)
	if pgsConfig == nil || err != nil {
		fmt.Printf("Unable to create connection pool: %v\n", err)
		return nil, err
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), pgsConfig)
	if err != nil || dbPool == nil {
		fmt.Printf("Unable to create connection pool: %v\n", err)
		return nil, err
	}

	DBClient = Postgres{Connection: dbPool}

	return &DBClient, nil
}

func (p *Postgres) Create(ctx context.Context, query string, args ...interface{}) error {
	pool := p.Connection
	_, err := pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) Read(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	pool := p.Connection

	rows, err := pool.Query(ctx, query, args...)
	if err != nil || rows == nil {
		return nil, err
	}

	return rows, nil
}

func (p *Postgres) Update(ctx context.Context, query string, args ...interface{}) error {
	pool := p.Connection
	_, err := pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) Delete(ctx context.Context, query string, args ...interface{}) error {
	pool := p.Connection
	_, err := pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
