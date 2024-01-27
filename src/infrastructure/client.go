package infrastructure

import (
	"fmt"

	"github.com/tamurakeito/memo-app-backend/src/domain/model"
	"github.com/tamurakeito/memo-app-backend/src/domain/repository"
)

type ClientDataRepository struct {
	SqlHandler
}

func NewClientRepository(sqlHandler SqlHandler) repository.ClientDataRepository {
	clientDataRepository := ClientDataRepository{sqlHandler}
	return &clientDataRepository
}

func (clientRepo *ClientDataRepository) Find() (data model.ClientData, err error) {
	row := clientRepo.SqlHandler.Conn.QueryRow("SELECT * FROM client_data")
	err = row.Scan(&data.Tab)
	if err != nil {
		fmt.Print(err)
		return
	}
	return
}

func (clientRepo *ClientDataRepository) Update(data model.ClientData) (model.ClientData, error) {
	_, err := clientRepo.SqlHandler.Conn.Exec("UPDATE client_data SET tab = ?", data.Tab)
	return data, err
}
