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

func InjectMemoRepository() repository.MemoRepository {
	sqlHandler := InjectDB()
	return infrastructure.NewMemoRepository(sqlHandler)
}

func InjectTaskRepository() repository.TaskRepository {
	sqlHandler := InjectDB()
	return infrastructure.NewTaskRepository(sqlHandler)
}

func InjectMemoUsecase() usecase.MemoUsecase {
	MemoRepo := InjectMemoRepository()
	TaskRepo := InjectTaskRepository()
	return usecase.NewMemoUsecase(MemoRepo, TaskRepo)
}

func InjectMemoHandler() presentation.MemoHandler {
	return presentation.NewMemoHandler(InjectMemoUsecase())
}
