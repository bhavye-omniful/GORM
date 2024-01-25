package create

import (
	"github.com/bhavye-omniful/GORM/database"
	"github.com/bhavye-omniful/GORM/enums"
	"github.com/bhavye-omniful/GORM/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateDay(ctx *gin.Context) {
	var day models.DayModel

	err := ctx.ShouldBindJSON(&day)
	if err != nil {
		log.Println("Error binding json to user : ", err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Println(day)

	if enums.StringToEnumDay[day.Today] == enums.Saturday || enums.StringToEnumDay[day.Today] == enums.Sunday {
		day.Type = enums.Holiday
	} else {
		day.Type = enums.Working
	}

	result := database.Db.Create(&day)
	if result.Error != nil {
		log.Println(result.Error.Error())
		ctx.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	ctx.JSON(200, day)
}
