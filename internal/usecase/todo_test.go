package usecase_test

import (
	"app/internal/domain"
	"app/internal/usecase"
	"errors"
	"reflect"
	"testing"
	"time"
)

type mockTodoRepository struct {
	calledCreate   bool
	calledReadAll  bool
	readAllError   error
	calledFindById bool
	calledUpdate   bool
	todo           []domain.Todo
	findError      error
	calledDelete   bool
	deleteError    error
}

func (m *mockTodoRepository) Create(todo *domain.Todo) error {
	m.calledCreate = true
	return nil
}

func (m *mockTodoRepository) ReadAll() ([]domain.Todo, error) {
	m.calledReadAll = true
	if m.readAllError != nil {
		return nil, m.readAllError
	}
	return m.todo, nil
}

func (m *mockTodoRepository) FindByID(id string) (*domain.Todo, error) {
	if m.todo == nil {
		return nil, m.findError
	}
	for i := range m.todo {
		if m.todo[i].ID == id {
			return &m.todo[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func (m *mockTodoRepository) Update(todo *domain.Todo) error {
	m.calledUpdate = true
	if m.findError != nil {
		return m.findError
	}
	for i, t := range m.todo {
		if t.ID == todo.ID {
			m.todo[i] = *todo
			return nil
		}
	}
	return errors.New("todo not found")
}

func (m *mockTodoRepository) Delete(todo *domain.Todo) error {
	m.calledDelete = true

	if m.deleteError != nil {
		return m.deleteError
	}

	for i, t := range m.todo {
		if t.ID == todo.ID {
			m.todo = append(m.todo[:i], m.todo[i+1:]...)
			return nil
		}
	}

	return errors.New("todo not found")
}

func TestCreate_success(t *testing.T) {
	repo := &mockTodoRepository{}
	u := usecase.NewTodoUsecase(repo)
	err := u.Create("test create")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if !repo.calledCreate {
		t.Errorf("expected Create to be called")
	}
}

func TestCreate_EmptyTitle_error(t *testing.T) {
	repo := &mockTodoRepository{}
	u := usecase.NewTodoUsecase(repo)
	err := u.Create("")
	if err == nil {
		t.Errorf("expected error but got %v", err)
	}
	if repo.calledCreate {
		t.Errorf("create should be not called")
	}
}

func TestReadAll_success(t *testing.T) {
	repo := &mockTodoRepository{
		todo: []domain.Todo{
			{
				ID:    "1",
				Title: "title",
			},
		},
	}
	u := usecase.NewTodoUsecase(repo)
	got, err := u.ReadAll()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	want := []domain.Todo{
		{
			ID:    "1",
			Title: "title",
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
	if !repo.calledReadAll {
		t.Errorf("expected ReadAll to be called")
	}
}

func TestReadAll_RepoError(t *testing.T) {
	repo := &mockTodoRepository{
		readAllError: errors.New("repo error"),
	}
	u := usecase.NewTodoUsecase(repo)
	_, err := u.ReadAll()
	if err == nil {
		t.Errorf("expected error but got %v", err)
	}
	if !repo.calledReadAll {
		t.Errorf("expected ReadAll to be called")
	}
}

func TestUpdate_success(t *testing.T) {
	todo, _ := domain.NewTodo("old Todo")

	repo := &mockTodoRepository{
		todo: []domain.Todo{*todo},
	}

	u := usecase.NewTodoUsecase(repo)

	before := repo.todo[0].UpdatedAt
	time.Sleep(time.Millisecond) // UpdatedAtの更新を確実にするために待機

	err := u.Update(todo.ID, "new Todo", true)
	if err != nil {
		t.Fatalf("expected no error but got %v", err)
	}

	after := repo.todo[0].UpdatedAt

	if !after.After(before) {
		t.Errorf("expected UpdatedAt to be updated, before=%v, after=%v", before, after)
	}

	want := domain.Todo{
		ID:        todo.ID,
		Title:     "new Todo",
		Completed: true,
	}

	updated := repo.todo[0]

	if want.Title != updated.Title {
		t.Errorf("expected title to be updated, expected=%v, got=%v", want.Title, updated.Title)
	}

	if want.Completed != updated.Completed {
		t.Errorf("expected completed to be updated, expected=%v, got=%v", want.Completed, updated.Completed)
	}
	if updated.UpdatedAt.IsZero() {
		t.Errorf("UpdatedAt should not be zero")
	}
	if !repo.calledUpdate {
		t.Errorf("expected Update to be called")
	}
}

func TestUpdate_EmptyTitle_error(t *testing.T) {
	todo, _ := domain.NewTodo("old Todo")
	repo := &mockTodoRepository{todo: []domain.Todo{*todo}}
	u := usecase.NewTodoUsecase(repo)
	err := u.Update(todo.ID, "", true)
	if err == nil {
		t.Errorf("expected error but got %v", err)
	}
	if repo.calledUpdate {
		t.Errorf("update should be not called")
	}
}

func TestFindByID_notFound_error(t *testing.T) {
	repo := &mockTodoRepository{findError: errors.New("not found")}
	u := usecase.NewTodoUsecase(repo)
	err := u.Update("existing id", "new Todo", true)
	if err == nil {
		t.Errorf("expected error but got %v", err)
	}
	if repo.calledUpdate {
		t.Errorf("update should be not called")
	}
}

func TestUpdate_RepoError(t *testing.T) {
	todo, _ := domain.NewTodo("old Todo")
	repo := &mockTodoRepository{
		todo:      []domain.Todo{*todo},
		findError: errors.New("repo error")}
	u := usecase.NewTodoUsecase(repo)
	err := u.Update(todo.ID, "new Todo", true)
	if err == nil {
		t.Errorf("expected error but got %v", err)
	}
	if !repo.calledUpdate {
		t.Errorf("expected update to be called")
	}
}

func TestDelete_success(t *testing.T) {
	todo, _ := domain.NewTodo("old Todo")
	repo := &mockTodoRepository{
		todo: []domain.Todo{*todo},
	}
	t.Logf("got %v", repo.todo)
	u := usecase.NewTodoUsecase(repo)
	err := u.Delete(todo.ID)
	if err != nil {
		t.Fatalf("expected no error but got %v", err)
	}
	if !repo.calledDelete {
		t.Errorf("expected Delete to be called")
	}
	if len(repo.todo) != 0 {
		t.Errorf("expected todo to be deleted, got %v", repo.todo)
	} else {
		t.Logf("got %v", repo.todo)
	}
}

func TestDelete_RepoError(t *testing.T) {
	todo, _ := domain.NewTodo("todo")
	repo := &mockTodoRepository{
		todo:        []domain.Todo{*todo},
		deleteError: errors.New("repo error"),
	}
	u := usecase.NewTodoUsecase(repo)
	err := u.Delete(todo.ID)
	if err == nil {
		t.Errorf("expected error but got %v", err)
	}
	if !repo.calledDelete {
		t.Errorf("expected Delete to be called")
	}
}

func TestDelete_NotFound_error(t *testing.T) {
	todo, _ := domain.NewTodo("old Todo")
	repo := &mockTodoRepository{
		todo: []domain.Todo{*todo},
	}
	u := usecase.NewTodoUsecase(repo)
	err := u.Delete("non-existent id") // IDはUUIDで生成されるので、存在しえないIDとして指定
	if err == nil {
		t.Errorf("expected error but got %v", err)
	}
	if !repo.calledDelete {
		t.Errorf("expected Delete to be called")
	}
}
