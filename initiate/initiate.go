package initiate

import (
	"crud-apis-db-app/apis"
	"crud-apis-db-app/clients"
	"crud-apis-db-app/config"
	db "crud-apis-db-app/dal"
	"crud-apis-db-app/shared"
)

func Initiate() error {

	cfg, err := config.NewConfig()
	if cfg == nil || err != nil {
		return err
	}

	// Initializes the DB connections
	dbInstances, err := db.NewInstance(cfg)
	if err != nil {
		return err
	}

	//initializes client configs
	clientCfgInstances, err := clients.NewClientCfgInstances(cfg)
	if err != nil {
		return err
	}

	// loads all common dependencies
	dependencies := shared.Deps{
		Config:        cfg,
		Database:      dbInstances,
		ClientConfigs: clientCfgInstances,
	}

	// Initializes servers
	err = apis.InitServers(&dependencies)
	if err != nil {
		return err
	}

	return nil
}
