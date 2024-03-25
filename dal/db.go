package dal

import "crud-apis-db-app/config"

type Instances struct {
	PostgresDb IPostgresDb
	MongoDb    IMongoDb
}

// NewInstance creates an instance of initialized DBInstances
func NewInstance(conf config.IConfig) (*Instances, error) {
	dbInstances := &Instances{}
	postgresDbInstance, mongoDbIntance, err := initDbs(conf)
	if err != nil {
		return nil, err
	}

	// Sets db instances
	dbInstances.PostgresDb = postgresDbInstance
	dbInstances.MongoDb = mongoDbIntance

	return dbInstances, nil
}

// Simulates the initialization of a db connections
func initDbs(config config.IConfig) (IPostgresDb, IMongoDb, error) {
	//initialising postgres connection pool
	postgresDbInterface, err := NewPostgres(config)
	if err != nil {
		return nil, nil, err
	}

	mongoDbInterface, err := NewMongo(config)
	if err != nil {
		return nil, nil, err
	}

	return postgresDbInterface, mongoDbInterface, nil
}
