package injector

// import (
// 	"github.com/tamurakeito/memo-app-backend/src/domain/repository"
// 	"github.com/tamurakeito/memo-app-backend/src/infrastructure"
// 	"github.com/tamurakeito/memo-app-backend/src/presentation"
// 	"github.com/tamurakeito/memo-app-backend/src/usecase"
// )

// func InjectDB() infrastructure.SqlHandler {
//     sqlhandler := infrastructure.NewSqlHandler()
//     return *sqlhandler
// }

// /*
// TodoRepository(interface)に実装であるSqlHandler(struct)を渡し生成する。
// */
// func InjectTodoRepository() repository.TodoRepository {
//     sqlHandler := InjectDB()
//     return infrastructure.NewTodoRepository(sqlHandler)
// }

// func InjectTodoUsecase() usecase.TodoUsecase {
//     TodoRepo := InjectTodoRepository()
//     return usecase.NewTodoUsecase(TodoRepo)
// }

// func InjectTodoHandler() presentation.TodoHandler {
//     return presentation.NewTodoHandler(InjectTodoUsecase())
// }
