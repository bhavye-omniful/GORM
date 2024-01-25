package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	TableName string
	Comment   string
}
