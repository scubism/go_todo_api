package main

import (
	"./app"
	"./controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"os"
	"runtime"
	"time"
)

func main() {
	ConfigRuntime()
	StartGin()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartGin() {
	app.InitApp()
	defer app.CloseApp()

	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Go TODO API")
	})

	v1 := router.Group("/v1")
	{
		v1.GET("/todos", controllers.IndexTodos)
		v1.GET("/todos/:id", controllers.ViewTodo)
		v1.POST("/todos", controllers.CreateTodo)
		v1.PUT("/todos/:id", controllers.UpdateTodo)
		v1.DELETE("/todos/:id", controllers.DeleteTodo)

		v1.POST("/todos/:id/move", controllers.MoveTodo)
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	router.Run(":" + port)
}
