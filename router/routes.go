package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/marcsvinicius/go-api-jobs/docs"
	"github.com/marcsvinicius/go-api-jobs/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	path := "/api/v1"
	docs.SwaggerInfo.BasePath = path
	v1 := router.Group(path)
	{
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.GetAllOpeningsHandler)
	}
	// Inicializa Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
