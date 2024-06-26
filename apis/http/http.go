package http

import (
	"crud-apis-db-app/shared"
	"fmt"

	"crud-apis-db-app/apis/http/departments"
	"crud-apis-db-app/apis/http/employees"
	oauth "crud-apis-db-app/apis/http/oAuth"
	"crud-apis-db-app/apis/http/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer(deps *shared.Deps) error {

	address := deps.Config.Get().Server.Http.Address
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	corsConfig := cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowOrigins: []string{"http://localhost:8080"},
	}
	router.Use(cors.New(corsConfig))
	fmt.Printf("HTTP Server listening on : " + address)

	//register the routes
	users.NewUsersRouter(router, deps)
	employees.NewEmployeesRouter(router, deps)
	departments.NewDepartmentsRouter(router, deps)
	oauth.NewOAuthRouter(router, deps)

	// Start the server
	err := router.Run(address)
	if err != nil {
		return err
	}

	return nil
}
