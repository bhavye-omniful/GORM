package models

import (
	"errors"
	"github.com/bhavye-omniful/GORM/enums"
	"gorm.io/gorm"
)

type DayModel struct {
	gorm.Model
	Today string `json:"day"`
	Type  enums.DayType
}

func (d *DayModel) BeforeCreate(db *gorm.DB) (err error) {
	_, ok := enums.StringToEnumDay[d.Today]

	if !ok {
		return errors.New("Invalid day!")
	}

	return
}

//func (d *DayModel) AfterCreate(db *gorm.DB) (err error) {
//	if enums.StringToEnumDay[d.Today] == enums.Saturday || enums.StringToEnumDay[d.Today] == enums.Sunday {
//		db.Model(d).Update("type", enums.Holiday)
//	} else {
//		db.Model(d).Update("type", enums.Working)
//	}
//	return
//}

func (d *DayModel) AfterSave(db *gorm.DB) (err error) {
	log := Log{
		TableName: d.Today,
		Comment:   "Inserted something in day table",
	}
	db.Create(&log)
	return
}
