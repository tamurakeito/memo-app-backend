package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tamurakeito/memo-app-backend/src/domain/entity"
	"github.com/tamurakeito/memo-app-backend/src/domain/repository"
)

type OrderRepositoryy struct {
	SqlHandler
}

func NewOrderRepositoryy(sqlHandler SqlHandler) repository.OrderRepository {
	orderRepositoryy := OrderRepositoryy{sqlHandler}
	return &orderRepositoryy
}

func (orderRepo *OrderRepositoryy) Find() (data entity.MemoOrder, err error) {
	var jsonData string
	row := orderRepo.SqlHandler.Conn.QueryRow("SELECT `order` FROM memo_order")
	err = row.Scan(&jsonData)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (orderRepo *OrderRepositoryy) Update(data entity.MemoOrder) (entity.MemoOrder, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Print(err)
		return data, err
	}
	_, err = orderRepo.SqlHandler.Conn.Exec("UPDATE memo_order SET `order` = ?", jsonData)
	return data, err
}
