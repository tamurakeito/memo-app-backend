package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tamurakeito/memo-app-backend/src/domain/entity"
	"github.com/tamurakeito/memo-app-backend/src/domain/repository"
)

type OderRepository struct {
	SqlHandler
}

func NewOderRepository(sqlHandler SqlHandler) repository.OderRepository {
	oderRepository := OderRepository{sqlHandler}
	return &oderRepository
}

func (oderRepo *OderRepository) Find() (data entity.MemoOder, err error) {
	var jsonData string
	row := oderRepo.SqlHandler.Conn.QueryRow("SELECT oder FROM memo_oder")
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

func (oderRepo *OderRepository) Update(data entity.MemoOder) (entity.MemoOder, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Print(err)
		return data, err
	}
	_, err = oderRepo.SqlHandler.Conn.Exec("UPDATE memo_oder SET oder = ?", jsonData)
	return data, err
}
