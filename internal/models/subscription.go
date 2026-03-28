package models

type Subscription struct {
	ID          int     `json:"id" db:"id"`
	ServiceName string  `json:"service_name" db:"service_name" binding:"required"`
	Price       int     `json:"price" db:"price" binding:"required"`
	UserID      string  `json:"user_id" db:"user_id" binding:"required,uuid"`
	StartDate   string  `json:"start_date" db:"start_date" binding:"required"`
	EndDate     *string `json:"end_date,omitempty" db:"end_date"`
}

type CostRequest struct {
	StartDate   string `form:"start_date" binding:"required"`
	EndDate     string `form:"end_date" binding:"required"`
	UserID      string `form:"user_id"`
	ServiceName string `form:"service_name"`
}
