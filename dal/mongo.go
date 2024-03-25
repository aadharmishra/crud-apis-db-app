package dal

import (
	"context"
	"crud-apis-db-app/config"
	"crypto/tls"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type IMongoDb interface {
	Exists(ctx context.Context, db string, collection string, filter interface{}) (int64, error)
	Create(ctx context.Context, db string, collection string, doc []interface{}) error
	Read(ctx context.Context, db string, collection string, filter interface{}) (*mongo.Cursor, error)
	Update(ctx context.Context, db string, collection string, models []mongo.WriteModel) error
	Delete(ctx context.Context, db string, collection string, filter interface{}) error
}

type Mongo struct {
	Connection *mongo.Client
}

var MongoDbClient = Mongo{}

func NewMongo(config config.IConfig) (IMongoDb, error) {
	cfgMongo := config.Get().Mongodb

	mongoConnOptions := &options.ClientOptions{
		Hosts:      cfgMongo.Hosts,
		ReplicaSet: &cfgMongo.ReplicaSet,
		ServerAPIOptions: &options.ServerAPIOptions{
			ServerAPIVersion: options.ServerAPIVersion1,
		},
		AppName:      &cfgMongo.AppName,
		RetryWrites:  &cfgMongo.RetryWrites,
		WriteConcern: &writeconcern.WriteConcern{W: "majority"},
		MinPoolSize:  &cfgMongo.MinPoolSize,
		MaxPoolSize:  &cfgMongo.MaxPoolSize,
		Auth: &options.Credential{
			Username:    cfgMongo.Username,
			Password:    cfgMongo.Password,
			AuthSource:  cfgMongo.AuthSource,
			PasswordSet: true,
		},
		TLSConfig: new(tls.Config), //needed because ssl is enabled by default in mongo config
	}

	client, err := mongo.Connect(context.Background(), mongoConnOptions)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	fmt.Println("Successfully connected to MongoDB")

	MongoDbClient = Mongo{Connection: client}

	return &MongoDbClient, nil
}

func (m *Mongo) Create(ctx context.Context, db string, collection string, doc []interface{}) error {
	conn := m.Connection
	if conn == nil {
		return errors.New("empty connection")
	}

	handle := conn.Database(db).Collection(collection)
	if handle == nil {
		return errors.New("empty collection handle")
	}

	result, err := handle.InsertMany(ctx, doc)
	if err != nil || result == nil {
		return err
	}

	return nil
}

func (m *Mongo) Delete(ctx context.Context, db string, collection string, filter interface{}) error {
	conn := m.Connection
	if conn == nil {
		return errors.New("empty connection")
	}

	handle := conn.Database(db).Collection(collection)
	if handle == nil {
		return errors.New("empty collection handle")
	}

	result, err := handle.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	if result == nil {
		return errors.New("error while deleting records")
	}

	return nil
}

func (m *Mongo) Read(ctx context.Context, db string, collection string, filter interface{}) (*mongo.Cursor, error) {
	conn := m.Connection
	if conn == nil {
		return nil, errors.New("empty connection")
	}

	handle := conn.Database(db).Collection(collection)
	if handle == nil {
		return nil, errors.New("empty collection handle")
	}

	cursor, err := handle.Find(context.Background(), filter)

	if cursor == nil || err != nil {
		return nil, errors.New("error while finding data")
	}

	return cursor, nil
}

func (m *Mongo) Update(ctx context.Context, db string, collection string, models []mongo.WriteModel) error {
	conn := m.Connection
	if conn == nil {
		return errors.New("empty connection")
	}

	handle := conn.Database(db).Collection(collection)
	if handle == nil {
		return errors.New("empty collection handle")
	}

	result, err := handle.BulkWrite(ctx, models)
	if err != nil || result == nil {
		return err
	}

	return nil
}

func (m *Mongo) Exists(ctx context.Context, db string, collection string, filter interface{}) (int64, error) {
	var count int64
	var err error

	conn := m.Connection
	if conn == nil {
		return count, errors.New("empty connection")
	}

	handle := conn.Database(db).Collection(collection)
	if handle == nil {
		return count, errors.New("empty collection handle")
	}

	count, err = handle.CountDocuments(ctx, filter)
	if err != nil {
		return count, err
	}

	return count, err
}
