package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	// "fmt"
	"log"
	"net/http"
	"time"

	"github.com/bangarangler/go-multi-choice/app/models"
	model "github.com/bangarangler/go-multi-choice/app/models"
	"github.com/bangarangler/go-multi-choice/helpers"
)

func (r *mutationResolver) CreateAnswer(ctx context.Context, questionID string, optionID string) (*model.AnswerResponse, error) {
	ans := &models.Answer{
		QuestionID: questionID,
		OptionID:   optionID,
	}

	if ok, errorString := helpers.ValidateInputs(*ans); !ok {
		return &models.AnswerResponse{
			Message: errorString,
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	correctOpt, err := r.QuestionOptionService.GetQuestionOptionByID(optionID)
	if err != nil {
		return &models.AnswerResponse{
			Message: "Error getting question option",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	if correctOpt.IsCorrect == true {
		ans.IsCorrect = true
	} else {
		ans.IsCorrect = false
	}
	ans.CreatedAt = time.Now()
	ans.UpdatedAt = time.Now()

	answer, err := r.AnsService.CreateAnswer(ans)
	if err != nil {
		log.Println("Answer creation error: ", err)
		return &models.AnswerResponse{
			Message: "Error creating annswer",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.AnswerResponse{
		Message: "Successfully created answer",
		Status:  http.StatusCreated,
		Data:    answer,
	}, nil
}

func (r *mutationResolver) UpdateAnswer(ctx context.Context, id string, questionID string, optionID string) (*model.AnswerResponse, error) {
	ans, err := r.AnsService.GetAnswerByID(id)
	if err != nil {
		log.Println("Error getting the answer to update: ", err)
		return &models.AnswerResponse{
			Message: "Error getting the answer",
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	ans.OptionID = optionID
	ans.QuestionID = questionID
	ans.UpdatedAt = time.Now()

	if ok, errorString := helpers.ValidateInputs(*ans); !ok {
		return &model.AnswerResponse{
			Message: errorString,
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	correctOpt, err := r.AnsService.GetAnswerByID(optionID)
	if err != nil {
		return &models.AnswerResponse{
			Message: "Error getting question option",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	if correctOpt.IsCorrect == true {
		ans.IsCorrect = true
	} else {
		ans.IsCorrect = false
	}

	answer, err := r.AnsService.UpdateAnswer(ans)
	if err != nil {
		log.Println("Answer updating error: ", err)
		return &models.AnswerResponse{
			Message: "Error updating answer",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.AnswerResponse{
		Message: "Successfully updated answer",
		Status:  http.StatusOK,
		Data:    answer,
	}, nil
}

func (r *mutationResolver) DeleteAnswer(ctx context.Context, id string) (*model.AnswerResponse, error) {
	err := r.AnsService.DeleteAnswer(id)
	if err != nil {
		return &models.AnswerResponse{
			Message: "Something went wrong deletinng the answer.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.AnswerResponse{
		Message: "Successfully deleted answer",
		Status:  http.StatusOK,
	}, nil
}

func (r *queryResolver) GetOneAnswer(ctx context.Context, id string) (*model.AnswerResponse, error) {
	answer, err := r.AnsService.GetAnswerByID(id)
	if err != nil {
		log.Println("getting answer error: ", err)
		return &models.AnswerResponse{
			Message: "Something went wrong getting the answer.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.AnswerResponse{
		Message: "Successfully retrieved answer",
		Status:  http.StatusOK,
		Data:    answer,
	}, nil
}

func (r *queryResolver) GetAllQuestionAnswers(ctx context.Context, questionID string) (*model.AnswerResponse, error) {
	answers, err := r.AnsService.GetAllQuestionAnswers(questionID)
	if err != nil {
		log.Println("getting all questions error ", err)
		return &models.AnswerResponse{
			Message: "Something went wrong getting all questions.",
			Status:  http.StatusInternalServerError,
		}, nil
	}
	return &models.AnswerResponse{
		Message:  "Successfully retrieved all answers",
		Status:   http.StatusOK,
		DataList: answers,
	}, nil
}
