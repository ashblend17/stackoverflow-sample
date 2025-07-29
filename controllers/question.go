package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ashblend17/stackoverflow-sample/database"
	"github.com/ashblend17/stackoverflow-sample/models"

	"github.com/gin-gonic/gin"
)

type AnswerResponse struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	User      struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
}

type QuestionResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	User      struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
}

func CreateQuestion(ctx *gin.Context) {
	// Extract user_id from context (set by middleware)
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	question := models.Question{
		UserID: userID.(int),
		Title:  input.Title,
		Body:   input.Body,
	}

	if err := database.DB.Create(&question).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":     "Question created",
		"question_id": question.ID,
	})
}

func GetQuestionWithAnswers(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	// Load question
	var question models.Question
	if err := database.DB.Preload("User").First(&question, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	// Load answers
	var answers []models.Answer
	if err := database.DB.Preload("User").Where("question_id = ?", id).Order("created_at").Find(&answers).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch answers"})
		return
	}

	// Transform question
	questionRes := QuestionResponse{
		ID:        question.ID,
		Title:     question.Title,
		Body:      question.Body,
		CreatedAt: question.CreatedAt,
	}
	questionRes.User.ID = question.User.ID
	questionRes.User.Username = question.User.Username

	// Transform answers
	var answerRes []AnswerResponse
	for _, a := range answers {
		ar := AnswerResponse{
			ID:        a.ID,
			Body:      a.Body,
			CreatedAt: a.CreatedAt,
		}
		ar.User.ID = a.User.ID
		ar.User.Username = a.User.Username
		answerRes = append(answerRes, ar)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"question": questionRes,
		"answers":  answerRes,
	})

}
