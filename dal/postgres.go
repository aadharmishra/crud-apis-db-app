package dal

import (
	"context"
	"fmt"

	"crud-apis-db-app/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IPostgresDb interface {
	Create(ctx context.Context, query string, args ...interface{}) error
	Read(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Update(ctx context.Context, query string, args ...interface{}) error
	Delete(ctx context.Context, query string, args ...interface{}) error
	Execute(ctx context.Context, name string, statement string, args ...interface{}) error
}

type Postgres struct {
	Connection *pgxpool.Pool
}

var PostgresDbClient = Postgres{}

func NewPostgres(config config.IConfig) (IPostgresDb, error) {
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

	PostgresDbClient = Postgres{Connection: dbPool}

	return &PostgresDbClient, nil
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

func (p *Postgres) Execute(ctx context.Context, name string, statement string, args ...interface{}) error {
	pool := p.Connection

	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Prepare(ctx, name, statement)
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), name, args...)
	if err != nil {
		return err
	}

	return nil
}
