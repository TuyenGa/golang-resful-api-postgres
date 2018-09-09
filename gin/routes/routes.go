package routes

import (
	"gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	v1 := router.Group("api")
	{
		userGroup := v1.Group("people")
		{
			userGroup.GET("/", controllers.GetPeople)
			userGroup.POST("/", controllers.Create)
			userGroup.PUT("/{id:int}", controllers.Update)
			userGroup.DELETE("/{id:int}", controllers.Delete)
		}
	}

	return router
}
