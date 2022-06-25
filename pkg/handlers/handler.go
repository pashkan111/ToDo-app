package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	// api := router.Group("/api")
	// {
	// 	api.GET("/lists")
	// 	api.GET("/list/:id")
	// 	api.PUT("/list/:id")
	// 	api.POST("/create-list")
	// 	api.GET("/notes")
	// 	api.GET("/note/:id")
	// 	api.PUT("/note/:id")
	// 	api.POST("/create-note")
	// }
	return router
}
