package infrastructure

import (
	"fmt"

	"github.com/tamurakeito/memo-app-backend/src/domain/model"
	"github.com/tamurakeito/memo-app-backend/src/domain/repository"
)

type TaskRepository struct {
	SqlHandler
}

func NewTaskRepository(sqlHandler SqlHandler) repository.TaskRepository {
	taskRepository := TaskRepository{sqlHandler}
	return &taskRepository
}

func (taskRepo *TaskRepository) Find(list []int) (tasks []model.Task, err error) {
	var list_query string
	for i, v := range list {
		if i > 0 {
			list_query += ","
		}
		list_query += fmt.Sprint(v)
	}
	rows, err := taskRepo.SqlHandler.Conn.Query("SELECT * FROM tasks WHERE id IN (?)", list_query)
	defer rows.Close()
	if err != nil {
		fmt.Print(err)
		return
	}
	for rows.Next() {
		task := model.Task{}

		rows.Scan(&task.ID, &task.Name, &task.Complete)

		tasks = append(tasks, task)
	}
	return
}

func (taskRepo *TaskRepository) Create(task *model.Task) (*model.Task, error) {
	_, err := taskRepo.SqlHandler.Conn.Exec("INSERT INTO tasks (name,complete) VALUES (?, ?) ", task.Name, task.Complete)
	return task, err
}

func (taskRepo *TaskRepository) Update(task *model.Task) (*model.Task, error) {
	_, err := taskRepo.SqlHandler.Conn.Exec("UPDATE tasks SET name = ?,complete = ? WHERE id = ?", task.Name, task.Complete, task.ID)
	return task, err
}