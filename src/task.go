package main

import (
	"api-todolist/src/connection"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Task struct
type Task struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// TaskID find one task
type TaskID struct {
	ID int `json:"id"`
}

// Create Function create new Task
func (task Task) Create() error {
	db := connection.GetConnection()

	query := `INSERT INTO task (title, description, status, updated_at) VALUES (?,?,?,?)`

	statement, err := db.Prepare(query)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(task.Title, task.Description, task.Status, time.Now())

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// GetAll function
func (task *Task) GetAll() ([]Task, error) {
	db := connection.GetConnection()
	query := `SELECT *from task`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
		return []Task{}, err
	}

	defer rows.Close()

	tasks := []Task{}

	for rows.Next() {
		rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		tasks = append(tasks, *task)
	}

	return tasks, nil
}

// GetTask get one task
func (task *Task) GetTask(taskID int) (Task, error) {
	db := connection.GetConnection()
	query := `SELECT * FROM task WHERE id=?`

	err := db.QueryRow(query, taskID).Scan(
		&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt,
	)

	if err != nil {
		return Task{}, err
	}

	return *task, nil
}

//Update task
func (task Task) Update(taskID int) error {
	db := connection.GetConnection()

	query := `UPDATE task set title=?, description=?, updated_at=? WHERE id=?`

	statement, err := db.Prepare(query)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(task.Title, task.Description, time.Now(), taskID)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

//Delete task
func (task Task) Delete(taskID int) error {
	db := connection.GetConnection()

	query := `DELETE FROM task WHERE id=?`

	statement, err := db.Prepare(query)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(taskID)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
