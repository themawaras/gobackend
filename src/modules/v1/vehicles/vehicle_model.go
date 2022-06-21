package vehicles

import "time"

type Vehicle struct {
	Vehicle_id  int       `gorm:"primaryKey" json:"id"`
	Manufacture string    `json:"manufacture"`
	Model       string    `json:"model"`
	Type        string    `json:"type"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Year        string    `json:"year"`
}

type Vehicles []Vehicle
