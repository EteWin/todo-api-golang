package usecase

import (
	"app/internal/domain"
)

type TodoRepository interface {
	Create(todo *domain.Todo) error
	ReadAll() ([]domain.Todo, error)
	FindByID(id string) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(todo *domain.Todo) error
}

type TodoUsecase struct {
	repo TodoRepository
}

func NewTodoUsecase(repo TodoRepository) *TodoUsecase {
	return &TodoUsecase{repo: repo}
}

func (u *TodoUsecase) Create(title string) error {
	todo, err := domain.NewTodo(title)
	if err != nil {
		return err
	}
	return u.repo.Create(todo)
}

func (u *TodoUsecase) ReadAll() ([]domain.Todo, error) {
	return u.repo.ReadAll()
}

func (u *TodoUsecase) Update(id string, title string, completed bool) error {
	todo, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}

	if err := todo.UpdateTitle(title); err != nil {
		return err
	}
	return u.repo.Update(todo)
}

// func (u *TodoUsecase) Update(title string, completed bool) error {
// 	todo := &domain.Todo{
// 		Title:     title,
// 		Completed: completed,
// 	}
// 	if err := todo.Validate(); err != nil {
// 		return err
// 	}
// 	return u.repo.Update(todo)
// }

func (u *TodoUsecase) Delete(id string) error {
	todo := &domain.Todo{
		ID: id,
	}
	return u.repo.Delete(todo)
}
