package usecase

import (
	"log"

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
	// MemoOrder() (entity.MemoOrder, error)
	MemoOrderOverrode(data entity.MemoOrder) (entity.MemoOrder, error)
}

type memoUsecase struct {
	memoRepo   repository.MemoRepository
	taskRepo   repository.TaskRepository
	clientRepo repository.ClientDataRepository
	orderRepo  repository.OrderRepository
}

// repository.MemoRepository を usecase.MemoUsecase に型変換するだけ
func NewMemoUsecase(memoRepo repository.MemoRepository, taskRepo repository.TaskRepository, clientRepo repository.ClientDataRepository, orderRepo repository.OrderRepository) MemoUsecase {
	memoUsecase := memoUsecase{memoRepo: memoRepo, taskRepo: taskRepo, clientRepo: clientRepo, orderRepo: orderRepo}
	return &memoUsecase
}

func (usecase *memoUsecase) MemoSummary() (summaries []entity.MemoSummary, err error) {
	memos, err := usecase.memoRepo.FindAll()
	if err != nil {
		return
	}

	orders, err := usecase.orderRepo.Find()
	order := orders.Order

	// idをキーにしてMapを作成
	memoMap := make(map[int]*model.Memo)
	for _, memo := range memos {
		memoMap[memo.ID] = memo
	}

	// orderで並べ替える
	summaries = make([]entity.MemoSummary, 0)
	for _, id := range order {
		if memo, exists := memoMap[id]; exists {
			length, countErr := usecase.taskRepo.Count(memo.ID)
			if countErr != nil {
				return
			}
			summary := entity.MemoSummary{ID: memo.ID, Name: memo.Name, Tag: memo.Tag, Length: length}
			summaries = append(summaries, summary)
		}
	}
	return
}

func (usecase *memoUsecase) MemoDetail(id int) (detail entity.MemoDetail, err error) {
	memo, err := usecase.memoRepo.Find(id)
	if err != nil {
		log.Fatal(err)
		return detail, err
	}
	tasks, err := usecase.taskRepo.Find(memo.ID)
	detail = entity.MemoDetail{ID: memo.ID, Name: memo.Name, Tag: memo.Tag, Tasks: tasks}
	return
}

func (usecase *memoUsecase) AddMemo(memoDetail entity.MemoDetail) (entity.MemoDetail, error) {
	memo, err := usecase.memoRepo.Create(model.Memo{ID: memoDetail.ID, Name: memoDetail.Name, Tag: memoDetail.Tag})
	if err != nil {
		log.Fatal(err)
		return memoDetail, err
	}
	// task_listへタスクの追加
	for _, task := range memoDetail.Tasks {
		_, err = usecase.taskRepo.Create(task)
		if err != nil {
			log.Fatal(err)
			return memoDetail, err
		}
	}
	// orderの更新
	orders, err := usecase.orderRepo.Find()
	if err != nil {
		log.Fatal(err)
		return memoDetail, err
	}
	newOrder := append([]int{memo.ID}, orders.Order...)
	_, err = usecase.orderRepo.Update(entity.MemoOrder{Order: newOrder})
	if err != nil {
		log.Fatal(err)
		return memoDetail, err
	}
	return entity.MemoDetail{ID: memo.ID, Name: memo.Name, Tag: memo.Tag, Tasks: memoDetail.Tasks}, err
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
	if err != nil {
		log.Fatal(err)
		return id, err
	}
	// memoのtask一覧を取得して全部削除する
	tasks, err := usecase.taskRepo.Find(id)
	if err != nil {
		log.Fatal(err)
		return id, err
	}
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

// func (usecase *memoUsecase) MemoOrder() (entity.MemoOrder, error) {
// }

func (usecase *memoUsecase) MemoOrderOverrode(data entity.MemoOrder) (entity.MemoOrder, error) {
	data, err := usecase.orderRepo.Update(data)
	return data, err
}
