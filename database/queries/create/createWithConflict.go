package create

import (
	"github.com/bhavye-omniful/GORM/database"
	"github.com/bhavye-omniful/GORM/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CreateWithConflict(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		log.Println("Error binding json to user : ", err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Println(user)

	// Do nothing on conflict
	result := database.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	if result.Error != nil {
		log.Println(result.Error.Error())
		ctx.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	// Update columns to default value on `id` conflict
	result = database.Db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "email"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"email": strconv.FormatInt(time.Now().UnixNano(), 10),
		}),
	}).Create(&user)

	//result := database.Db.Create(&user)

	if result.Error != nil {
		log.Println(result.Error.Error())
		ctx.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	ctx.JSON(200, user)
}
