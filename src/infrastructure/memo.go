package infrastructure

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tamurakeito/memo-app-backend/src/domain/entity"
	"github.com/tamurakeito/memo-app-backend/src/domain/model"
	"github.com/tamurakeito/memo-app-backend/src/domain/repository"
)

type MemoRepository struct {
	SqlHandler
}

func NewMemoRepository(sqlHandler SqlHandler) repository.MemoRepository {
	memoRepository := MemoRepository{sqlHandler}
	return &memoRepository
}

func (memoRepo *MemoRepository) FindAll() (memos []*model.Memo, err error) {
	rows, err := memoRepo.SqlHandler.Conn.Query("SELECT id, name, tag FROM memo_list")
	defer rows.Close()
	if err != nil {
		fmt.Print(err)
		return
	}
	for rows.Next() {
		memo := model.Memo{}
		rows.Scan(&memo.ID, &memo.Name, &memo.Tag)

		memos = append(memos, &memo)
	}
	return
}

func (memoRepo *MemoRepository) Find(id int) (memo model.Memo, taskOrder entity.TaskOrder, err error) {
	var jsonData string
	row := memoRepo.SqlHandler.Conn.QueryRow("SELECT id, name, tag, task_order FROM memo_list WHERE id = ?", id)
	err = row.Scan(&memo.ID, &memo.Name, &memo.Tag, &jsonData)
	if err != nil {
		log.Fatal(err)
		return
	}
	var order struct {
		Order []int
	}
	err = json.Unmarshal([]byte(jsonData), &order)
	taskOrder = entity.TaskOrder{
		ID:    memo.ID,
		Order: order.Order,
	}
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (memoRepo *MemoRepository) Create(memo model.Memo) (model.Memo, error) {
	var result sql.Result
	var err error
	orders := struct {
		Order []int `json:"order"`
	}{
		Order: []int{},
	}
	jsonData, _ := json.Marshal(orders)
	result, err = memoRepo.SqlHandler.Conn.Exec("INSERT INTO memo_list (id,name,tag,task_order) VALUES (?, ?, ?, ?) ", memo.ID, memo.Name, memo.Tag, jsonData)
	if err != nil {
		result, err = memoRepo.SqlHandler.Conn.Exec("INSERT INTO memo_list (name,tag,task_order) VALUES (?, ?, ?) ", memo.Name, memo.Tag, jsonData)
		if err != nil {
			return memo, err
		}
	}

	// Get the last inserted ID
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return memo, err
	}

	// Retrieve the inserted memo
	row := memoRepo.SqlHandler.Conn.QueryRow("SELECT id, name, tag FROM memo_list WHERE id = ?", lastInsertId)
	err = row.Scan(&memo.ID, &memo.Name, &memo.Tag)
	if err != nil {
		return memo, err
	}
	return memo, err
}

func (memoRepo *MemoRepository) Update(memo model.Memo) (model.Memo, error) {
	_, err := memoRepo.SqlHandler.Conn.Exec("UPDATE memo_list SET name = ?,tag = ? WHERE id = ?", memo.Name, memo.Tag, memo.ID)
	return memo, err
}

func (memoRepo *MemoRepository) UpdateTask(taskOrder entity.TaskOrder) (entity.TaskOrder, error) {
	orders := struct {
		Order []int `json:"order"`
	}{
		Order: taskOrder.Order,
	}
	jsonData, err := json.Marshal(orders)
	if err != nil {
		log.Fatal(err)
		return taskOrder, err
	}
	_, err = memoRepo.SqlHandler.Conn.Exec("UPDATE memo_list SET task_order = ? WHERE id = ?", jsonData, taskOrder.ID)
	return taskOrder, err
}

func (memoRepo *MemoRepository) Delete(id int) (int, error) {
	_, err := memoRepo.SqlHandler.Conn.Exec("DELETE FROM memo_list WHERE id = ?", id)
	return id, err
}
