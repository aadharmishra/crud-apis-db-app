package shared

import (
	"crud-apis-db-app/config"
	db "crud-apis-db-app/dal"
)

type Deps struct {
	Config   config.IConfig
	Database *db.Instances
}
