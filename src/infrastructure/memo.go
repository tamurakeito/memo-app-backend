package infrastructure

import (
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
	rows, err := memoRepo.SqlHandler.Conn.Query("SELECT * FROM memos")
	defer rows.Close()
	if err != nil {
		fmt.Print(err)
		return
	}
	for rows.Next() {
		memo := model.Memo{}

		rows.Scan(&memo.ID, &memo.Name, &memo.Tag, &memo.List)

		memos = append(memos, &memo)
	}
	return
}

func (memoRepo *MemoRepository) Find(id int) (memo model.Memo, err error) {
	rows, err := memoRepo.SqlHandler.Conn.Query("SELECT * FROM memos WHERE id = ?", id)
	defer rows.Close()
	if err != nil {
		fmt.Print(err)
		return
	}
	for rows.Next() {
		memo := model.Memo{}
		rows.Scan(&memo.ID, &memo.Name, &memo.Tag, &memo.List)
	}
	return
}

func (memoRepo *MemoRepository) Create(memo *model.Memo) (*model.Memo, error) {
	_, err := memoRepo.SqlHandler.Conn.Exec("INSERT INTO memos (name,tag,list) VALUES (?, ?, ?) ", memo.Name, memo.Tag, memo.List)
	return memo, err
}

func (memoRepo *MemoRepository) Update(memo *model.Memo) (*model.Memo, error) {
	_, err := memoRepo.SqlHandler.Conn.Exec("UPDATE memos SET name = ?,tag = ? ,list = ? WHERE id = ?", memo.Name, memo.Tag, memo.List, memo.ID)
	return memo, err
}
