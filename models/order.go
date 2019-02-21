package models

import "time"

// Order struct defines the order model
type Order struct {
	ID        int
	StartLat  float64
	StartLong float64
	EndLat    float64
	EndLong   float64
	Distance  int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
