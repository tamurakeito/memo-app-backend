package presentation

import (
	"net/http"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, memoHandler MemoHandler) {

	// ローカルdockerテスト
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Memo App!")
	})

	e.GET("/memo-summary", memoHandler.MemoSummary())
	e.GET("/memo-detail/:id", memoHandler.MemoDetail())
	e.POST("/add-memo", memoHandler.AddMemo())
	e.POST("/add-task", memoHandler.AddTask())
	e.PUT("/restatus-memo", memoHandler.RestatusMemo())
	e.PUT("/restatus-task", memoHandler.RestatusTask())
	e.DELETE("/delete-memo/:id", memoHandler.DeleteMemo())
	e.DELETE("/delete-task/:id", memoHandler.DeleteTask())
	e.GET("/client-data", memoHandler.ClientData())
	e.PUT("/client-data-override", memoHandler.ClientDataOverrode())
	// e.GET("/memo-order", memoHandler.MemoOrder())
	e.PUT("/memo-order-override", memoHandler.MemoOrderOverrode())
}
