package models

type Routine struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	DailyAverageHours float64 `json:"daily_average_hours"`
	Monday           float64 `json:"monday"`
	Tuesday          float64 `json:"tuesday"`
	Wednesday        float64 `json:"wednesday"`
	Thursday         float64 `json:"thursday"`
	Friday           float64 `json:"friday"`
	Saturday         float64 `json:"saturday"`
	Sunday           float64 `json:"sunday"`
	Comment          string  `json:"comment"`
}