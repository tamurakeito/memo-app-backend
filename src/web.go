package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/tamurakeito/memo-app-backend/src/handler"
	"github.com/tamurakeito/memo-app-backend/src/injector"
)

func main() {
    fmt.Println("sever start")
    todoHandler := injector.InjectTodoHandler()
    e := echo.New()
    handler.InitRouting(e, todoHandler)
    // Logger.Fatalはエラーメッセージをログに出力しアプリケーションを停止する
    // 重要なエラーが発生した場合に使用される
    // 普通のエラーは通常のエラーハンドリングを使おう
    e.Logger.Fatal(e.Start(":8080"))
}
