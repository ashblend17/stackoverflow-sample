package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ashblend17/stackoverflow-sample/database"
	"github.com/ashblend17/stackoverflow-sample/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func VoteHandler(itemType string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Get("user_id")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		itemIDStr := ctx.Param("id")
		itemID, err := strconv.Atoi(itemIDStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
			return
		}

		var input struct {
			Vote string `json:"vote"` // "upvote", "downvote", or "remove"
		}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if input.Vote != "upvote" && input.Vote != "downvote" && input.Vote != "remove" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vote type"})
			return
		}

		// Check if vote already exists
		var existing models.Vote
		err = database.DB.Where("user_id = ? AND item_id = ? AND item_type = ?", userID, itemID, itemType).First(&existing).Error

		if input.Vote == "remove" {
			if err == nil {
				// Vote exists → delete it
				database.DB.Delete(&existing)
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Vote removed"})
			return
		}

		if err == nil {
			// Vote exists → update it
			existing.VoteType = input.Vote
			existing.UpdatedAt = time.Now()
			database.DB.Save(&existing)
			ctx.JSON(http.StatusOK, gin.H{"message": "Vote updated"})
			return
		}

		// New vote
		vote := models.Vote{
			UserID:    userID.(int),
			ItemID:    itemID,
			ItemType:  itemType,
			VoteType:  input.Vote,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = database.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "item_id"}, {Name: "item_type"}},
			UpdateAll: true,
		}).Create(&vote).Error

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cast vote"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Vote recorded"})
	}
}
