package interface_

import (
	"app/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	usecase *usecase.TodoUsecase
}

func NewTodoHandler(u *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{usecase: u}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.Create(req.Title); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *TodoHandler) ReadAllTodos(c *gin.Context) {
	todos, err := h.usecase.ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	var req struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.Update(req.Title, req.Completed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo updated"})
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	var req struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.Delete(req.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}
