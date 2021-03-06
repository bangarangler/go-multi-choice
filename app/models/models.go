// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type Answer struct {
	ID         string    `json:"id"`
	QuestionID string    `json:"questionId"`
	OptionID   string    `json:"optionId"`
	IsCorrect  bool      `json:"isCorrect"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type AnswerResponse struct {
	Message  string    `json:"message"`
	Status   int       `json:"status"`
	Data     *Answer   `json:"data"`
	DataList []*Answer `json:"dataList"`
}

type Question struct {
	ID             string            `json:"id"`
	Title          string            `json:"title"`
	QuestionOption []*QuestionOption `json:"questionOption"`
	CreatedAt      time.Time         `json:"createdAt"`
	UpdatedAt      time.Time         `json:"updatedAt"`
}

type QuestionInput struct {
	Title   string                 `json:"title"`
	Options []*QuestionOptionInput `json:"options"`
}

type QuestionOption struct {
	ID         string    `json:"id"`
	QuestionID string    `json:"questionId"`
	Title      string    `json:"title"`
	Position   int       `json:"position"`
	IsCorrect  bool      `json:"isCorrect"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type QuestionOptionInput struct {
	Title     string `json:"title"`
	Position  int    `json:"position"`
	IsCorrect bool   `json:"isCorrect"`
}

type QuestionResponse struct {
	Message  string      `json:"message"`
	Status   int         `json:"status"`
	Data     *Question   `json:"data"`
	DataList []*Question `json:"dataList"`
}
