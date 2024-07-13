package repository

import (
	"github.com/tamurakeito/memo-app-backend/src/domain/entity"
	"github.com/tamurakeito/memo-app-backend/src/domain/model"
)

type MemoRepository interface {
	FindAll() (memos []*model.Memo, err error)
	Find(id int) (memo model.Memo, order entity.TaskOrder, err error)
	Create(memo model.Memo) (model.Memo, error)
	Update(memo model.Memo) (model.Memo, error)
	UpdateTask(taskOrder entity.TaskOrder) (entity.TaskOrder, error)
	Delete(id int) (int, error)
}

type TaskRepository interface {
	Find(memoID int) (tasks []model.Task, err error)
	Count(memoID int) (length int, err error)
	Create(task model.Task) (model.Task, error)
	Update(task model.Task) (model.Task, error)
	Delete(id int) (int, error)
}

type ClientDataRepository interface {
	Find() (model.ClientData, error)
	Update(data model.ClientData) (model.ClientData, error)
}

type OrderRepository interface {
	Find() (entity.MemoOrder, error)
	Update(data entity.MemoOrder) (entity.MemoOrder, error)
}
