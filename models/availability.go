package models

import "time"

// Availability represents the availability of a user
// swagger:model Availability
type Availability struct {
	// The ID of the availability record
	// required: true
	// example: 1
	ID uint `json:"id" gorm:"primaryKey"`

	// The ID of the user
	// required: true
	// example: 123
	UserID uint `json:"user_id"`

	// The start time of the availability
	// required: true
	// example: 2023-06-18T15:00:00Z
	StartTime time.Time `json:"start_time"`

	// The end time of the availability
	// required: true
	// example: 2023-06-18T17:00:00Z
	EndTime time.Time `json:"end_time"`
}
