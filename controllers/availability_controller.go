package controllers

import (
	"net/http"
	"strconv"

	"github.com/Glitchfix/coding-project/models"
	"github.com/Glitchfix/coding-project/services"
	"github.com/gin-gonic/gin"
)

// GetAvailability godoc
// @Summary Get availability for a user
// @Description Get availability for a given user ID
// @Tags availability
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Success 200 {array} models.Availability
// @Failure 400 {object} gin.H "error response"
// @Failure 500 {object} gin.H "error response"
// @Router /availability/{user_id} [get]
func GetAvailability(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	availabilities, err := services.GetAvailability(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, availabilities)
}

// CreateAvailability godoc
// @Summary Create availability for a user
// @Description Create availability for a user
// @Tags availability
// @Accept  json
// @Produce  json
// @Param availability body models.Availability true "Availability"
// @Success 201 {object} models.Availability
// @Failure 400 {object} gin.H "error response"
// @Failure 500 {object} gin.H "error response"
// @Router /availability [post]
func CreateAvailability(c *gin.Context) {
	var availability models.Availability
	if err := c.ShouldBindJSON(&availability); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateAvailability(&availability); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, availability)
}

// FindOverlap godoc
// @Summary Find overlapping availability between two users
// @Description Find overlapping availability between two users
// @Tags availability
// @Accept  json
// @Produce  json
// @Param user_id1 path int true "User ID 1"
// @Param user_id2 path int true "User ID 2"
// @Success 200 {array} models.Availability
// @Failure 400 {object} gin.H "error response"
// @Failure 500 {object} gin.H "error response"
// @Router /overlap/{user_id1}/{user_id2} [get]
func FindOverlap(c *gin.Context) {
	userID1, err := strconv.Atoi(c.Param("user_id1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID1"})
		return
	}
	userID2, err := strconv.Atoi(c.Param("user_id2"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID2"})
		return
	}

	availabilities, err := services.FindOverlap(uint(userID1), uint(userID2))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, availabilities)
}
