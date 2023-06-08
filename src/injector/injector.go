package injector

import (
	"github.com/tamurakeito/memo-app-backend/src/domain/repository"
	"github.com/tamurakeito/memo-app-backend/src/handler"
	"github.com/tamurakeito/memo-app-backend/src/infra"
	"github.com/tamurakeito/memo-app-backend/src/usecase"
)

func InjectDB() infra.SqlHandler {
    sqlhandler := infra.NewSqlHandler()
    return *sqlhandler
}

/* 
TodoRepository(interface)に実装であるSqlHandler(struct)を渡し生成する。
*/
func InjectTodoRepository() repository.TodoRepository {
    sqlHandler := InjectDB()
    return infra.NewTodoRepository(sqlHandler)
}

func InjectTodoUsecase() usecase.TodoUsecase {
    TodoRepo := InjectTodoRepository()
    return usecase.NewTodoUsecase(TodoRepo)
}

func InjectTodoHandler() handler.TodoHandler {
    return handler.NewTodoHandler(InjectTodoUsecase())
}


