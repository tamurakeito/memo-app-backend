package infrastructure

import (
	"database/sql"
	"fmt"

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
	rows, err := memoRepo.SqlHandler.Conn.Query("SELECT * FROM memo_list")
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

func (memoRepo *MemoRepository) Find(id int) (memo model.Memo, err error) {
	row := memoRepo.SqlHandler.Conn.QueryRow("SELECT * FROM memo_list WHERE id = ?", id)
	// defer rows.Close()
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }
	// for rows.Next() {
	// 	memo := model.Memo{}
	// 	rows.Scan(&memo.ID, &memo.Name, &memo.Tag, &memo.Length)
	// }
	err = row.Scan(&memo.ID, &memo.Name, &memo.Tag)
	if err != nil {
		fmt.Print(err)
		return
	}
	return

	// memo = model.Memo{
	// 	ID:     1,
	// 	Name:   "テスト",
	// 	Tag:    true,
	// 	Length: 7,
	// }
	// err = nil
	// return
}

func (memoRepo *MemoRepository) Create(memo model.Memo) (model.Memo, error) {
	var result sql.Result
	var err error
	result, err = memoRepo.SqlHandler.Conn.Exec("INSERT INTO memo_list (id,name,tag) VALUES (?, ?, ?) ", memo.ID, memo.Name, memo.Tag)
	if err != nil {
		result, err = memoRepo.SqlHandler.Conn.Exec("INSERT INTO memo_list (name,tag) VALUES (?, ?) ", memo.Name, memo.Tag)
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

func (memoRepo *MemoRepository) Delete(id int) (int, error) {
	_, err := memoRepo.SqlHandler.Conn.Exec("DELETE FROM memo_list WHERE id = ?", id)
	return id, err
}
