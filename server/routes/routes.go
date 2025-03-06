package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sabrodigan/newWebBoxes/server/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	router := r.Group("/")

	// error handling middleware
	router.Use(middleware.ErrorHandling())

	userRoutes := router.Group("/users")
	{
		RegisterUserRoutes(userRoutes)

	}
	authRoutes := router.Group("/auth")
	{
		RegisterAuthRoutes(authRoutes)
	}
}
