package entity

import "github.com/tamurakeito/memo-app-backend/src/domain/model"

type MemoSummary struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Tag    bool   `json:"tag"`
	Length int    `json:"length"`
}

type MemoDetail struct {
	ID    int          `json:"id"`
	Name  string       `json:"name"`
	Tag   bool         `json:"tag"`
	Tasks []model.Task `json:"tasks"`
}

type MemoOrder struct {
	Order []int `json:"order"`
}
