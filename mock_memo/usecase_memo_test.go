package mock_repository

import (
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/tamurakeito/memo-app-backend/src/domain/entity"
	"github.com/tamurakeito/memo-app-backend/src/domain/model"
	"github.com/tamurakeito/memo-app-backend/src/usecase"
)

func TestMemoSummary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var mockFindAll []*model.Memo
	var expected []entity.MemoSummary

	var err error

	mockMemoSample := NewMockMemoRepository(ctrl)
	mockMemoSample.EXPECT().FindAll().Return(mockFindAll, err)

	mockTaskSample := NewMockTaskRepository(ctrl)

	memoUsecase := usecase.NewMemoUsecase(mockMemoSample, mockTaskSample)
	result, err := memoUsecase.MemoSummary()

	if err != nil {
		t.Error("Actual FindAll() is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual FindAll() is not same as expected")
	}

}

func TestMemoDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := 1
	name := "テストメモ"
	tag := false
	length := 5

	taskId := 7
	taskName := "テストタスク"
	complete := false

	var mockMemoFind = model.Memo{
		ID:     id,
		Name:   name,
		Tag:    tag,
		Length: length,
	}
	var mockTaskFind = []model.Task{
		{
			ID:       taskId,
			Name:     taskName,
			Complete: complete,
		},
	}
	var err error
	var expected = entity.MemoDetail{
		ID:    id,
		Name:  name,
		Tag:   tag,
		Tasks: mockTaskFind,
	}

	mockMemoSample := NewMockMemoRepository(ctrl)
	mockMemoSample.EXPECT().Find(id).Return(mockMemoFind, err)

	mockTaskSample := NewMockTaskRepository(ctrl)
	mockTaskSample.EXPECT().Find(id).Return(mockTaskFind, err)

	memoUsecase := usecase.NewMemoUsecase(mockMemoSample, mockTaskSample)
	result, err := memoUsecase.MemoDetail(id)

	if err != nil {
		t.Error("Actual FindAll() is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual FindAll() is not same as expected")
	}

}

func TestRestatusTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	task := model.Task{
		ID:       0,
		Name:     "タスク１",
		Complete: true,
	}

	mockMemoSample := NewMockMemoRepository(ctrl)

	mockTaskSample := NewMockTaskRepository(ctrl)
	mockTaskSample.EXPECT().Update(task).Return(task, nil)

	memoUsecase := usecase.NewMemoUsecase(mockMemoSample, mockTaskSample)
	task, err := memoUsecase.RestatusTask(task)

	if err != nil {
		t.Error("Actual FindAll() is not same as expected")
	}
}
