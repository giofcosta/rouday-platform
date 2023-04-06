// cmd/api/data/models/week-configuration.go
package models

import "time"

type WeekConfiguration struct {
	ID              string    `json:"id"`
	WeekStart       time.Time `json:"week_start"`
	WorkHoursPerDay float64   `json:"work_hours_per_day"`
	UtilDaysPerWeek int       `json:"util_days_per_week"`
	Routines        []Routine `json:"routines"`
}