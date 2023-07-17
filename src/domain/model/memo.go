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
	MemoID   int    `json:"memo_id"`
	Complete bool   `json:"complete"`
}
