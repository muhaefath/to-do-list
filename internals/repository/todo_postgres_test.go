package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/muhaefath/to-do-list/internals/models"
)

func TestTodoListRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTodoListRepository(db)

	expectedTodos := []models.Todo{
		{ID: "1", Title: "Task 1", Completed: false},
		{ID: "2", Title: "Task 2", Completed: true},
	}

	mock.ExpectQuery("SELECT id, title, completed FROM todolist").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "completed"}).
			AddRow(expectedTodos[0].ID, expectedTodos[0].Title, expectedTodos[0].Completed).
			AddRow(expectedTodos[1].ID, expectedTodos[1].Title, expectedTodos[1].Completed))

	todos, err := repo.GetAll()
	if err != nil {
		t.Fatalf("Error calling GetAll: %v", err)
	}

	if len(todos) != len(expectedTodos) {
		t.Fatalf("Expected %d todos, got %d", len(expectedTodos), len(todos))
	}
	for i := range todos {
		if todos[i].ID != expectedTodos[i].ID || todos[i].Title != expectedTodos[i].Title || todos[i].Completed != expectedTodos[i].Completed {
			t.Fatalf("Unexpected todo at index %d: %v", i, todos[i])
		}
	}
}

func TestTodoListRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTodoListRepository(db)

	expectedTodo := models.Todo{
		ID:        "1",
		Title:     "Task 1",
		Completed: false,
	}

	mock.ExpectQuery("SELECT id, title, completed FROM todolist WHERE id = ?").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "completed"}).
			AddRow(expectedTodo.ID, expectedTodo.Title, expectedTodo.Completed))

	todo, err := repo.GetByID("1")
	if err != nil {
		t.Fatalf("Error calling GetByID: %v", err)
	}

	if todo.ID != expectedTodo.ID || todo.Title != expectedTodo.Title || todo.Completed != expectedTodo.Completed {
		t.Fatalf("Unexpected todo: %+v", todo)
	}
}

func TestTodoListRepository_CreateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTodoListRepository(db)

	todo := models.Todo{
		Title:     "New Task",
		Completed: false,
	}

	mock.ExpectExec("INSERT INTO todolist").
		WithArgs(todo.Title, todo.Completed).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateTodo(&todo)
	if err != nil {
		t.Fatalf("Error calling CreateTodo: %v", err)
	}
}

func TestTodoListRepository_UpdateTodoByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTodoListRepository(db)

	updatedTodo := models.Todo{
		ID:        "1",
		Title:     "Updated Task",
		Completed: true,
	}

	mock.ExpectExec("UPDATE todolist SET").
		WithArgs(updatedTodo.Title, updatedTodo.Completed, updatedTodo.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.UpdateTodoByID("1", &updatedTodo)
	if err != nil {
		t.Fatalf("Error calling UpdateTodoByID: %v", err)
	}
}

func TestTodoListRepository_DeleteTodoByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTodoListRepository(db)

	mock.ExpectExec("DELETE FROM todolist WHERE").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteTodoByID("1")
	if err != nil {
		t.Fatalf("Error calling DeleteTodoByID: %v", err)
	}
}
