package model

type Car struct {
	CarId     *int    `json:"car_id,omitempty"`
	CarName   string  `json:"car_name"`
	DayRate   float64 `json:"day_rate"`
	MonthRate float64 `json:"month_rate"`
	Image     string  `json:"image"`
}
