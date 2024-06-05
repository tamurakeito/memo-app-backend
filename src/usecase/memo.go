package usecase

import (
	"github.com/tamurakeito/memo-app-backend/src/domain/entity"
	"github.com/tamurakeito/memo-app-backend/src/domain/model"
	"github.com/tamurakeito/memo-app-backend/src/domain/repository"
)

type MemoUsecase interface {
	MemoSummary() (summary []entity.MemoSummary, err error)
	MemoDetail(id int) (detail entity.MemoDetail, err error)
	AddMemo(memo entity.MemoDetail) (entity.MemoDetail, error)
	AddTask(task model.Task) (model.Task, error)
	RestatusMemo(memo model.Memo) (model.Memo, error)
	RestatusTask(task model.Task) (model.Task, error)
	DeleteMemo(id int) (int, error)
	DeleteTask(id int) (int, error)
	ClientData() (model.ClientData, error)
	ClientDataOverrode(data model.ClientData) (model.ClientData, error)
}

type memoUsecase struct {
	memoRepo   repository.MemoRepository
	taskRepo   repository.TaskRepository
	clientRepo repository.ClientDataRepository
}

// repository.MemoRepository を usecase.MemoUsecase に型変換するだけ
func NewMemoUsecase(memoRepo repository.MemoRepository, taskRepo repository.TaskRepository, clientRepo repository.ClientDataRepository) MemoUsecase {
	memoUsecase := memoUsecase{memoRepo: memoRepo, taskRepo: taskRepo, clientRepo: clientRepo}
	return &memoUsecase
}

func (usecase *memoUsecase) MemoSummary() (summaries []entity.MemoSummary, err error) {
	memos, err := usecase.memoRepo.FindAll()
	if err != nil {
		return
	}
	summaries = make([]entity.MemoSummary, 0)
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

func (usecase *memoUsecase) AddMemo(memoDetail entity.MemoDetail) (entity.MemoDetail, error) {
	_, err := usecase.memoRepo.Create(model.Memo{ID: memoDetail.ID, Name: memoDetail.Name, Tag: memoDetail.Tag})
	// task_listへタスクの追加
	for _, task := range memoDetail.Tasks {
		_, err = usecase.taskRepo.Create(task)
	}
	return memoDetail, err
}

func (usecase *memoUsecase) AddTask(task model.Task) (model.Task, error) {
	task, err := usecase.taskRepo.Create(task)
	return task, err
}

func (usecase *memoUsecase) RestatusMemo(memo model.Memo) (model.Memo, error) {
	memo, err := usecase.memoRepo.Update(memo)
	return memo, err
}

func (usecase *memoUsecase) RestatusTask(task model.Task) (model.Task, error) {
	task, err := usecase.taskRepo.Update(task)
	return task, err
}

func (usecase *memoUsecase) DeleteMemo(id int) (int, error) {
	_, err := usecase.memoRepo.Delete(id)
	// memoのtask一覧を取得して全部削除する
	tasks, err := usecase.taskRepo.Find(id)
	for _, task := range tasks {
		_, err := usecase.taskRepo.Delete(task.ID)
		if err != nil {
			return id, err
		}
	}
	return id, err
}

func (usecase *memoUsecase) DeleteTask(id int) (int, error) {
	id, err := usecase.taskRepo.Delete(id)
	return id, err
}

func (usecase *memoUsecase) ClientData() (model.ClientData, error) {
	data, err := usecase.clientRepo.Find()
	return data, err
}

func (usecase *memoUsecase) ClientDataOverrode(data model.ClientData) (model.ClientData, error) {
	data, err := usecase.clientRepo.Update(data)
	return data, err
}
