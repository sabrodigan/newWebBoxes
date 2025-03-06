package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sabrodigan/newWebBoxes/server/config"
	"github.com/sabrodigan/newWebBoxes/server/forBoxes"
	"github.com/sabrodigan/newWebBoxes/server/routes"
)

func main() {

	// make sure the backend is there, with a ping
	forBoxes.Ping()

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	fmt.Println("Starting application...")
	gin.SetMode(gin.ReleaseMode)
	config.ConfEnvChecks()

	app := gin.Default()
	app.Use(gin.Recovery())

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":    "pong",
			"statusCode": 200,
		})
	})
	routes.RegisterRoutes(app)

	fmt.Println("The application is now running!")

	port, err := config.GetEnvProperty("port")
	if err != nil {
		fmt.Println("Error getting port:", err)
		return
	}

	fmt.Println("Running on port:", port)
	err = app.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Println("Error running application:", err)
		return
	}
}
