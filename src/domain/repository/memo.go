package repository

import (
	"github.com/tamurakeito/memo-app-backend/src/domain/model"
)

type MemoRepository interface {
	FindAll() (memos []*model.Memo, err error)
	Find(id int) (memo model.Memo, err error)
	Create(memo *model.Memo) (*model.Memo, error)
	Update(memo *model.Memo) (*model.Memo, error)
	Delete(id int) (int, error)
}

type TaskRepository interface {
	Find(memoID int) (tasks []model.Task, err error)
	Count(memoID int) (length int, err error)
	Create(task *model.Task) (*model.Task, error)
	Update(task model.Task) (model.Task, error)
	Delete(id int) (int, error)
}
