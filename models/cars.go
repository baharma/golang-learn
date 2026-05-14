package models

import "time"

type Car struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	ProductsId int       `json:"products_id" gorm:"index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
