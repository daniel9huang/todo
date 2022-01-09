package main

/// Go fmt import
import (
	//"fmt"
	"github.com/gin-gonic/gin"
)

// Go main function
func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/todo")
	{
		v1.GET("/", getAllTasks)
		v1.GET("/:title", getOneTaskByTitle)
		v1.POST("/", insertTask)
		v1.PUT("/", updateTaskBodybyTitle)
		v1.DELETE("/:title", deleteTaskByTitle)
	}
	router.Run()
}
