package presentation

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/tamurakeito/memo-app-backend/src/usecase"
)

type MemoHandler struct {
	memoUsecase usecase.MemoUsecase
}

func NewMemoHandler(memoUsecase usecase.MemoUsecase) MemoHandler {
	todoHandler := MemoHandler{memoUsecase: memoUsecase}
	return todoHandler
}

func (handler *MemoHandler) MemoSummary() echo.HandlerFunc {

	return func(c echo.Context) error {
		models, err := handler.memoUsecase.MemoSummary()
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}

}

func (handler *MemoHandler) MemoDetail() echo.HandlerFunc {

	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.QueryParam("id"))
		models, err := handler.memoUsecase.MemoDetail(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}

}

// func (handler *MemoHandler) View() echo.HandlerFunc {

// 	return func(c echo.Context) error {
// 		models, err := handler.memoUsecase.View()
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, models)
// 		}
// 		return c.JSON(http.StatusOK, models)
// 	}

// }
// func (handler *MemoHandler) Search() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		word := c.QueryParam("word")
// 		models, err := handler.memoUsecase.Search(word)
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, models)
// 		}
// 		return c.JSON(http.StatusOK, models)
// 	}
// }
// func (handler *MemoHandler) Add() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var todo model.Todo
// 		c.Bind(&todo)
// 		err := handler.memoUsecase.Add(&todo)
// 		return c.JSON(http.StatusOK, err)
// 	}
// }
// func (handler *MemoHandler) Edit() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var todo model.Todo
// 		c.Bind(&todo)
// 		err := handler.memoUsecase.Edit(&todo)
// 		return c.JSON(http.StatusOK, err)
// 	}
// }
