package main

import (
	db "app/internal/infrastructure"
	"app/internal/infrastructure/persistence"
	interface_ "app/internal/interface"
	"app/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.Init()
	db.Migrate(database)

	repo := persistence.NewTodoRepositoryImpl(database)
	usecase := usecase.NewTodoUsecase(repo)
	handler := interface_.NewTodoHandler(usecase)

	root := gin.Default()
	root.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	root.POST("/todo", handler.CreateTodo)
	root.GET("/todo", handler.ReadAllTodos)
	root.PUT("/todo", handler.UpdateTodo)
	root.DELETE("/todo", handler.DeleteTodo)
	root.Run(":8080")
}
