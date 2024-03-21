package apis

import (
	"crud-apis-db-app/apis/http"
	"crud-apis-db-app/shared"
)

func InitServers(deps *shared.Deps) error {
	err := http.StartServer(deps)
	if err != nil {
		return err
	}
	return nil
}
