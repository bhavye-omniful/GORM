package models

import (
	"context"
	"errors"
	"github.com/bhavye-omniful/GORM/enums"
	"github.com/bhavye-omniful/GORM/redis"
	"gorm.io/gorm"
	"strconv"
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

	// asdfghj

	res := redis.RedisClient.HSet(context.Background(), "DayMap:"+strconv.FormatUint(uint64(d.ID), 10), map[string]string{
		"CreatedAt": d.CreatedAt.String(),
		"UpdatedAt": d.UpdatedAt.String(),
		"DeletedAt": d.DeletedAt.Time.String(),
		"Today":     d.Today,
		"Type":      string(d.Type),
	})

	if res.Err() != nil {
		return errors.New("Problem inserting in redis" + res.Err().Error())
	}

	db.Create(&log)
	return
}
