package services

import (
	models "github.com/Glitchfix/coding-project/models"
	repositories "github.com/Glitchfix/coding-project/repositories"
)

func GetAvailability(userID uint) ([]models.Availability, error) {
	return repositories.GetAvailabilityByUserID(userID)
}

func CreateAvailability(availability *models.Availability) error {
	return repositories.CreateAvailability(availability)
}

func FindOverlap(userID1, userID2 uint) ([]models.Availability, error) {
	return repositories.GetOverlappingAvailabilities(userID1, userID2)
}
