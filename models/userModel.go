package models

import (
	"gorm.io/gorm"
)	

type User struct {
	gorm.Model
	ID			int		`gorm:"UNIQUE_INDEX:compositeindex"`
	Username	string	`gorm:"UNIQUE_INDEX:compositeindex" json:"username" binding:"required"`
	FullName	string	`json:"full_name" binding:"required"`
	Password 	string	`json:"-"`
}