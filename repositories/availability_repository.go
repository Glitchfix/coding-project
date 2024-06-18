package repositories

import (
	config "github.com/Glitchfix/coding-project/config"
	models "github.com/Glitchfix/coding-project/models"
)

func GetAvailabilityByUserID(userID uint) ([]models.Availability, error) {
	var availabilities []models.Availability
	result := config.DB.Where("user_id = ?", userID).Find(&availabilities)
	return availabilities, result.Error
}

func CreateAvailability(availability *models.Availability) error {
	result := config.DB.Create(availability)
	return result.Error
}

func GetOverlappingAvailabilities(userID1, userID2 uint) ([]models.Availability, error) {
	var availabilities []models.Availability
	query := `SELECT * FROM availabilities WHERE user_id IN (?, ?) AND start_time < (SELECT end_time FROM availabilities WHERE user_id = ? ORDER BY end_time DESC LIMIT 1) AND end_time > (SELECT start_time FROM availabilities WHERE user_id = ? ORDER BY start_time ASC LIMIT 1)`
	result := config.DB.Raw(query, userID1, userID2, userID1, userID2).Scan(&availabilities)
	return availabilities, result.Error
}
