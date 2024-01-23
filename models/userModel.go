package models

import "gorm.io/gorm"

type Address struct {
	HouseNo string `json:"houseNo"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type User struct {
	gorm.Model
	Name    string  `gorm:"not null" json:"name"`
	Email   string  `gorm:"unique;not null" json:"email"`
	Address Address `gorm:"embedded;embeddedPrefix:User_" json:"address"`
}
