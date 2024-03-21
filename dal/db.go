package dal

import "crud-apis-db-app/config"

type Instances struct {
	Db DBInterface
}

// NewInstance creates an instance of initialized DBInstances
func NewInstance(conf config.IConfig) (*Instances, error) {
	dbInstances := &Instances{}
	myDBInstance, err := initDB(conf)
	if err != nil {
		return nil, err
	}

	// Sets db instance
	dbInstances.Db = myDBInstance

	return dbInstances, nil
}

// Simulates the initialization of a db connection
func initDB(config config.IConfig) (DBInterface, error) {
	dbInterface, err := NewPostgres(config)
	if err != nil {
		return nil, err
	}
	return dbInterface, nil
}
