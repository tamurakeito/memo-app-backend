package usecase

import (
	"github.com/tamurakeito/memo-app-backend/src/domain/entity"
	"github.com/tamurakeito/memo-app-backend/src/domain/model"
	"github.com/tamurakeito/memo-app-backend/src/domain/repository"
)

type MemoUsecase interface {
	MemoSummary() (summary []entity.MemoSummary, err error)
	MemoDetail(id int) (detail entity.MemoDetail, err error)
	RestatusTask(task model.Task) (model.Task, error)
	DeleteTask(id int) (int, error)
}

type memoUsecase struct {
	memoRepo repository.MemoRepository
	taskRepo repository.TaskRepository
}

// repository.MemoRepository を usecase.MemoUsecase に型変換するだけ
func NewMemoUsecase(memoRepo repository.MemoRepository, taskRepo repository.TaskRepository) MemoUsecase {
	memoUsecase := memoUsecase{memoRepo: memoRepo, taskRepo: taskRepo}
	return &memoUsecase
}

func (usecase *memoUsecase) MemoSummary() (summaries []entity.MemoSummary, err error) {
	memos, err := usecase.memoRepo.FindAll()
	// if err != nil {
	// 	return
	// }
	for _, memo := range memos {
		length, countErr := usecase.taskRepo.Count(memo.ID)
		if countErr != nil {
			return
		}
		summary := entity.MemoSummary{ID: memo.ID, Name: memo.Name, Tag: memo.Tag, Length: length}
		summaries = append(summaries, summary)
	}
	return
}

func (usecase *memoUsecase) MemoDetail(id int) (detail entity.MemoDetail, err error) {
	memo, err := usecase.memoRepo.Find(id)
	tasks, err := usecase.taskRepo.Find(memo.ID)
	detail = entity.MemoDetail{ID: memo.ID, Name: memo.Name, Tag: memo.Tag, Tasks: tasks}
	return
}

func (usecase *memoUsecase) RestatusTask(task model.Task) (model.Task, error) {
	task, err := usecase.taskRepo.Update(task)
	return task, err
}

func (usecase *memoUsecase) DeleteTask(id int) (int, error) {
	id, err := usecase.taskRepo.Delete(id)
	return id, err
}
