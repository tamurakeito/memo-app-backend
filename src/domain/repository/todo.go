package repository

import (
	"github.com/tamurakeito/memo-app-backend/src/domain/model"
)

//TodoRepository is interface for infrastructurestructure
type TodoRepository interface {
    FindAll() (todos []*model.Todo, err error)
    Find(word string) (todos []*model.Todo, err error)
    Create(todo *model.Todo) (*model.Todo, error)
    Update(todo *model.Todo) (*model.Todo, error)
}
