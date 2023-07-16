package injector

import (
	"github.com/tamurakeito/memo-app-backend/src/domain/repository"
	"github.com/tamurakeito/memo-app-backend/src/infrastructure"
	"github.com/tamurakeito/memo-app-backend/src/presentation"
	"github.com/tamurakeito/memo-app-backend/src/usecase"
)

func InjectDB() infrastructure.SqlHandler {
	sqlhandler := infrastructure.NewSqlHandler()
	return *sqlhandler
}

/*
Memo
*/
func InjectMemoRepository() repository.MemoRepository {
	sqlHandler := InjectDB()
	return infrastructure.NewMemoRepository(sqlHandler)
}

func InjectMemoUsecase() usecase.MemoUsecase {
	MemoRepo := InjectMemoRepository()
	return usecase.NewMemoUsecase(MemoRepo)
}

func InjectMemoHandler() presentation.MemoHandler {
	return presentation.NewMemoHandler(InjectMemoUsecase())
}

// /*
// Task
// */
// func InjectTaskRepository() repository.TaskRepository {
// 	sqlHandler := InjectDB()
// 	return infrastructure.NewTaskRepository(sqlHandler)
// }

// func InjectTaskUsecase() usecase.TaskUsecase {
// 	TaskRepo := InjectTaskRepository()
// 	return usecase.NewTaskUsecase(TaskRepo)
// }

// func InjectTaskHandler() presentation.TaskHandler {
// 	return presentation.NewTaskHandler(InjectTaskUsecase())
// }
