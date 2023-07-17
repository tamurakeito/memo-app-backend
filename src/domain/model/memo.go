package model

type Memo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Tag    bool   `json:"tag"`
	Length int    `json:"length"`
}

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}
