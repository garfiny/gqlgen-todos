// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Dates struct {
	CreatedDate    Date  `json:"createdDate"`
	DueDate        *Date `json:"dueDate,omitempty"`
	CompletionDate *Date `json:"completionDate,omitempty"`
}

type Mutation struct {
}

type Query struct {
}
