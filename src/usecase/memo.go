package usecase

import (
	"github.com/tamurakeito/memo-app-backend/src/domain/entity"
	"github.com/tamurakeito/memo-app-backend/src/domain/repository"
)

type MemoUsecase interface {
	ListSummary() (summary []entity.MemoSummary, err error)
	ListDetail(id int) (detail entity.MemoDetail, err error)
}

type memoUsecase struct {
	memoRepo repository.MemoRepository
	taskRepo repository.TaskRepository
}

// repository.MemoRepository を usecase.MemoUsecase に型変換するだけ
func NewMemoUsecase(memoRepo repository.MemoRepository) MemoUsecase {
	memoUsecase := memoUsecase{memoRepo: memoRepo}
	return &memoUsecase
}

func (usecase *memoUsecase) ListSummary() (summaries []entity.MemoSummary, err error) {
	memos, err := usecase.memoRepo.FindAll()
	for _, memo := range memos {
		summary := entity.MemoSummary{ID: memo.ID, Name: memo.Name, Tag: memo.Tag, Length: len(memo.List)}
		summaries = append(summaries, summary)
	}
	return
}

func (usecase *memoUsecase) ListDetail(id int) (detail entity.MemoDetail, err error) {
	memo, err := usecase.memoRepo.Find(id)
	tasks, err := usecase.taskRepo.Find(memo.List)
	detail = entity.MemoDetail{ID: memo.ID, Name: memo.Name, Tag: memo.Tag, Tasks: tasks}
	return
}

// // Search 入力された内容でTodoを検索する
// func (usecase *todoUsecase) Search(word string) (todo []*model.Todo, err error) {
// 	todo, err = usecase.todoRepo.Find(word)
// 	return
// }
