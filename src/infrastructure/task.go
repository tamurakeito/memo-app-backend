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

func (taskRepo *TaskRepository) Find(memoID int) (tasks []model.Task, err error) {
	// var list_query string
	// for i, v := range list {
	// 	if i > 0 {
	// 		list_query += ","
	// 	}
	// 	list_query += fmt.Sprint(v)
	// }
	rows, err := taskRepo.SqlHandler.Conn.Query("SELECT id, name, complete FROM task_list WHERE memo_id = ?", memoID)
	defer rows.Close()
	if err != nil {
		fmt.Print(err)
		return
	}
	for rows.Next() {
		task := model.Task{}
		task.MemoID = memoID

		rows.Scan(&task.ID, &task.Name, &task.Complete)

		tasks = append(tasks, task)
	}
	return
}

func (taskRepo *TaskRepository) Count(memoID int) (length int, err error) {
	row := taskRepo.SqlHandler.Conn.QueryRow("SELECT COUNT(*) FROM task_list WHERE memo_id = ?", memoID)
	err = row.Scan(&length)
	return
}

func (taskRepo *TaskRepository) Create(task model.Task) (model.Task, error) {
	result, err := taskRepo.SqlHandler.Conn.Exec("INSERT INTO task_list (name,memo_id,complete) VALUES (?, ?, ?) ", task.Name, task.MemoID, task.Complete)
	id, err := result.LastInsertId()
	createdTask := model.Task{ID: int(id), Name: task.Name, MemoID: task.MemoID, Complete: false}
	return createdTask, err
}

func (taskRepo *TaskRepository) Update(task model.Task) (model.Task, error) {
	_, err := taskRepo.SqlHandler.Conn.Exec("UPDATE task_list SET name = ?,complete = ? WHERE id = ?", task.Name, task.Complete, task.ID)
	return task, err
}

func (taskRepo *TaskRepository) Delete(id int) (int, error) {
	_, err := taskRepo.SqlHandler.Conn.Exec("DELETE FROM task_list WHERE id = ?", id)
	return id, err
}
