package models

import "github.com/google/uuid"

type Subscription struct {
	ID          int       `json:"id"`
	ServiceName string    `json:"service_name" binding:"required"`
	Price       int       `json:"price" binding:"required"`
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	StartDate   string    `json:"start_date" binding:"required"`
	EndDate     *string   `json:"end_date,omitempty"`
}

type CostRequest struct {
	StartDate   string `form:"start_date" binding:"required"`
	EndDate     string `form:"end_date" binding:"required"`
	UserID      string `form:"user_id"`
	ServiceName string `form:"service_name"`
}
