package controllers

import (
	"net/http"

	"github.com/ashblend17/stackoverflow-sample/database"
	"github.com/ashblend17/stackoverflow-sample/models"

	"github.com/gin-gonic/gin"
)

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
