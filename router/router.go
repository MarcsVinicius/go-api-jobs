package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicia Rotas
	r := gin.Default()
	initializeRoutes(r)
	// Roda servidor
	r.Run(":8080")
}
