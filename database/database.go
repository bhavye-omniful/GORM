package database

import (
	"github.com/bhavye-omniful/GORM/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func Init() {
	dsn := "host=arjuna.db.elephantsql.com user=tufpnzxa password=InhZoxtF_vwoqasJXn09_gdM_VFoth5r dbname=tufpnzxa port=5432"

	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	migrate()
}

func migrate() {
	err := Db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error migrating user model : ", err.Error())
	}
}
