package clients

import (
	oAuth "crud-apis-db-app/clients/oAuth"
	"crud-apis-db-app/config"
)

type ClientCfgInstances struct {
	OAuthClientCfg oAuth.IOAuthCfg
}

// NewClientInstance creates an instance of initialized ClientCfgInstances
func NewClientCfgInstances(conf config.IConfig) (*ClientCfgInstances, error) {
	clientCfgInstances := &ClientCfgInstances{}
	oAuthClientCfgInstance, err := initClients(conf)
	if err != nil {
		return nil, err
	}

	// Sets client config instances
	clientCfgInstances.OAuthClientCfg = oAuthClientCfgInstance

	return clientCfgInstances, nil
}

// Simulates the initialization of a client configs
func initClients(config config.IConfig) (oAuth.IOAuthCfg, error) {
	//initialising oAuth client config
	googleOAuthCfgInterface, err := oAuth.NewOAuthClientCfg(config)
	if err != nil {
		return nil, err
	}

	return googleOAuthCfgInterface, nil
}
