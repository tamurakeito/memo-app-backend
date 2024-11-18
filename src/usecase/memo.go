package usecase

import (
	"fmt"
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
	ClientDataOverride(data model.ClientData) (model.ClientData, error)
	MemoOrderOverride(data entity.MemoOrder) (entity.MemoOrder, error)
	TaskOrderOverride(data entity.TaskOrder) (entity.TaskOrder, error)
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
	memo, orders, err := usecase.memoRepo.Find(id)
	if err != nil {
		log.Fatal(err)
		return detail, err
	}
	tasks, err := usecase.taskRepo.Find(memo.ID)

	taskMap := make(map[int]model.Task)
	for _, task := range tasks {
		taskMap[task.ID] = task
	}
	// orderで並べ替える
	orderTasks := make([]model.Task, 0)
	for _, id := range orders.Order {
		if task, exists := taskMap[id]; exists {
			orderTasks = append(orderTasks, task)
			delete(taskMap, id) // 追加されたタスクをtaskMapから削除
		}
	}

	// 残ったタスクを追加する
	for _, task := range taskMap {
		orderTasks = append(orderTasks, task)
	}

	detail = entity.MemoDetail{ID: memo.ID, Name: memo.Name, Tag: memo.Tag, Tasks: orderTasks}
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
	createdTask, err := usecase.taskRepo.Create(task)
	if err != nil {
		log.Fatal(err)
		return task, err
	}
	fmt.Printf("createdTask: id: %d, memoId: %d\n", createdTask.ID, createdTask.MemoID)

	_, orders, err := usecase.memoRepo.Find(createdTask.MemoID)
	if err != nil {
		log.Fatal(err)
		return createdTask, err
	}
	fmt.Printf("orders ID: %d, orders.Order: %v\n", orders.ID, orders.Order)

	if orders.Order == nil {
		orders.Order = []int{}
	}

	// 型を揃えてからappend
	newOrder := append([]int{int(createdTask.ID)}, orders.Order...)
	fmt.Printf("newOrder: %v\n", newOrder)

	newTaskOrder := entity.TaskOrder{ID: orders.ID, Order: newOrder}
	_, err = usecase.memoRepo.UpdateTask(newTaskOrder)
	if err != nil {
		log.Fatal(err)
		return createdTask, err
	}

	return createdTask, nil
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

func (usecase *memoUsecase) ClientDataOverride(data model.ClientData) (model.ClientData, error) {
	data, err := usecase.clientRepo.Update(data)
	return data, err
}

func (usecase *memoUsecase) MemoOrderOverride(data entity.MemoOrder) (entity.MemoOrder, error) {
	data, err := usecase.orderRepo.Update(data)
	return data, err
}

func (usecase *memoUsecase) TaskOrderOverride(data entity.TaskOrder) (entity.TaskOrder, error) {
	data, err := usecase.memoRepo.UpdateTask(data)
	return data, err
}
