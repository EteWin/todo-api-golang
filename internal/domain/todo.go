package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        string    `gorm:"type:uuid;primaryKey"` // UUID
	Title     string    `gorm:"title"`
	Completed bool      `gorm:"completed"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func NewTodo(title string) (*Todo, error) {
	t := &Todo{
		ID:        uuid.NewString(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := t.Validate(); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Todo) Validate() error {
	if t.Title == "" {
		return fmt.Errorf("title is required")
	}
	return nil
}
