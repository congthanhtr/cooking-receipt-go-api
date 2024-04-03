package model

import "gorm.io/gorm"

type CookingReceipt struct {
	gorm.Model
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Ingredients []*Ingredient `json:"ingredients" gorm:"many2many:recipe_ingredients;" `
	Photo       string        `json:"photo"`
}
