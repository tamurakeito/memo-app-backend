package presentation

import (
	"net/http"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, memoHandler MemoHandler) {

	// ローカルdockerテスト
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/memo-summary", memoHandler.MemoSummary())
	e.GET("/memo-detail/:id", memoHandler.MemoDetail())
	e.PUT("/restatus-task", memoHandler.RestatusTask())
	e.DELETE("/delete-task/:id", memoHandler.DeleteTask())
}
