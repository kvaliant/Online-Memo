package models

import (
	"gorm.io/gorm"
)

type Memo struct {
	gorm.Model
	Title	string	`json:"title"`
	Body	string	`json:"body"`

	UserID	int		`json:"-"`
	User 	User	`json:"-"`
}
