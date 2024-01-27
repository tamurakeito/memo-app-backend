package presentation

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/tamurakeito/memo-app-backend/src/domain/model"
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
		id, err := strconv.Atoi(c.Param("id"))
		model, err := handler.memoUsecase.MemoDetail(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model)
		}
		return c.JSON(http.StatusOK, model)
	}

}

func (handler *MemoHandler) AddMemo() echo.HandlerFunc {

	return func(c echo.Context) error {
		body := model.Memo{}
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, body)
		}
		model, err := handler.memoUsecase.AddMemo(body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model)
		}
		return c.JSON(http.StatusOK, model)
	}

}

func (handler *MemoHandler) AddTask() echo.HandlerFunc {

	return func(c echo.Context) error {
		body := model.Task{}
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, body)
		}
		model, err := handler.memoUsecase.AddTask(body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model)
		}
		return c.JSON(http.StatusOK, model)
	}

}

func (handler *MemoHandler) RestatusMemo() echo.HandlerFunc {

	return func(c echo.Context) error {
		body := model.Memo{}
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, body)
		}
		model, err := handler.memoUsecase.RestatusMemo(body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model)
		}
		return c.JSON(http.StatusOK, model)
	}

}

func (handler *MemoHandler) RestatusTask() echo.HandlerFunc {

	return func(c echo.Context) error {
		body := model.Task{}
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, body)
		}
		model, err := handler.memoUsecase.RestatusTask(body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model)
		}
		return c.JSON(http.StatusOK, model)
	}

}

func (handler *MemoHandler) DeleteMemo() echo.HandlerFunc {

	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		model, err := handler.memoUsecase.DeleteMemo(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model)
		}
		return c.JSON(http.StatusOK, model)
	}

}

func (handler *MemoHandler) DeleteTask() echo.HandlerFunc {

	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		model, err := handler.memoUsecase.DeleteTask(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model)
		}
		return c.JSON(http.StatusOK, model)
	}

}

func (handler *MemoHandler) ClientData() echo.HandlerFunc {

	return func(c echo.Context) error {
		models, err := handler.memoUsecase.ClientData()
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}

}

func (handler *MemoHandler) ClientDataOverrode() echo.HandlerFunc {

	return func(c echo.Context) error {
		body := model.ClientData{}
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, body)
		}
		model, err := handler.memoUsecase.ClientDataOverrode(body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model)
		}
		return c.JSON(http.StatusOK, model)
	}

}
