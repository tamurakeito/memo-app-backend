package presentation

import (
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, memoHandler MemoHandler) {

	// e.GET("/", todoHandler.View())

	// e.GET("/search", todoHandler.Search())

	// e.POST("/todoCreate", todoHandler.Add())

	// e.POST("/todoEdit", todoHandler.Edit())

	e.GET("/list-summary", memoHandler.ListSummary())

	e.GET("/list-detail", memoHandler.ListDetail())
}
