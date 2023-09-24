package model

type Memo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  bool   `json:"tag"`
}

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}
