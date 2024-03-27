package dal

import "crud-apis-db-app/config"

type Instances struct {
	PostgresDb IPostgresDb
	MongoDb    IMongoDb
	RedisDb    IRedisDb
}

// NewInstance creates an instance of initialized DBInstances
func NewInstance(conf config.IConfig) (*Instances, error) {
	dbInstances := &Instances{}
	postgresDbInstance, mongoDbInstance, redisDbInstance, err := initDbs(conf)
	if err != nil {
		return nil, err
	}

	// Sets db instances
	dbInstances.PostgresDb = postgresDbInstance
	dbInstances.MongoDb = mongoDbInstance
	dbInstances.RedisDb = redisDbInstance

	return dbInstances, nil
}

// Simulates the initialization of a db connections
func initDbs(config config.IConfig) (IPostgresDb, IMongoDb, IRedisDb, error) {
	//initialising postgres connection pool
	postgresDbInterface, err := NewPostgres(config)
	if err != nil {
		return nil, nil, nil, err
	}

	//initialising mongodb connection pool
	mongoDbInterface, err := NewMongo(config)
	if err != nil {
		return nil, nil, nil, err
	}

	//initialising redis connection pool
	redisDbInterface, err := NewRedis(config)
	if err != nil {
		return nil, nil, nil, err
	}

	return postgresDbInterface, mongoDbInterface, redisDbInterface, nil
}
