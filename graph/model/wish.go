package model

type Wish struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Todos []Todo `json: todos`
}

type NewWish struct {
	Text string `json:"text"`
}
