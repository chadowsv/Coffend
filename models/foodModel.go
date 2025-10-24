package models

import "time"

type Food struct {
	//Serializacion de mapeo
	FoodID      int       `json:"food_id"`
	Name        string    `json:"name" validate:"required,min=2,max=100"`
	Description string    `json:"description" validate:"required"`
	Price       float64   `json:"price" validate:"required"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	MenuID      int       `json:"menu_id" validate:"required"`
}
type FoodMenu struct {
	MenuID      int     `json:"menu_id" validate:"required"`
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}
