package controllers

import (
	"net/http"
	"strconv"

	"github.com/ashblend17/stackoverflow-sample/database"
	"github.com/ashblend17/stackoverflow-sample/models"

	"github.com/gin-gonic/gin"
)

func CreateAnswer(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	questionIDStr := ctx.Param("id")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	var input struct {
		Body string `json:"body"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	answer := models.Answer{
		UserID:     userID.(int),
		QuestionID: questionID,
		Body:       input.Body,
	}

	if err := database.DB.Create(&answer).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to post answer"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":   "Answer posted",
		"answer_id": answer.ID,
	})
}
