package oauth

import (
	"crud-apis-db-app/shared"

	"github.com/gin-gonic/gin"
)

func NewOAuthRouter(router *gin.Engine, deps *shared.Deps) {
	bindRoutes(router, deps)
}

func bindRoutes(router *gin.Engine, deps *shared.Deps) {
	service := NewOAuthService(deps)
	oAuthRouterGrp := router.Group("/google")
	{
		oAuthRouterGrp.GET("/login", service.Login)
		oAuthRouterGrp.GET("/callback", service.Callback)
		oAuthRouterGrp.POST("/drive/files/create", service.CreateDriveFile)
	}
	youtubeRouterGrp := router.Group("/youtube")
	{
		youtubeRouterGrp.GET("/search", service.YoutubeSearch)
	}
}
