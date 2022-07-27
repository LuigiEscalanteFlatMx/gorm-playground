package httpd

import (
	"github.com/gin-gonic/gin"
	"github.com/luigiescalante/proyect-template/core/repository/api/handler"
)

func StartApiServer() {
	router := gin.Default()
	apiV1 := router.Group("/flat/v1")
	{
		apiV1.GET("/users", handler.GetUsers)
		apiV1.GET("/users/id/:id", handler.GetUserById)
		apiV1.POST("/users", handler.Save)
		apiV1.PUT("/users/id/:id", handler.Update)
		apiV1.DELETE("/users/id/:id", handler.Delete)
	}
	router.GET("/", handler.HealthyCheck)
	router.GET("/_healthy-check", handler.Ping)
	router.Run()
}
