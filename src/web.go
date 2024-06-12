package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tamurakeito/memo-app-backend/src/injector"
	handler "github.com/tamurakeito/memo-app-backend/src/presentation"
)

func main() {
	fmt.Println("sever start")
	memoHandler := injector.InjectMemoHandler()
	e := echo.New()
	// CORSの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"http://localhost:3000", "http://tamurakeito.tokyo/memo-app"},
		AllowOrigins: []string{"*"},
	}))

	handler.InitRouting(e, memoHandler)
	// Logger.Fatalはエラーメッセージをログに出力しアプリケーションを停止する
	// 重要なエラーが発生した場合に使用される
	// 普通のエラーは通常のエラーハンドリングを使おう
	e.Logger.Fatal(e.Start(":8080"))
}
