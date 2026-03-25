package persistence

import (
	"app/internal/domain"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepositoryImpl(db *gorm.DB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{db: db}
}

func (r *TodoRepositoryImpl) Create(todo *domain.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepositoryImpl) ReadAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	return todos, r.db.Find(&todos).Error
}

func (r *TodoRepositoryImpl) FindByID(id string) (*domain.Todo, error) {
	var todo domain.Todo
	if err := r.db.First(&todo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepositoryImpl) Update(todo *domain.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepositoryImpl) Delete(todo *domain.Todo) error {
	return r.db.Delete(&domain.Todo{}, "id = ?", todo.ID).Error
}
