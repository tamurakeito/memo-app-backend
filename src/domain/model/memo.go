package model

type Memo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  bool   `json:"tag"`
	List []int  `json:"list"`
}

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}
