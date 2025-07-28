package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ashblend17/stackoverflow-sample/database"
	"github.com/ashblend17/stackoverflow-sample/models"
	"github.com/ashblend17/stackoverflow-sample/utils"

	"github.com/gin-gonic/gin"
)

func SummarizeQuestion(ctx *gin.Context) {
	questionIDStr := ctx.Param("id")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	// Get the question
	var question models.Question
	if err := database.DB.First(&question, questionID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	// Get all answers
	var answers []models.Answer
	if err := database.DB.Where("question_id = ?", questionID).Find(&answers).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch answers"})
		return
	}

	// Extract answer bodies
	answerTexts := []string{}
	for _, ans := range answers {
		answerTexts = append(answerTexts, ans.Body)
	}

	// Get summary from OpenAI
	summary, err := utils.SummarizeQnA(question.Body, answerTexts)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "LLM summarization failed"})
		fmt.Println("Error summarizing question and answers:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"question_id": questionID,
		"summary":     summary,
	})
}
