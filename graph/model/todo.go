package model

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID int64  `json:"userId"`
	User   *User  `json:"user"`
	Dates  Dates  `json:"dates"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID int64  `json:"userId"`
}
