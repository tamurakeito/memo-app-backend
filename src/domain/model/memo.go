package model

type Memo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  bool   `json:"tag"`
}

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	MemoID   int    `json:"memo_id"`
	Complete bool   `json:"complete"`
}

type ClientData struct {
	Tab int `json:"tab"`
}
